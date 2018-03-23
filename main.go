package main

import (
	"../MongoData"
	"proto/adventure"
	"./src"
	"github.com/micro/go-micro"
	"github.com/tangxuyao/Frameworks"
	"fmt"
)

func main() {
	mg := &mongo.MongoDB{}
	if err := mg.Dial(""); err != nil {
		panic("连接[mongo]数据失败，请检查相关参数")
	}

	service := micro.NewService(micro.Name("AdventureService"))
	service.Init()

	defer func() {
		service.Server().Deregister()
	}()

	handler := &adventure.AdventureService{M: mg}
	adventrue_api.RegisterAdventureServiceHandler(service.Server(), handler)

	fmt.Println("AdventureService starting...")

	d := Frameworks.NewDispatcher(10)
	d.Run()

	if err := service.Run(); err != nil {
		panic(err)
	}
}
