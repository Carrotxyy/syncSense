package main

import (
	"fmt"
	"github.com/Carrotxyy/syncSense/common/db"
	"github.com/Carrotxyy/syncSense/common/setting"
	"github.com/Carrotxyy/syncSense/work"
	"github.com/facebookgo/inject"
	"log"
)

type Init struct {
	Work *work.Work `inject:""`
}

func main() {
	wxConfig := setting.Config{}
	wxConfig.WxAddr = "http://xyz.szlimaiyun.cn"
	wxConfig.DbType = "mysql"
	wxConfig.DbUser = "root"
	wxConfig.DbPassword = "123456"
	wxConfig.DbIP = "127.0.0.1"
	wxConfig.DbName = "gin-vue"
	wxConfig.TablePrefix = "go_"


	init,_ := Create(wxConfig)


	init.Work.OrginfoUpload()
}


func Create(config setting.Config)(*Init,error){

	conn := db.DB{}
	err := conn.Connect(config)
	if err != nil {
		fmt.Println("连接数据库错误:",err)
		return nil,err
	}

	var init Init

	// 依赖注入
	var injector inject.Graph
	if err := injector.Provide(
		&inject.Object{Value: &init},
		&inject.Object{Value: &conn},
		&inject.Object{Value: &config},
	); err != nil {
		log.Fatal("inject fatal: ", err)
	}
	if err := injector.Populate(); err != nil {
		log.Fatal("injector fatal: ", err)
	}

	return &init,nil
}