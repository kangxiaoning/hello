package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

var message = `
	{
		"msg": "", 
		"code": 0, 
		"data": {
			"views": [
				{
					"title": "monitor.table-ckv.b101010345-sum(reads)", 
					"start_time": 1521528840, 
					"id": "ade7dbd9db971f9fe0e219ddcf06b561", 
					"values": [
						70679
					], 
					"end_time": 1521528900, 
					"timestamps": [
						1521528840000
					], 
					"metric": {
						"instance": "ckv.b101010345", 
						"app": "cmongo", 
						"role_path": "monitor.table", 
						"target": "sum(reads)"
					}
				}, 
				{
					"title": "monitor.table-ckv.b101010345-sum(updates)", 
					"start_time": 1521528840, 
					"id": "2055c337e6f5f6d87a4d17d636e1d2b7", 
					"values": [
						24904
					], 
					"end_time": 1521528900, 
					"timestamps": [
						1521528840000
					], 
					"metric": {
						"instance": "ckv.b101010345", 
						"app": "cmongo", 
						"role_path": "monitor.table", 
						"target": "sum(updates)"
					}
				}
			], 
			"page_count": 1, 
			"page": 1, 
			"page_size": 45, 
			"raw_count": 2
		}
	}`

type timestamp int64

type metric struct {
	Instance string `json:"instance"`
	Target   string `json:"target"`
}

// embedded struct can not be Unmarshal correctly
type view struct {
	Value     []int64     `json:"values"`
	Timestamp []timestamp `json:"timestamps"`
	Metric    metric
}

type data struct {
	Views []view
}

type result struct {
	Code int `json:"code"`
	Data data
}

func main() {

	var ret result
	err := json.Unmarshal([]byte(message), &ret)
	fmt.Println(err)
	fmt.Println(ret.Code)
	// fmt.Println(ret.Data.Views)
	for _, v := range ret.Data.Views {
		// fmt.Println(strings.Replace(v.Metric.Instance, "ckv.b", "", 1), v.Metric.Target, v.Timestamp[0].toTime(), v.Value[0])
		fmt.Println(strings.Replace(v.Metric.Instance, "ckv.b", "", 1), v.Metric.Target, v.Timestamp[0].toTime(), v.Value[0])
	}
}

func (t timestamp) toTime() string {
	return time.Unix(0, int64(t)*int64(time.Millisecond)).Format("2006-01-02 15:04:05")
}
