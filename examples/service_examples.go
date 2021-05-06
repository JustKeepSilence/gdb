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
	"fmt"
	"github.com/JustKeepSilence/gdb/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

const (
	ip = "192.168.1.104:8082"
)

// mock client
func httpsClient() {
	if cred, err := credentials.NewClientTLSFromFile("./ssl/gdbServer.crt", "github.com"); err != nil {
		log.Fatal(err)
	} else {
		if conn, err := grpc.Dial(ip, grpc.WithTransportCredentials(cred)); err != nil {
			log.Fatal(err)
		} else {
			client := model.NewPageClient(conn)
			if r, err := client.UserLogin(context.Background(), &model.AuthInfo{
				UserName: "admin",
				PassWord: "685a6b21dc732a9702a96e6731811ec9",
			}); err != nil {
				log.Fatal(err)
			} else {
				fmt.Println(r.GetToken())
			}

		}
	}
}
