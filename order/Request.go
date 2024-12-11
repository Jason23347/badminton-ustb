package order

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func SendOrder(nextDate string, cdType string, cdNumbers []int) (string, error) {
	cdstring := getCDString(cdType, cdNumbers)
	searchParam := fmt.Sprintf(
		`{"datestring":"%s","cdstring":"%s","paytype":"M"}`,
		nextDate, cdstring)
	encodedSearchParam := url.QueryEscape(searchParam)
	// 定义参数
	wxKey := "E7E7EB4C8EC1A817B3858271B986FBBA0ECE35796DD6B28956063323C0239EA6C7F2D849B78B6638C0697E569096ECF966494AF46C5281B2A1F4AEFD32105705A7DCD001993984734F02187DE30B34202ECA2A07D0ED39E44011DEF523F3414772AE1EF9236B36ADDEA3E5EA9D3D6D95"
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
