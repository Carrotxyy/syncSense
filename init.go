package main

import (
	"fmt"
	"github.com/Carrotxyy/syncSense/common/api"
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

	// 同步Org_SenseMark = "1" 的新增数据
	init.Work.AddOrginfoUpload()
	// 同步Org_SenseMark = "2" 的修改数据
	init.Work.UpdateOrginfoUpload()
	// 同步Org_SenseMark = "2" 且 Org_SenseID = "" 特殊情况
	init.Work.OtherOrginfoUpload()
	// 同步Org_SenseMark = "3" 的删除数据
	init.Work.DeleteOrginfoUpload()
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
		&inject.Object{Value: &api.Api{}},
	); err != nil {
		log.Fatal("inject fatal: ", err)
	}
	if err := injector.Populate(); err != nil {
		log.Fatal("injector fatal: ", err)
	}

	return &init,nil
}