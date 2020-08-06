package work

import (
	"fmt"
	"github.com/Carrotxyy/syncSense/models"
)

/**
	上传同步 新增机构 数据

	返回参数 (错误)
 */
func (w *Work)AddOrginfoUpload()error{
	fmt.Println("新增机构同步开始")
	var orginfos []*models.Orginfo
	where := models.Orginfo{
		Org_SenseMark: "1",	// 表示 是新添加的数据
		Org_Status: "1",	// 表示 数据是有效的
	}
	// 获取数据
	err,count := w.Repository.Get(where,&orginfos,"Org_ID,Org_Name,Org_SenseMark,Org_SenseID","")

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

	err = w.AddOrginfoController(orginfos)
	fmt.Println("新增机构同步 结束")
	return err
}

/**
	将修改机构操作，同步数据到云
 */
func (w *Work)UpdateOrginfoUpload()error{
	fmt.Println("修改机构同步 开始")
	// 定义数据对象
	var orginfos []*models.Orginfo
	// 获取 添加操作 的数据   2 => 修改操作
	where := models.Orginfo{Org_SenseMark: "2",Org_Status: "1"}
	// 获取数据,  trim(Org_SenseID)!=''  表示Org_SenseID不为空串
	err,count := w.Repository.Get(&where,&orginfos,"Org_ID,Org_Name,Org_SenseID"," trim(Org_SenseID)!='' ")

	if err != nil {
		fmt.Println("获取数据错误!")
		return err
	}

	for _, value := range orginfos {
		fmt.Println("数据:",value)
	}
	if count <= 0 {
		fmt.Println("暂无数据!")
		return nil
	}
	err = w.UpdateOrginfoController(orginfos)
	fmt.Println("修改机构同步 结束")
	return err
}

/**
	将删除机构操作的数据，同步到sense
 */
func (w *Work)DeleteOrginfoUpload()error{
	fmt.Println("删除机构同步 开始")
	var orginfos []*models.Orginfo

	where := models.Orginfo{Org_SenseMark: "3",Org_Status: "0"}
	// 获取数据，需要删除的数据
	err,count := w.Repository.Get(where,&orginfos,"Org_ID,Org_Name,Org_SenseID","")
	if err != nil {
		fmt.Println("删除机构同步 => 获取数据错误:",err)
		return err
	}
	if count <= 0 {
		fmt.Println("暂无数据")
		return nil
	}

	err = w.DeleteOrginfoController(orginfos)
	fmt.Println("删除机构同步 结束")
	return err
}

/**
	有时候，当管理员新建一个机构后，在同步程序没有执行之前，
	修改了机构的数据，导致标志位Org_SenseMark="2",但是这个机构又不存在Org_SenseID
	所以会导致这个机构无法同步到商汤系统中，
	并且添加到这个机构下的人员也无法同步到商汤的系统中

	所以，这里要找出Org_SenseMark="2",但没有 Org_SenseID的数据，调用添加机构的方法
 */
func (w *Work)OtherOrginfoUpload()error{
	fmt.Println("特殊机构同步 开始")
	// 定义数据对象
	var orginfos []*models.Orginfo
	where := models.Orginfo{Org_SenseMark: "2",Org_Status: "1"}
	// trim(Org_SenseID)='' 表示 Org_SenseID 为 空串
	err,count := w.Repository.Get(&where,&orginfos,"Org_ID,Org_Name,Org_SenseID"," trim(Org_SenseID)='' ")
	if err != nil {
		fmt.Println("特殊机构同步，数据库数据获取错误:",err)
		return err
	}
	if count <= 0 {
		fmt.Println("特殊机构,暂无数据")
		return nil
	}

	err = w.AddOrginfoController(orginfos)
	fmt.Println("特殊机构同步 结束")
	return err
}

/**
	更新 新增数据 具体操作

	@orginfos 需要同步的数据

 */
func (w *Work)AddOrginfoController(orginfos []*models.Orginfo)error{

	// 构建访问路径
	path := w.SplicUrl("/cloudSync/sense/orginfos")

	// 数据类型
	contentType := "application/json"
	var res struct{
		Code int `json:"code"`
		Orginfos []models.Orginfo `json:"data"`
	}
	// 发起请求
	err := w.HttpApi.Post(path,orginfos,&res,contentType)

	if err != nil || res.Code != 200 {
		fmt.Println("上传机构数据错误:",err)
		return err
	}
	// 返回参数的Org_ID,用于恢复标志位
	var ids []int
	// 更新 Org_SenseID
	for _, item := range res.Orginfos {
		// 更新条件
		where := &models.Orginfo{Org_ID: item.Org_ID}
		// 更新内容
		newObj := &models.Orginfo{Org_SenseID: item.Org_SenseID}
		err = w.Repository.Update(models.Orginfo{},where,newObj)

		if err != nil {
			fmt.Println("更新机构 sense ID 错误：",err)
		}else{
			// 只有更新成功的才能恢复标志位
			ids = append(ids,item.Org_ID)
		}
	}
	//上传成功后，要将标志位恢复成 "0" 已同步状态
	err = w.Repository.UpdateMark(models.Orginfo{},"Org_ID",ids,models.Orginfo{Org_SenseMark: "0"})
	if err != nil {
		fmt.Println("更新标志位失败:",err)
		return err
	}

	fmt.Println("返回的参数:",res)
	return nil
}

/**
	更新 修改数据 具体操作

	@orginfos 需要同步的数据

*/
func (w *Work)UpdateOrginfoController(orginfos []*models.Orginfo)error{

	// 构建访问路径
	path := w.SplicUrl("/cloudSync/sense/orginfos")

	// 接受响应数据
	var res struct{
		Code int `json:"code"`
		Orginfos []models.Orginfo `json:"data"`
	}

	// 发起请求
	err := w.HttpApi.PUT(path,orginfos,&res)

	if err != nil {
		fmt.Println("请求失败:",err)
		return err
	}

	//上传成功后，要将标志位恢复成 0 已同步状态
	var ids []int
	for _, item := range res.Orginfos {
		ids = append(ids,item.Org_ID)
	}
	// 更新标志位
	err = w.Repository.UpdateMark(models.Orginfo{},"Org_ID",ids,&models.Orginfo{Org_SenseMark: "0"})

	if err != nil {
		fmt.Println("更新标志位失败:",err)
		return err
	}

	return nil
}

/**
	更新 删除数据 具体操作

	@orginfos 需要同步的数据

*/
func (w *Work)DeleteOrginfoController(orginfos []*models.Orginfo)error{
	// 构建路径
	path := w.SplicUrl("/cloudSync/sense/orginfos")

	var res struct{
		Code int `json:"code"`
		Orginfos []models.Orginfo `json:"data"`
	}

	err := w.HttpApi.Delete(path,orginfos,&res)
	if err != nil {
		fmt.Println("请求失败:",err)
		return err
	}
	//上传成功后，要将标志位恢复成 0 已同步状态
	var ids []int
	for _, item := range res.Orginfos {
		ids = append(ids,item.Org_ID)
	}
	// 更新标志位
	err = w.Repository.UpdateMark(models.Orginfo{},"Org_ID",ids,&models.Orginfo{Org_SenseMark: "0"})
	if err != nil {
		fmt.Println("更新标志位失败:",err)
		return err
	}

	return nil
}