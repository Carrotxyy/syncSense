package work

import (
	"fmt"
	"github.com/Carrotxyy/syncSense/models"
)

/**
上传 添加人员 数据
*/
func (w *Work) AddPersonUpload() {
	fmt.Println("新增人员同步 开始")

	var persons []*models.Personinfo

	where := models.Personinfo{
		Per_SenseMark: "1", // 新增数据
		Per_Status:    "1", // 有效数据
	}
	// 获取数据
	err, count := w.Repository.GetReferBelongsTo(where, &persons, []string{"Per_OrgRefer"}, "", "`Per_Image` != \"\" ")
	if err != nil {
		fmt.Println("获取新增人员数据错误:", err)
		return
	}
	if count <= 0 {
		fmt.Println("暂无新增数据!")
		return
	}

	for _, item := range persons {
		fmt.Println("数据:", item)
	}
	w.AddPersonController(persons)
	fmt.Println("新增人员 结束")
}

func (w *Work) AddPersonController(persons []*models.Personinfo) error {
	// 构建访问路径
	path := w.SplicUrl("/cloudSync/sense/persons")

	// 数据类型
	contentType := "application/json"
	var res struct {
		Code    int                 `json:"code"`
		Persons []models.Personinfo `json:"data"`
	}
	// 发起请求
	err := w.HttpApi.Post(path, persons, &res, contentType)

	if err != nil || res.Code != 200 {
		fmt.Println("上传机构数据错误:", err)
		return err
	}

	var ids []int
	// 保存senseid
	for _, item := range res.Persons {
		// 更新条件
		where := models.Personinfo{Per_ID: item.Per_ID}
		// 更新内容
		newObj := models.Personinfo{Per_SensePerID: item.Per_SensePerID}
		err = w.Repository.Update(models.Personinfo{}, where, newObj)

		if err != nil {
			fmt.Println("更新人员 senseid 错误：", err)
		} else {
			// 只有更新成功的才能将标志位恢复
			ids = append(ids, item.Per_ID)
		}
	}

	//上传成功后，要将标志位恢复成 0 已同步状态
	// 更新标志位
	err = w.Repository.UpdateMark(models.Personinfo{}, "Per_ID", ids, &models.Personinfo{Per_SenseMark: "0"})

	if err != nil {
		fmt.Println("更新标志位失败:", err)
		return err
	}

	return nil
}

/**
上传 修改人员 数据
*/
func (w *Work) UpdatePersonUpload() {
	fmt.Println("修改人员 开始")
	var persons []*models.Personinfo

	where := models.Personinfo{
		Per_Status:    "1", // 数据有效
		Per_SenseMark: "2", // 修改操作
	}
	// 获取数据
	err, count := w.Repository.GetReferBelongsTo(where, &persons, []string{"Per_OrgRefer"}, ""," trim(Per_SensePerID)!='' ")
	if err != nil {
		fmt.Println("获取 修改人员 数据错误!", err)
		return
	}
	if count <= 0 {
		fmt.Println("暂无数据")
		return
	}

	w.UpdatePersonController(persons)
	fmt.Println("修改人员 结束")
}

func (w *Work) UpdatePersonController(persons []*models.Personinfo) error {
	// 构建访问路径
	path := w.SplicUrl("/cloudSync/sense/persons")

	var res struct {
		Code int                 `json:"code"`
		Data []models.Personinfo `json:"data"`
	}

	err := w.HttpApi.PUT(path, persons, &res)
	//
	if err != nil {
		fmt.Println("请求出现错误：", err)
		return err
	}

	//上传成功后，要将标志位恢复成 0 已同步状态
	var ids []int
	for _, item := range res.Data {
		ids = append(ids, item.Per_ID)
	}
	// 更新标志位
	err = w.Repository.UpdateMark(models.Personinfo{}, "Per_ID", ids, &models.Personinfo{Per_SenseMark: "0"})

	if err != nil {
		fmt.Println("更新标志位失败:", err)
		return err
	}

	return nil
}

/**
上传 删除人员 数据
*/
func (w *Work) DeletePersonUpload() {
	fmt.Println("删除人员同步 开始")

	var persons []*models.Personinfo
	where := models.Personinfo{
		Per_SenseMark: "3",
		Per_Status:    "0",
	}
	// 获取需要删除的数据
	err, count := w.Repository.Get(where, &persons, "Per_ID,Per_SensePerID", "")
	if err != nil {
		fmt.Println("删除人员同步 => 获取数据错误:", err)
		return
	}
	if count <= 0 {
		fmt.Println("暂无数据")
		return
	}
	for _, item := range persons {
		fmt.Println("数据:",item)
	}
	w.DeletePersonController(persons)
}

func (w *Work) DeletePersonController(persons []*models.Personinfo) error {

	// 构建路径
	path := w.SplicUrl("/cloudSync/sense/persons")

	var res struct {
		Code int   `json:"code"`
		Data []int `json:"data"`
	}

	err := w.HttpApi.Delete(path, persons, &res)
	if err != nil {
		fmt.Println("删除人员，请求远程失败:", err)
		return err
	}

	//上传成功后，要将标志位恢复成 0 已同步状态
	// 更新标志位
	err = w.Repository.UpdateMark(models.Personinfo{}, "Per_ID", res.Data, &models.Personinfo{Per_SenseMark: "0"})
	if err != nil {
		fmt.Println("更新标志位失败:", err)
		return err
	}

	return nil
}

/**

有时候，当管理员新建一个人员后，在同步程序没有执行之前，
修改了人员的数据，导致标志位Per_SenseMark="2",但是这个人员又不存在Per_SensePerID
所以会导致这个机构无法同步到商汤系统中，

所以，这里要找出Per_SenseMark="2",但没有 Per_SensePerID的数据，调用添加人员的方法
 */

func (w *Work) OtherPersonUpload(){
	fmt.Println("特殊情况同步开始")
	var persons []*models.Personinfo

	where := models.Personinfo{
		Per_SenseMark: "2",
		Per_Status: "1" ,
	}

	err, count := w.Repository.GetReferBelongsTo(where, &persons, []string{"Per_OrgRefer"}, "", " trim(Per_SensePerID)='' ")

	if err != nil {
		fmt.Println("特殊人员情况，获取数据库数据错误:",err)
		return
	}

	if count <= 0 {
		fmt.Println("特殊人员情况，暂无数据")
		return
	}

	w.AddPersonController(persons)
}