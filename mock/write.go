/*
creatTime: 2021/1/21
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.15.3
*/

package mock

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// write date to gdb

// send post request
const (
	ip               = "http://192.168.0.199"
	port             = "8082"
	TimeFormatString = "2006-01-02 15:04:05"
)

func main() {
	mockWrite("test.xlsx")
}

func mockWrite(filePath string) {
	contents, _ := readExcel(filePath)
	count := 1
	s := rand.NewSource(64)
	rr := rand.New(s)
	var values1 []string
	for range contents {
		values1 = append(values1, strconv.Itoa(rr.Intn(1000)))
	}
	for {
		ct := time.Now().Second()
		if ct == 10 || ct == 11 || ct == 12 || ct == 13 || ct == 20 || ct == 21 || ct == 22 || ct == 23 {
			c1 := [][]string{contents, values1}
			content := map[string]interface{}{"itemValues": c1, "groupName": "1DCS"}
			requestData, _ := json.Marshal(content)
			r, err := post(requestData, count)
			fmt.Printf(", 写入结果: %s, v: %s\n", fmt.Sprintf("%s", r), values1[0])
			if err != nil {
				log.Fatal(err)
			}
			count++
			time.Sleep(time.Second)
		} else {
			var values []string
			for range contents {
				values = append(values, strconv.Itoa(rr.Intn(1000)))
			}
			c1 := [][]string{contents, values}
			content := map[string]interface{}{"itemValues": c1, "groupName": "1DCS"}
			requestData, _ := json.Marshal(content)
			r, err := post(requestData, count)
			fmt.Printf(", 写入结果: %s, v: %s\n", fmt.Sprintf("%s", r), values[0])
			if err != nil {
				log.Fatal(err)
			}
			count++
			time.Sleep(time.Second)
		}
	}
}

func readExcel(filePath string) ([]string, error) {
	f, err := excelize.OpenFile(filePath) // 打开excel文件
	if err != nil {
		return nil, err
	}
	rows, err := f.Rows("Sheet1")
	if err != nil {
		return nil, err
	}
	// test.xls中对应的列分别为0(pointName), 3(Description),4(Unit),5(Source),7(Upper),8(Lower)
	var insertedData []string
	for rows.Next() {
		row, err := rows.Columns()
		if row[0] == "JL1_10DAS27A:LBQ21AA001ZF.CIN111111" {
			fmt.Println(111)
		}
		insertedData = append(insertedData, row[0])
		if err != nil {
			return nil, err
		}

	}
	return insertedData[1:], nil
}

func post(requestBody []byte, count int) ([]byte, error) {
	sb := strings.Builder{}
	sb.Write([]byte(ip))
	sb.Write([]byte(":"))
	sb.Write([]byte(port))
	sb.Write([]byte("/data/batchWrite"))
	url := sb.String() // 请求的url
	buf := bytes.NewBuffer(requestBody)
	request, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return nil, fmt.Errorf("构造请求失败: %s", err)
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8") //添加请求头
	request.Header.Set("Authorization", `{"username": "admin", "password": "admin@123"}`)
	client := http.Client{} //创建客户端
	t0 := time.Now()
	resp, err := client.Do(request.WithContext(context.TODO())) //发送请求
	t1 := time.Now()
	fmt.Printf("[%s]: 第%d次写入耗时:%d ms", time.Now().Format(TimeFormatString), count, t1.Sub(t0).Milliseconds())
	if err != nil {
		return nil, fmt.Errorf("请求失败: %s", err)
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取response失败: %s", err)
	}
	return respBytes, nil
}
