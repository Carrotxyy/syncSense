package work

import (
	"fmt"
	"github.com/Carrotxyy/syncSense/models"
)

/**
	上传 访客数据
 */
func (w *Work) AddVisitorUpload(){
	var visitors []*models.Visitor
	where := models.Visitor{
		Vis_SenseMark: "1",	// 新增的访客数据
		Vis_State : "1",	// 业主同意访问
	}

	//err,count := w.Repository.GetReferBelongsTo(where,&visitors,[]string{"Vis_PersonRefer"},"","")
	err,count := w.Repository.GetUploadVisitor(where,&visitors)
	if err != nil {
		fmt.Println("同步商汤系统 => 上传访客系统，获取数据错误 =>",err)
		return
	}
	if count <= 0 {
		fmt.Println("同步商汤系统 => 上传访客系统，暂无数据")
		return
	}

	for _, item := range visitors {
		fmt.Println("数据:",item)
	}
}


