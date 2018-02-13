package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"github.com/anaskhan96/soup"
	"golang.org/x/text/transform"
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"errors"
	"strings"
)

// GBK编码转换为UTF-8
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// 查找需要解析的tbody
func getTbody(doc soup.Root) (soup.Root, error) {
	var tbody soup.Root
	for _, tb := range doc.FindAll("tbody") {
		id, ok := tb.Attrs()["id"]
		// 获取属性id等于bid_list_tbody的tbody
		if ok && id == "bid_list_tbody" {
			tbody = tb
			log.Println("find tbody(id='bid_list_tbody')")
			return tbody, nil
		}
	}
	// 如果没找到则返回错误信息
	return tbody, errors.New("not found tbody")
}

func main() {
	// 屏蔽敏感信息
	url := "http://xxx"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	// 认证
	req.SetBasicAuth("xxx", "xxx")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)

	// 网页源码charset=gb2312，不转换会乱码
	bodyText, err = GbkToUtf8(bodyText)
	s := string(bodyText)
	doc := soup.HTMLParse(s)

	tbody, err := getTbody(doc)
	if err != nil {
		log.Println("not found tbody")
	}
	// 遍历tr
	for _, tr := range tbody.FindAll("tr") {
		// 遍历td
		for i, td := range tr.FindAll("td")[:8] {
			if i == 0 || i == 1 {
				// 获取td中a标签的值
				fmt.Printf("%d \t %s \n", i, td.Find("a").Text())
			} else if i == 5 || i == 6 {
				// 获取td中title属性的值
				title := td.Attrs()["title"]
				// 获取td的值
				value := td.Text()
				fmt.Printf("%d \t %s \n", i, title+value)
			} else if i == 7 {
				temp := []string{}
				// 获取td中font标签的值
				for _, v := range td.FindAll("font") {
					temp = append(temp, v.Text())
				}
				fmt.Printf("%d \t %s \n", i, strings.Join(temp, "|"))
			} else {
				// 获取td的值
				fmt.Printf("%d \t %s \n", i, td.Text())
			}
		}
	}
}

// 需要解析的tr部分信息示例
//<tr align="center">
//<td align="left"><a href="#bid_add_pos" onClick="xxx">20500001</a></td>
//<td style="word-break:break-all"  align="left" ><a href="xxx">xxx</a></td>
//<td align="left"><span class='style4'>热备</span></td>
//<td align="left">73</td>
//<td align="left">100</td>
//<td align="left" title="max=964651100 (67%)">651847897</td>
//<td align="left" title="max=73873 M (84%)">62711 M</td>
//<td align="left"><font color=green>84%</font> | <font color=green>85%</font></td>
//<td align="left">空闲</td>
//<td align="left" title="">均衡</td>
//...
//</tr>

// 结果

//0 	 20500007
//1 	 xxx
//2
//3 	 77:85
//4 	 11
//5 	 max=101314204 (48%)49290122
//6 	 max=8255 M (81%)6721 M
//7 	 81%|85%
