package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"fmt"
	"time"
)

// 注意: 不加tags时struct里的字段必须和配置文件中的拼写一致
type Config struct {
	Database  map[string]database
	Parameter parameter
}

type database struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type parameter struct {
	MaxThreads int       `toml:"max_threads"`
	IsEnabled  bool      `toml:"is_enabled"`
	StartTime  time.Time `toml:"start_time"`
}

func main() {
	// 从文件解析配置
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(meta)
	for k, v := range config.Database {
		fmt.Println(k, v)
	}

	// 获取配置
	if source, ok := config.Database["source"]; !ok {
		fmt.Println("'source' not exist in configuration")
	} else {
		fmt.Println(source.Host)
		fmt.Println(source.Port)
		fmt.Println(source.User)
		fmt.Println(source.Password)
		fmt.Println(source.Database)
	}

	// 测试不存在的配置
	if test, ok := config.Database["test"]; !ok {
		fmt.Println("'test' not exist in configuration")
	} else {
		fmt.Println(test.Host)
		fmt.Println(test.Port)
		fmt.Println(test.User)
		fmt.Println(test.Password)
		fmt.Println(test.Database)
	}

	fmt.Println("配置时间是： ", config.Parameter.StartTime)
}
