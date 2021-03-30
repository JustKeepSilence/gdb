/*
creatTime: 2021/3/30
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.15.3
*/

package examples

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/JustKeepSilence/gdb/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

// mock client
func client() {
	if certificate, err := tls.LoadX509KeyPair("./ssl/gdbClient.crt", "./ssl/gdbClient.key"); err != nil {
		log.Fatal(err)
	} else {
		cred := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: false,
			ServerName:         "gdb.dev",
			Certificates:       []tls.Certificate{certificate},
		})
		if conn, err := grpc.Dial("localhost:8082", grpc.WithTransportCredentials(cred)); err != nil {
			log.Fatal(err)
		} else {
			defer conn.Close()
			c := model.NewGroupClient(conn)
			if r, err := c.GetGroupProperty(context.Background(), &model.QueryGroupPropertyInfo{
				GroupName: "1DCS",
				Condition: "1=1",
			}); err != nil {
				log.Fatal(err)
			} else {
				fmt.Println(r.ItemColumnNames)
			}
		}
	}
}
