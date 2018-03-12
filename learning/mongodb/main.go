package main

import (
	"fmt"
	"log"
	"sync"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{"10.10.10.10:6000"},
		Database: "database",
		Source:   "admin",
		Username: "username",
		Password: "password",
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Println(err)
	}
	defer session.Close()
	// 切换database
	db := session.DB("database")
	// 列出所有collections
	names, err := db.CollectionNames()
	if err != nil {
		log.Println(err)
	}
	var wg sync.WaitGroup
	for i, v := range names {
		wg.Add(1)
		// 执行db.collection.stats()
		go func(i int, v string) {
			var doc bson.M
			// 运行database命令
			// example: https://godoc.org/gopkg.in/mgo.v2#Database.Run
			// command: https://docs.mongodb.com/manual/reference/command/
			err = session.DB("database").Run(bson.D{
				{Name: "collStats", Value: v},
				{"scale", 1024 * 1024}, //collStata的参数
			}, &doc)
			if err != nil {
				log.Println(err)
			}
			if size, ok := doc["storageSize"]; !ok {
				log.Println("not found storageSize")
			} else {
				fmt.Printf("%s storageSize: %d Mb\n", v, size)
			}
			wg.Done()
		}(i, v)
	}
	wg.Wait()
	fmt.Println("complete")
}
