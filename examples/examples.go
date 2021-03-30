package examples

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/JustKeepSilence/gdb/db"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	ip   = "http://192.168.1.2"
	port = "9000"
)

func main() {
	//errorGroupTest()
	mockWritingData("./test.xlsx")
	//dbPath := "./db"         // path of data
	//itemDbPath := "./itemDb" // path of itemDb
	//g, err := db.NewGdb(dbPath, itemDbPath)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Open db successfully")
	////add groups:
	//groupInfos := []db.AddGroupInfo{{
	//	GroupName:   "1DCS",
	//	ColumnNames: []string{"groupName", "type", "description", "unit", "source"}, // every group has two cols: id and itemName
	//}}
	//if _, err := g.AddGroups(groupInfos...); err != nil {
	//	log.Fatal(err)
	//} else {
	//	fmt.Println("add group successfully")
	//}
	////add items
	//if _, err := g.AddItems(db.AddItemInfo{
	//	GroupName: "1DCS",
	//	Values: []map[string]string{{"itemName": "testItem1", "groupName": "1DCS", "type": "", "description": "testItem1", "unit": "", "source": ""},
	//		{"itemName": "testItem2", "type": "", "groupName": "1DCS", "description": "testItem2", "unit": "", "source": ""},
	//		{"itemName": "testItem3", "type": "", "groupName": "1DCS", "description": "testItem3", "unit": "", "source": ""},
	//		{"itemName": "testItem4", "type": "", "groupName": "1DCS", "description": "testItem4", "unit": "", "source": ""},
	//		{"itemName": "testItem5", "type": "", "groupName": "1DCS", "description": "testItem5", "unit": "", "source": ""},
	//		{"itemName": "testItem6", "type": "", "groupName": "1DCS", "description": "testItem6", "unit": "", "source": ""},
	//		{"itemName": "testItem7", "type": "", "groupName": "1DCS", "description": "testItem7", "unit": "", "source": ""},
	//		{"itemName": "testItem8", "type": "", "groupName": "1DCS", "description": "testItem8", "unit": "", "source": ""}},
	//}); err != nil {
	//	log.Fatal(err)
	//} else {
	//	fmt.Println("add items successfully!")
	//}
	////add items by excel
	//if _, err := g.AddItemsByExcel("1DCS", "./test.xlsx"); err != nil {
	//	log.Fatal(err)
	//} else {
	//	fmt.Println("add items by excel successfully!")
	//}
	////write realTime Data without timeStamp
	//if _, err := g.BatchWrite(db.BatchWriteString{
	//	GroupName: "1DCS",
	//	ItemValues: []db.ItemValue{{
	//		ItemName: "testItem1",
	//		Value:    "-100",
	//	}, {
	//		ItemName: "testItem2",
	//		Value:    "0",
	//	}, {
	//		ItemName: "testItem3",
	//		Value:    "100",
	//	}, {
	//		ItemName: "testItem4",
	//		Value:    "200",
	//	}, {
	//		ItemName: "testItem5",
	//		Value:    "300",
	//	}},
	//	WithTimeStamp: false,
	//}); err != nil {
	//	log.Fatal(err)
	//} else {
	//	fmt.Println("Write successfully")
	//}
	//// write realTime data without timeStamp
	//t := fmt.Sprintf("%d", time.Now().Add(-1*time.Hour).Unix()) // unix timeStamp
	//if _, err := g.BatchWrite(db.BatchWriteString{
	//	GroupName: "1DCS",
	//	ItemValues: []db.ItemValue{{
	//		ItemName:  "testItem6",
	//		Value:     "400",
	//		TimeStamp: t,
	//	}, {
	//		ItemName:  "testItem7",
	//		Value:     "500",
	//		TimeStamp: t,
	//	}, {
	//		ItemName:  "testItem8",
	//		Value:     "600",
	//		TimeStamp: t,
	//	}},
	//	WithTimeStamp: true,
	//}); err != nil {
	//	log.Fatal(err)
	//} else {
	//	fmt.Println("Write with timeStamp successfully")
	//}
	//// get realTime data, return the latest updated data
	//itemNames := []string{"testItem1", "testItem2", "testItem3", "testItem4", "testItem5", "testItem6", "testItem7", "testItem8"}
	//if c, err := g.GetRealTimeData(itemNames...); err != nil {
	//	log.Fatal(err)
	//} else {
	//	r, _ := json.Marshal(c)
	//	fmt.Println(fmt.Sprintf("%s", r))
	//}
	//if c, err := g.GetRawHistoricalData(itemNames...); err != nil {
	//	log.Fatal(err)
	//} else {
	//	r, _ := json.Marshal(c)
	//	fmt.Println(fmt.Sprintf("%s", r))
	//}
	//// get historical data with timeStamp
	//timeStamps := [][]int{{1612413561}, {1612413561}, {1612413561}, {1612413561}, {1612413561}}
	//if c, err := g.GetHistoricalDataWithStamp([]string{"testItem1", "testItem2", "testItem3", "testItem4", "testItem5"}, timeStamps...); err != nil {
	//	log.Fatal(err)
	//} else {
	//	r, _ := json.Marshal(c)
	//	fmt.Println(fmt.Sprintf("%s", r))
	//}
}

// mock writing data

func mockWritingData(filePath string) {
	contents, err := readExcel(filePath)
	if err != nil {
		fmt.Println(err)
		time.Sleep(60 * time.Second)
	}
	count := 1
	s := rand.NewSource(64)
	rr := rand.New(s)
	itemValues := []db.ItemValue{}
	for index, itemName := range contents {
		if itemName == "JL1_10DAS27A:HAG46AA101ZF.CIN          " {
			fmt.Println(index)
		}
		itemValues = append(itemValues, db.ItemValue{
			ItemName: strings.Trim(itemName, " "),
			Value:    strconv.Itoa(rr.Intn(1000)),
		})
	}
	for {
		ct := time.Now().Second()
		if ct == 10 || ct == 11 || ct == 12 || ct == 13 || ct == 20 || ct == 21 || ct == 22 || ct == 23 {
			content := map[string]interface{}{"itemValues": itemValues, "groupName": "1DCS"}
			requestData, _ := json.Marshal(content)
			r, err := SendPost(requestData, count)
			fmt.Printf(", wirting result: %s, v: %s\n", fmt.Sprintf("%s", r), itemValues[0].Value)
			if err != nil {
				log.Fatal(err)
			}
			count++
			time.Sleep(time.Second)
		} else {
			values := []db.ItemValue{}
			for _, itemName := range contents {
				values = append(values, db.ItemValue{
					ItemName: strings.Trim(itemName, " "),
					Value:    strconv.Itoa(rr.Intn(1000)),
				})
			}
			content := map[string]interface{}{"itemValues": values, "groupName": "1DCS"}
			requestData, _ := json.Marshal(content)
			r, err := SendPost(requestData, count)
			fmt.Printf(", writing result: %s, v: %s\n", fmt.Sprintf("%s", r), values[0].Value)
			if err != nil {
				log.Fatal(err)
			}
			count++
			time.Sleep(time.Second)
		}
	}
}

func readExcel(filePath string) ([]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	rows, err := f.Rows("Sheet1")
	if err != nil {
		return nil, err
	}
	insertedData := []string{}
	for rows.Next() {
		row, err := rows.Columns()
		insertedData = append(insertedData, row[0])
		if err != nil {
			return nil, err
		}

	}
	return insertedData[1:], nil
}

func SendPost(requestBody []byte, count int) ([]byte, error) {
	sb := strings.Builder{}
	sb.Write([]byte(ip))
	sb.Write([]byte(":"))
	sb.Write([]byte(port))
	sb.Write([]byte("/data/batchWrite"))
	url := sb.String()
	buf := bytes.NewBuffer(requestBody)
	request, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return nil, fmt.Errorf("fail in constructing request: %s", err)
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", `{"username": "admin", "password": "admin@123"}`)
	client := http.Client{}
	t0 := time.Now()
	resp, err := client.Do(request.WithContext(context.TODO()))
	t1 := time.Now()
	fmt.Printf("[%s]: %d wirting consuming :%d ms", time.Now().Format("2006-01-02 15:04:05"), count, t1.Sub(t0).Milliseconds())
	if err != nil {
		return nil, fmt.Errorf("fail in requesting: %s", err)
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("fail in reading response: %s", err)
	}
	return respBytes, nil
}
