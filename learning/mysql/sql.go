package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:Paic1234@tcp(127.0.0.1:3306)/config?charset=utf8")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from cmongo_metadata")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	type row struct {
		id         int
		name       string
		cl5_mod_id string
		cl5_cmd_id string
		proxy      string
		dbname     string
		source     string
		username   string
		password   string
		enabled    bool // boolean type in mysql
	}
	got := []row{}
	for rows.Next() {
		var r row
		err = rows.Scan(&r.id,
			&r.name,
			&r.cl5_mod_id,
			&r.cl5_cmd_id,
			&r.proxy,
			&r.dbname,
			&r.source,
			&r.password,
			&r.username,
			&r.enabled,
		)
		if err != nil {
			log.Println(err)
		}
		got = append(got, r)
	}
	for _, v := range got {
		fmt.Println(v)
	}
}
