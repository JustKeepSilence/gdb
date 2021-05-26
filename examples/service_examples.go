/*
creatTime: 2021/3/30
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.15.3
*/

// examples about how to use gRCP of gdb in go, for more details about api examples or gRPC in nodeJS
// you can see https://github.com/JustKeepSilence/gdbUI/blob/master/src/renderer/api/index.js

package examples

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/JustKeepSilence/gdb/model"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"math"
	"math/rand"
	"time"
)

const (
	ip = "192.168.0.199:8082"
)

func httpsClient() (*grpc.ClientConn, error) {
	if cred, err := credentials.NewClientTLSFromFile("./ssl/gdbServer.crt", "gdb.com"); err != nil {
		return nil, err
	} else {
		if conn, err := grpc.Dial(ip, grpc.WithTransportCredentials(cred)); err != nil {
			return nil, err
		} else {
			return conn, nil
		}
	}
}

func httpClient() (*grpc.ClientConn, error) {
	if conn, err := grpc.Dial(ip, grpc.WithInsecure()); err != nil {
		return nil, err
	} else {
		return conn, err
	}
}

func mockWrite(mode string) {
	var conn *grpc.ClientConn
	var err error
	if mode == "http" {
		conn, err = httpClient()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		conn, err = httpsClient()
		if err != nil {
			log.Fatal(err)
		}
	}
	client := model.NewDataClient(conn)
	itemValues, _ := json.Marshal([]map[string]interface{}{{
		"groupName": "2DCS",
		"itemName":  "item1",
		"value":     2.0,
	}, {
		"groupName": "calc",
		"itemName":  "item1",
		"value":     1.0,
	}})
	if r, err := client.BatchWrite(context.Background(), &model.BatchWriteString{ItemValues: string(itemValues)}); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(r.GetEffectedRows())
	}
}

func mockWriteWithStream(mode string) {
	var conn *grpc.ClientConn
	var err error
	if mode == "http" {
		conn, err = httpClient()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		conn, err = httpsClient()
		if err != nil {
			log.Fatal(err)
		}
	}
	client := model.NewDataClient(conn)
	if stream, err := client.BatchWriteWithStream(context.Background()); err != nil {
		log.Fatal(err)
	} else {
		r := rand.New(rand.NewSource(99))
		itemValue, _ := json.Marshal([]map[string]interface{}{{"itemName": "item1", "groupName": "2DCS", "value": float64(r.Intn(3600)) * math.Pi},
			{"groupName": "calc", "itemName": "item1", "value": float64(r.Intn(3600)) * math.E}})
		if err := stream.Send(&model.BatchWriteString{ItemValues: string(itemValue)}); err != nil {
			log.Fatal(err)
		}
		if r, err := stream.CloseAndRecv(); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.GetEffectedRows())
		}
	}
}

func mockWriteHistoricalData(mode string) {
	// mock write historical data
	durations := 24 * 3600 // day ago
	v1, v2, ts1, ts2 := make([]interface{}, durations), make([]interface{}, durations), make([]int, durations), make([]int, durations)
	endTime := time.Now()
	st := time.Now()
	g := errgroup.Group{}
	for i := 0; i < durations; i++ {
		index := i
		g.Go(func() error {
			r := rand.New(rand.NewSource(99))
			duration := time.Second * time.Duration(-1*index)
			startTime := endTime.Add(duration)
			v1[index] = float64(r.Intn(index+1)) * math.Pi
			v2[index] = float64(r.Intn(index+1)) * math.E
			ts1[index] = int(startTime.Unix()) + 8*3600
			ts2[index] = int(startTime.Unix()) + 8*3600
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	} else {
		et := time.Now()
		fmt.Println(et.Sub(st).Seconds())
		itemValues, _ := json.Marshal([]map[string]interface{}{{"groupName": "2DCS", "itemName": "item1", "values": v1, "timeStamps": ts1},
			{"groupName": "calc", "itemName": "item1", "values": v2, "timeStamps": ts2}})
		var conn *grpc.ClientConn
		var err error
		if mode == "http" {
			conn, err = httpClient()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			conn, err = httpsClient()
			if err != nil {
				log.Fatal(err)
			}
		}
		client := model.NewDataClient(conn)
		if r, err := client.BatchWriteHistoricalData(context.Background(), &model.BatchWriteHistoricalString{HistoricalItemValues: string(itemValues)}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.String())
		}
	}
}

func mockWriteHistoricalDataWithStream(mode string) {
	var conn *grpc.ClientConn
	var err error
	if mode == "http" {
		conn, err = httpClient()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		conn, err = httpsClient()
		if err != nil {
			log.Fatal(err)
		}
	}
	client := model.NewDataClient(conn)
	if stream, err := client.BatchWriteHistoricalDataWithStream(context.Background()); err != nil {
		log.Fatal(err)
	} else {
		durations := 24 * 3600
		endTime := time.Now()
		for i := 0; i < durations; i++ {
			r := rand.New(rand.NewSource(99))
			duration := time.Second * time.Duration(-1*i)
			startTime := endTime.Add(duration)
			v1 := []interface{}{float64(r.Intn(i+1)) * math.Pi * -1}
			v2 := []interface{}{float64(r.Intn(i+1)) * math.E * -1}
			ts1 := []int{int(startTime.Unix()) + 8*3600}
			ts2 := []int{int(startTime.Unix()) + 8*3600}
			itemValue, _ := json.Marshal([]map[string]interface{}{{"groupName": "2DCS", "itemName": "item1", "values": v1, "timeStamps": ts1},
				{"groupName": "calc", "itemName": "item1", "values": v2, "timeStamps": ts2}})
			if err := stream.Send(&model.BatchWriteHistoricalString{HistoricalItemValues: string(itemValue)}); err != nil {
				log.Fatal(err)
			}
		}
		if r, err := stream.CloseAndRecv(); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r.String())
		}
	}
}
