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
	senseConfig := setting.Config{}
	senseConfig.WxAddr = "http://xyz.szlimaiyun.cn"
	senseConfig.DbType = "mysql"
	senseConfig.DbUser = "root"
	senseConfig.DbPassword = "123456"
	senseConfig.DbIP = "127.0.0.1"
	senseConfig.DbName = "gin-vue"
	senseConfig.TablePrefix = "go_"


	init,_ := Create(senseConfig)

	// 机构同步
	//// 同步Org_SenseMark = "1" 的新增数据
	//init.Work.AddOrginfoUpload()
	//// 同步Org_SenseMark = "2" 的修改数据
	//init.Work.UpdateOrginfoUpload()
	//// 同步Org_SenseMark = "2" 且 Org_SenseID = "" 特殊情况
	//init.Work.OtherOrginfoUpload()
	//// 同步Org_SenseMark = "3" 的删除数据
	//init.Work.DeleteOrginfoUpload()

	// 人员同步
	//// 同步Per_SenseMark = "1" 的新增数据
	//init.Work.AddPersonUpload()
	//// 同步Per_SenseMark = "2" 的修改数据
	//init.Work.UpdatePersonUpload()
	//// 同步Per_SenseMark = "2" 且 Per_SensePerID = "" 特殊情况
	//init.Work.OtherPersonUpload()
 	//// 同步Per_SenseMark = "3" 的删除数据
	//init.Work.DeletePersonUpload()

	// 访客同步
	init.Work.AddVisitorUpload()
}


func (init *Init)RunWork(){

	// 同步Org_SenseMark = "1" 的新增数据
	init.Work.AddOrginfoUpload()
	// 同步Org_SenseMark = "2" 的修改数据
	init.Work.UpdateOrginfoUpload()
	// 同步Org_SenseMark = "2" 且 Org_SenseID = "" 特殊情况
	init.Work.OtherOrginfoUpload()
	// 同步Org_SenseMark = "3" 的删除数据
	init.Work.DeleteOrginfoUpload()
}

// 获取配置对象
func GetConfig()setting.Config{
	return setting.Config{}
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