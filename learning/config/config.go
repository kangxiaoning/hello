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
	Number    int
	Cmem      map[string]cmem
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

// cmem
type master struct {
	DbIp     string   `toml:"db_ip"`
	MasterIp []string `toml:"master_ip"`
}

type cmem struct {
	Url    string   `toml:"url"`
	Master []master `toml:"master"`
}

//type Config struct {
//	Number int
//	Cmem   map[string]cmem
//}

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
	cmem := config.Cmem
	if sz, ok := cmem["sz"]; !ok {
		fmt.Println("not found sz config")
	} else {
		fmt.Println(sz.Url)
		for _, m := range sz.Master {
			fmt.Println(m.DbIp)
			for _, mip := range m.MasterIp {
				fmt.Printf("\t%s\n", mip)
			}
		}
	}
	//fmt.Println(cmem)
}

// 执行结果
//➜  config git:(master) ✗ ./config
//target {192.192.192.10 3306 user01 123456 test01}
//source {192.192.192.10 3306 user01 123456 test01}
//192.192.192.10
//3306
//user01
//123456
//test01
//'test' not exist in configuration
//配置时间是：  2018-02-11 18:00:00 +0000 UTC
//http://example.com/cgi-bin/trmem_list_bid.cgi?act=list&start=0&num=%s&ver=3.0&masterport=9020&dbip=%s&masterip=%s
//192.137.5.25
//192.137.155.29
//192.129.129.151
//192.175.127.158
//192.137.5.18
//192.175.117.74
//192.185.19.147
//➜  config git:(master) ✗
