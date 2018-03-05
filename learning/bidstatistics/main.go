// log soup channel 练习
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/anaskhan96/soup"
	"gopkg.in/natefinch/lumberjack.v2"
)

// init log configuration
func init() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "bidstatistics.log",
		MaxSize:    1,
		MaxBackups: 3,
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

type auth struct {
	User     string `toml:"username"`
	Password string `toml:"password"`
}

type Config struct {
	Regions       []string `toml:"regions"`
	BusinessGroup []string `toml:"business_group"`
	Database      database
	Auth          auth
}

// retrieve regions from configuration
func retrieveRegions(regions []string, regionChannel chan string) {
	defer close(regionChannel)
	for _, region := range regions {
		regionChannel <- region
	}
}

// retrieve houses from regions
func retrieveHouse(regionChannel chan string, houseChannel chan string) {
	defer close(houseChannel)
	var wg sync.WaitGroup
	for region := range regionChannel {
		wg.Add(1)
		// pass the **pointer** to the sync.WaitGroup instead of a sync.WaitGroup
		go parseRegion(region, houseChannel, &wg)
	}
	wg.Wait()
	log.Println("[INFO] retrieveHouse complete")
}

// parse region, return house url
func parseRegion(region string, houseChannel chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("[DEBUG] parseRegion start: ", region)
	// construct authed http request
	client := http.DefaultClient
	req, err := http.NewRequest("GET", region, nil)
	req.SetBasicAuth(config.Auth.User, config.Auth.Password)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[ERROR]", err)
		return
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[ERROR]", err)
		return
	}
	doc := soup.HTMLParse(string(bodyText))
	tbody := doc.Find("tbody")
	if tbody.NodeValue == "" {
		log.Println("[WARN] not found tbody in", region)
		return
	}
	u, err := url.Parse(region)
	if err != nil {
		log.Println("[ERROR]", err)
		return
	}
	baseUrl := strings.Replace(region, u.RequestURI(), "", 1)
	for _, tr := range tbody.FindAll("tr") {
		cols := tr.FindAll("td")
		if len(cols) < 3 {
			log.Println("[INFO] invalid td ", cols)
			return
		}
		bg := cols[2].Text()
		for _, ok := range config.BusinessGroup {
			if bg == ok {
				house := baseUrl + cols[0].Find("a").Attrs()["href"]
				log.Println("[DEBUG] house url:", house)
				houseChannel <- house
			}
		}
	}
	log.Println("[DEBUG] parseRegion end: ", region)
}

func retrieveSet(houseChannel chan string, setChannel chan string) {
	defer close(setChannel)
	var wg sync.WaitGroup
	for house := range houseChannel {
		wg.Add(1)
		go parseSet(house, setChannel, &wg)
	}
	wg.Wait()
	log.Println("[INFO] retrieveSet complete")
}

func parseSet(house string, setChannel chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("[DEBUG] parseSet start: ", house)
	client := http.DefaultClient
	req, err := http.NewRequest("GET", house, nil)
	req.SetBasicAuth(config.Auth.User, config.Auth.Password)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[ERROR]", err)
		return
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[ERROR]", err)
		return
	}
	doc := soup.HTMLParse(string(bodyText))
	tbody := doc.Find("tbody")
	if tbody.NodeValue == "" {
		log.Println("[WARN] not found tbody in", house)
		return
	}
	u, err := url.Parse(house)
	if err != nil {
		log.Println("[ERROR]", err)
		return
	}
	queryString, _ := url.ParseQuery(u.RawQuery)
	dbip := queryString["dbip"][0]
	params := "xxx&dbip=%s&masterip=%s"
	baseUrl := strings.Replace(house, u.RequestURI(), "", 1)
	rows := tbody.FindAll("tr")
	for _, row := range rows {
		cols := row.FindAll("td")
		if len(cols) < 2 {
			log.Println("[INFO] invalid td ", cols)
			return
		}
		if cols[1].Text() != "-" {
			masterIp := cols[1].Text()
			set := baseUrl + fmt.Sprintf(params, dbip, masterIp)
			setChannel <- set
			log.Println("[DEBUG] set ip:", set)
		}
	}
	log.Println("[DEBUG] parseSet end: ", house)
}

var config Config

func main() {

	_, err := toml.DecodeFile("bidstatistics.toml", &config)
	if err != nil {
		log.Println("[ERROR]", err)
	}

	// read regions value from configuration, write to channel
	regionChannel := make(chan string, len(config.Regions))
	go retrieveRegions(config.Regions, regionChannel)

	// read region url from channel regionChannel, parse house url and write to channel houseChannel
	houseChannel := make(chan string, 100)
	go retrieveHouse(regionChannel, houseChannel)

	// read house url from channel houseChanel, parse seturl and write to chanel setChannel
	setChannel := make(chan string, 100)
	go retrieveSet(houseChannel, setChannel)
	for set := range setChannel {
		fmt.Println(set)
	}
}
