package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
	以post方式，向指定路径，发送指定类型的数据

	@path 请求路径
	@data 数据
	@out 装载响应数据
	@contentType 请求体类型

	返回参数 (错误)
*/

func Post(path string , data , out interface{} , contentType string)error{
	// 将数据json化
	jsonStr, _ := json.Marshal(data)
	// 发送请求
	resp, err := http.Post(path, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("发送请求错误:", err)
		return err
	}
	// 关闭响应流
	defer resp.Body.Close()
	// 读取响应
	result, _ := ioutil.ReadAll(resp.Body)
	// 解析响应数据
	err = json.Unmarshal(result, out)
	if err != nil {
		fmt.Println("解析响应数据错误!", err)
		return err
	}
	return nil
}

