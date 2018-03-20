package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	monitor := "http://monitor.example.com/monitor/api/common/query/"

	message := `
{
    "rsp_format": 2, 
    "diff": 0, 
    "metrics": [
        {
            "target": "sum(reads)"
        },
        {
            "target": "sum(updates)"
        }
    ], 
    "fields": {
        "app": "mongodb", 
        "role_path": "monitor.table", 
        "instance": "db.collection"
    }, 
    "time": {
        "start": 1520522952184,
        "end": 1521522952184
    }, 
    "interval": "1m", 
    "page": 1, 
    "page_size": 45
}`
	var jsonStr = []byte(message)

	// post json
	req, err := http.NewRequest("POST", monitor, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	// set timeout
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
