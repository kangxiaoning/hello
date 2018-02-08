package main

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {
	config, err := yaml.ReadFile("dbconfig.yaml")
	if err != nil {
		fmt.Println(err)
	}
	// 获取所有配置
	fmt.Println(config.Root)
	// 获取source配置
	src, err := yaml.Child(config.Root, "source")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(src)

	// 获取host信息
	if s_host, err := yaml.Child(src, "host"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("source host value: ", s_host)
	}
}
