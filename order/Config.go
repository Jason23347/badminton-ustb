package order

import (
	"flag"
	"sync"
	"time"
)

type User struct {
	WXKey string
}

const (
	TEST_MODE int = 0
	LOOP_MODE int = 1
)

type Config struct {
	Date                string
	FieldCount          int
	RequestMode         int
	MaxLoopCount        int
	LoopIntervalSeconds int
	RequestForm         bool
}

var (
	userOnce       sync.Once
	userInstance   *User
	configOnce     sync.Once
	configInstance *Config
)

// GetInstance 返回单例实例
func GetUserInstance() *User {
	userOnce.Do(func() {
		// 只会执行一次
		userInstance = &User{
			WXKey: "E7E7EB4C8EC1A817B3858271B986FBBA0ECE35796DD6B28956063323C0239EA6C7F2D849B78B6638C0697E569096ECF966494AF46C5281B2A1F4AEFD32105705A7DCD001993984734F02187DE30B34202ECA2A07D0ED39E44011DEF523F3414772AE1EF9236B36ADDEA3E5EA9D3D6D95",
		}
	})
	return userInstance
}

func GetConfigInstance() *Config {
	configOnce.Do(func() {
		// 只会执行一次
		configInstance = &Config{}
	})
	return configInstance
}

func ParseConfig() *Config {
	defaultDate := time.Now().AddDate(0, 0, 7).Format("2006-01-02")
	config := GetConfigInstance()

	// 定义命令行参数
	flag.IntVar(&config.FieldCount, "n", 2, "要订的场地数")
	flag.IntVar(&config.RequestMode, "m", TEST_MODE, "请求模式 (0: 一次, 1: 循环)，默认一次")
	flag.IntVar(&config.MaxLoopCount, "r", 5, "场地信息最大请求次数")
	flag.IntVar(&config.LoopIntervalSeconds, "s", 1, "场地信息请求间隔秒数")
	flag.StringVar(&config.Date, "d", defaultDate, "订场日期，默认today + 7")
	flag.BoolVar(&config.RequestForm, "f", false, "订场前先请求场地信息")

	// 解析命令行参数
	flag.Parse()

	return config
}
