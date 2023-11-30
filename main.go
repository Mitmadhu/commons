package main

import (
	"fmt"

	"github.com/Mitmadhu/commons/clients"
	mysqlDto "github.com/Mitmadhu/mysqlDB/dto"
)

func main() {
	obj := clients.MysqlClientImpl{}
	resp, err := obj.Login("http://localhost:8081", "POST", mysqlDto.ValidateUserRequest{
		Username: "ayush",
		Password: "123",
		BaseRequest: mysqlDto.BaseRequest{
			MsgID: "123",
		},
	})
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("%+v", resp)
}
