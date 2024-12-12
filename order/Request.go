package order

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func SendOrder(fieldType string, fields []Field) (string, error) {
	var cdNumbers []int
	for _, field := range fields {
		cdNumbers = append(cdNumbers, field.Number)
	}
	cdstring := getCDString(fieldType, cdNumbers)
	searchParam := fmt.Sprintf(
		`{"datestring":"%s","cdstring":"%s","paytype":"M"}`,
		GetConfigInstance().Date, cdstring)
	encodedSearchParam := url.QueryEscape(searchParam)
	// 定义参数
	wxKey := GetUserInstance().WXKey
	urlStr := "http://32901.koksoft.com/HomefuntionV2json.aspx"

	// 设置请求头
	headers := map[string]string{
		"Content-Type":     "application/x-www-form-urlencoded; charset=UTF-8",
		"Origin":           "http://32901.koksoft.com",
		"Referer":          "http://32901.koksoft.com/weixinordernewv7.aspx?wxkey=" + wxKey,
		"User-Agent":       "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Mobile Safari/537.36 Edg/131.0.0.0",
		"X-Requested-With": "XMLHttpRequest",
	}

	// 准备请求数据
	data := url.Values{}
	data.Set("searchparam", encodedSearchParam)
	data.Set("wxkey", wxKey)
	data.Set("classname", "saasbllclass.CommonFuntion")
	data.Set("funname", "MemberOrderfromWx")

	// 创建 POST 请求
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating request:", err)
		return "", err
	}

	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 执行请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading response:", err)
		return "", err
	}

	// 输出响应
	return string(body), nil
}

func getCDString(cdType string, cdNumbers []int) string {
	// 生成 cdstring
	cdstring := ""
	for _, num := range cdNumbers {
		cd := Field{Type: cdType, Number: num}
		cdstring += cd.String()
	}

	return cdstring
}

func GetForm() (IntervalForm, error) {
	// 构造请求 URL 和参数
	baseURL := "http://32901.koksoft.com/GetForm.aspx"
	params := url.Values{}
	params.Set("datatype", "viewchangdi4weixinv")
	params.Set("pagesize", "0")
	params.Set("pagenum", "0")
	params.Set("searchparam", "orderdate="+GetConfigInstance().Date)
	params.Set("wxkey", GetUserInstance().WXKey)

	// 发送 GET 请求
	reqURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return IntervalForm{}, fmt.Errorf("error creating request: %v", err)
	}

	// 设置请求头
	req.Header.Set("Referer", "http://32901.koksoft.com/weixinordernewv7.aspx?wxkey="+GetUserInstance().WXKey)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Mobile Safari/537.36 Edg/131.0.0.0")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	// 执行请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return IntervalForm{}, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// 使用 json.Decoder 进行流式解析
	decoder := json.NewDecoder(resp.Body)

	// 解析数组的开头
	if _, err := decoder.Token(); err != nil {
		return IntervalForm{}, fmt.Errorf("error reading array start token: %v", err)
	}

	for {
		// 获取每个数组元素
		t, err := decoder.Token()
		if err == io.EOF {
			break // 结束解析
		}
		if err != nil {
			return IntervalForm{}, fmt.Errorf("error reading token: %v", err)
		}

		switch v := t.(type) {
		case bool:
			// 处理布尔值
			continue
		case string:
			// 如果是 JSON 字符串类型，将其解析为 RawForm

			// 创建一个新的 Decoder 来解析 rawContent
			rawContentDecoder := json.NewDecoder(io.Reader(&stringReader{data: v}))

			{
				_, err := rawContentDecoder.Token() // ???
				if err != nil {
					return IntervalForm{}, fmt.Errorf("error reading token: %v", err)
				}
				_, err = rawContentDecoder.Token() // "rows"
				if err != nil {
					return IntervalForm{}, fmt.Errorf("error reading token: %v", err)
				}
				_, err = rawContentDecoder.Token() // array start
				if err != nil {
					return IntervalForm{}, fmt.Errorf("error reading token: %v", err)
				}
			}

			// 逐个解析 rows 数组中的 Form
			for {
				// 解析 Form 对象
				var form IntervalForm
				if err := rawContentDecoder.Decode(&form); err == io.EOF {
					// 如果解析完成，退出
					break
				}

				// 一旦找到 Timemc == "12:00" 的 Form，立即返回
				if form.StartTime == "12:00" {
					return form, nil
				}
			}
		default:
			fmt.Fprintln(os.Stderr, "JSON parsing: Unknown type", v)
		}
	}

	return IntervalForm{}, nil
}

// 自定义 stringReader，用于将字符串作为 io.Reader 传递给 Decoder
type stringReader struct {
	data string
	pos  int
}

func (s *stringReader) Read(p []byte) (n int, err error) {
	// 计算剩余的字符串长度
	remaining := len(s.data) - s.pos
	if remaining <= 0 {
		return 0, io.EOF
	}

	// 确保不会读取超出字符串的部分
	toRead := remaining
	if toRead > len(p) {
		toRead = len(p)
	}

	// 将数据复制到 p 中
	copy(p, s.data[s.pos:s.pos+toRead])
	s.pos += toRead

	return toRead, nil
}
