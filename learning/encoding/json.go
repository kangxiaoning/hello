package main

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"reflect"
)

// 注意变量名要大写，因为json和main是不同的package，大写才能exported被json使用
type DBConfig struct {
	Host     string `json: "host"`
	Port     int    `json: "port"`
	User     string `json: "user"`
	Password string `json: "password"`
	Database string `json: "database"`
}

func main() {
	file, err := os.Open("./db.json")
	defer file.Close()
	if err != nil {
		log.Fatalln("open ./db.json error")
	}

	// method 1: json.Unmarshal
	byteValue, _ := ioutil.ReadAll(file)
	var config DBConfig
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatalln("json decode failed")
	}
	defer file.Close()
	fmt.Println("Method 1: use json.Unmarshal(byteValue, &config)")
	fmt.Printf("\thost: %s\n", config.Host)
	fmt.Printf("\tport: %d\n", config.Port)
	fmt.Printf("\tuser: %s\n", config.User)
	fmt.Printf("\tpassword: %s\n", config.Password)
	fmt.Printf("\tdatabase: %s\n", config.Database)
	fmt.Println()

	// json.Marshal() example
	db := &DBConfig{
		Host:     "192.168.1.10",
		Port:     1521,
		User:     "oracle",
		Password: "ILoveChina",
		Database: "orcl",
	}
	fmt.Println("\tjson.Marshal(db) example")
	fmt.Printf("\t\tbefore json.Marshal type is: %s\n", reflect.TypeOf(db))
	// func Marshal(v interface{}) ([]byte, error)
	ret, err := json.Marshal(db)
	if err != nil {
		log.Fatalln(err)
	}
	// bytes to string

	fmt.Println("\t\tstring representation:\n\t\t\t", string(ret))
	fmt.Printf("\t\tafter json.Marshal type is: %s\n", reflect.TypeOf(ret))
	fmt.Println()
	// method 2: json.Decoder
	// func (dec *Decoder) Decode(v interface{}) error
	file, err = os.Open("./db.json")
	defer file.Close()
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Method 2: use json.NewDecoder(file).Decode(&config)")
	fmt.Printf("\thost: %s\n", config.Host)
	fmt.Printf("\tport: %d\n", config.Port)
	fmt.Printf("\tuser: %s\n", config.User)
	fmt.Printf("\tpassword: %s\n", config.Password)
	fmt.Printf("\tdatabase: %s\n", config.Database)
}
