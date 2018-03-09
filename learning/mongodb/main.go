package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{"127.0.0.0:3000"},
		Database: "abc",
		Source:   "admin",
		Username: "abc",
		Password: "abc",
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Println(err)
	}
	defer session.Close()
	db := session.DB("ckv")
	// 列出所有collections
	names, err := db.CollectionNames()
	if err != nil {
		log.Println(err)
	}
	for i, v := range names {
		fmt.Println(i, v)
	}

	// 执行db.collection.stats()
	var doc bson.M
	err = session.DB("ckv").Run(bson.D{
		{Name: "collStats", Value: "abc"}}, &doc)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(doc)

}
