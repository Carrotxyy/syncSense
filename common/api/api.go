package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Api struct {

}

/**
	以post方式，向指定路径，发送指定类型的数据-- 添加数据使用

	@path 请求路径
	@data 数据
	@out 装载响应数据
	@contentType 请求体类型

	返回参数 (错误)
*/

func (a *Api)Post(path string , data , out interface{} , contentType string)error{
	// 将数据json化
	jsonStr, _ := json.Marshal(data)
	// 发送请求
	res, err := http.Post(path, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("发送请求错误:", err)
		return err
	}
	// 关闭响应流
	defer res.Body.Close()
	// 获取数据
	err = GetResponse(res,out)
	return err
}


/**
	以PUT方式，向指定路径，发送指定类型的数据-- 修改数据使用

	@path 请求路径
	@data 数据
	@out 装载响应数据

	返回参数 (错误)
*/

func (a *Api)PUT(path string , data , out interface{} )error{
	// 将数据json化
	jsonstr,_ := json.Marshal(data)
	fmt.Println(string(jsonstr))
	// 创建请求体
	req,err := http.NewRequest("PUT",path,bytes.NewBuffer(jsonstr))
	if err != nil {
		fmt.Println("创建请求体错误:",err)
		return err
	}
	// 设置数据类型
	req.Header.Set("Content-Type","application/json")
	// 发起请求
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("发起请求错误:",err)
		return err
	}
	// 延迟关闭响应流
	defer res.Body.Close()
	// 获取响应数据
	err = GetResponse(res,out)
	return err
}

func (a *Api)Delete(path string,data,out interface{})error{
	// 将数据json化
	jsonStr,_ := json.Marshal(data)
	// 构建请求体
	req,err := http.NewRequest("DELETE",path,bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("构建delete请求体错误:",err)
		return err
	}
	// 设置数据类型
	req.Header.Set("Content-Type","application/json")

	client := &http.Client{}
	// 发起请求
	res,err := client.Do(req)
	if err != nil {
		fmt.Println("请求错误:",err)
		return err
	}
	defer res.Body.Close()
	// 获取响应数据
	err = GetResponse(res,out)
	return err
}

// 获取响应数据
func GetResponse(res *http.Response,out interface{})error{
	// 获取响应流
	body,err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("读取数据错误:",err)
		return err
	}
	// 绑定数据
	err = json.Unmarshal(body,out)
	if err != nil {
		fmt.Println("json化数据错误:",err)
		return err
	}
	return nil
}