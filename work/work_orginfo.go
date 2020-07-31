package work

import (
	"fmt"
	"github.com/Carrotxyy/syncSense/common/api"
	"github.com/Carrotxyy/syncSense/models"
)

/**
	上传同步机构数据

	返回参数 (错误)
 */
func (w *Work)OrginfoUpload()error{
	var orginfos []*models.Orginfo

	// 0 是已同步 ，not 反向获取变动的数据
	not := models.Orginfo{
		Org_SenseMark: "0",
	}

	// 获取数据
	err,count := w.Repository.Get_Not(not,&orginfos,"Org_ID,Org_Name,Org_SenseMark,Org_SenseID")

	if err != nil {
		fmt.Println("获取机构数据错误!",err)
		return err
	}
	// 没有最新数据
	if count <= 0 {
		fmt.Println("暂无新数据!")
		return nil
	}
	for _, item := range orginfos {

		fmt.Println("机构数据:",item)
	}

	// 构建访问路径
	path := w.SplicUrl("/cloudSync/sense/orginfos")

	// 数据类型
	contentType := "application/json"
	var res struct{
		Code int `json:"code"`
		Orginfos []models.Orginfo `json:"data"`
	}
	// 发起请求
	err = api.Post(path,orginfos,&res,contentType)

	if err != nil {
		fmt.Println("上传机构数据错误:",err)
		return err
	}else{
		//上传成功后，要将标志位恢复成 0 已同步状态
		// ...
		// 更新 Org_SenseID
		for _, item := range res.Orginfos {
			err = w.Repository.Update(&models.Orginfo{},&models.Orginfo{Org_ID: item.Org_ID},&models.Orginfo{Org_SenseID: item.Org_SenseID})
			if err != nil {
				fmt.Println("更新机构 sense ID 错误：",err)
			}
		}
	}
	fmt.Println("返回的参数:",res)
	return nil
}