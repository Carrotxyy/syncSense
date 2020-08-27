package work

import (
	"fmt"
	"github.com/Carrotxyy/syncSense/models"
)

/**
	上传 访客数据
 */
func (w *Work) AddVisitorUpload(){
	// 获取数据
	err,visitors  := w.Repository.GetUploadVisitor()
	if err != nil {
		fmt.Println("同步商汤系统 => 上传访客系统，获取数据错误 =>",err)
		return
	}
	if len(visitors) <= 0 {
		fmt.Println("同步商汤系统 => 上传访客系统，暂无数据")
		return
	}

	for _, item := range visitors {
		fmt.Println("数据:",item)
	}


	w.AddVisitorController(visitors)
}

func (w *Work) AddVisitorController(visitors []*models.Visitor){
	// 构建访问路径
	path := w.SplicUrl("/cloudSync/sense/visitor")

	// 数据类型
	contentType := "application/json"
	var res struct{
		Code int `json:"code"`
		Visitors []models.Visitor `json:"data"`
	}
	// 发起请求
	err := w.HttpApi.Post(path,visitors,&res,contentType)
	if err != nil {
		fmt.Println("上传访客数据错误")
		return
	}
	// 获取添加成功的访客id,用于恢复标志位
	var ids []int
	// 设置 Vis_SenseVisID
	for _, item := range res.Visitors {
		// 更新条件
		where := &models.Visitor{Vis_ID: item.Vis_ID}
		// 更新内容
		newObj := &models.Visitor{Vis_SenseVisID: item.Vis_SenseVisID}
		err = w.Repository.Update(models.Visitor{},where,newObj)

		if err != nil {
			fmt.Println("更新访客 sense ID 错误：",err)
		}else{
			// 只有更新成功的才能恢复标志位
			ids = append(ids,item.Vis_ID)
		}
	}
	// 恢复标志位
	err = w.Repository.UpdateMark(models.Visitor{},"Vis_ID",ids,models.Visitor{Vis_SenseMark: "0"})
	if err != nil {
		fmt.Println("更新标志位失败:",err)
	}
	fmt.Println("上传访客返回的参数:",res)
}
