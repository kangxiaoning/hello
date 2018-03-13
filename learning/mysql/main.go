package main

import (
	"fmt"
	"log"
	"sync"

	"time"

	"github.com/natefinch/lumberjack"

	"database/sql"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var config Config

// init log configuration
func init() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "cold_data_stats.log",
		MaxSize:    10,
		MaxBackups: 2,
	})
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type database struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type cluster struct {
	Addrs    []string
	Database string
	Source   string
	Username string
	Password string
}

type Config struct {
	Database database
	Clusters []cluster `toml:"cluster"`
}

func queryBidStats(cluster cluster, wg *sync.WaitGroup, c chan result) {
	defer wg.Done()
	start := time.Now()
	dialInfo := &mgo.DialInfo{
		Addrs:    cluster.Addrs,
		Database: cluster.Database,
		Source:   cluster.Source,
		Username: cluster.Username,
		Password: cluster.Password}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	defer session.Close()
	db := session.DB(cluster.Database)
	// get all collections
	names, err := db.CollectionNames()
	if err != nil {
		log.Println("[ERROR]", err)
	}
	var qwg sync.WaitGroup
	for _, v := range names {
		qwg.Add(1)
		go func(bid string) {
			defer qwg.Done()
			var doc bson.M
			// run database command
			// example: https://godoc.org/gopkg.in/mgo.v2#Database.Run
			// command: https://docs.mongodb.com/manual/reference/command/
			err = session.DB(cluster.Database).Run(bson.D{
				{Name: "collStats", Value: bid},
				//collStats parameter
				{"scale", 1024 * 1024 * 1024},
			}, &doc)
			if err != nil {
				log.Println("[ERROR]", err)
			}
			if size, ok := doc["storageSize"]; !ok {
				log.Println("[WARN] not found storageSize")
			} else {
				var r = result{bid, size}
				c <- r
				log.Printf("[INFO] got %s %d\n", bid, size)
			}
		}(v)
	}
	qwg.Wait()
	log.Printf("[INFO] %s querystats elapsed time %v",
		cluster.Addrs,
		time.Since(start))
}

// save to database
func save(c chan result) {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
			config.Database.User,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Database))
	if err != nil {
		log.Println("[ERROR]", err)
	}
	defer db.Close()
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(20)
	rowCount := 0
	for record := range c {
		_, err := db.Exec("INSERT INTO test_cmongo_stats VALUES (?,?)",
			record.bid,
			record.storageSize)
		if err != nil {
			log.Println("[ERROR]", err)
		} else {
			rowCount++
		}
		log.Printf("[INFO] save %s %d\n", record.bid, record.storageSize)
	}
	log.Printf("[INFO] inserted %d rows", rowCount)
}

type result struct {
	bid         string
	storageSize interface{}
}

func Worker(clusters []cluster, c chan result) {
	defer close(c)
	var wg sync.WaitGroup
	for _, cluster := range clusters {
		wg.Add(1)
		go queryBidStats(cluster, &wg, c)
	}
	wg.Wait()
	log.Println("[INFO] Worker complete")
}

// load config
func loadConfig(fpath string) {
	_, err := toml.DecodeFile(fpath, &config)
	if err != nil {
		fmt.Println(err)
		log.Println("[ERROR]", err)
	}
}

func main() {
	start := time.Now()
	loadConfig("collstats.toml")
	// data channel
	resultCh := make(chan result, 100)
	// generate data on all clusters in concurrency
	go Worker(config.Clusters, resultCh)
	// receive data and save to mysql
	save(resultCh)
	log.Println("[INFO] elapsed time", time.Since(start))
}
