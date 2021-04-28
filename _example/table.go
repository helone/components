package main

import (
	"fmt"
	"github.com/helone/components/table"
)

type User struct {
	UID      string `json:"uid,index"`
	Nickname string `json:"nickname"`
}

func main() {
	c := table.New(&table.Config{
		EndPoint:        "",
		InstanceName:    "",
		AccessKeyId:     "",
		AccessKeySecret: "",
	})
	user := User{
		UID:      "",
		Nickname: "",
	}
	err := c.Insert(user)
	fmt.Println(err)
}
