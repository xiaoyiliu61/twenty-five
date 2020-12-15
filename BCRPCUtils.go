package main

import (
	"BitCoinConnect1214/entity"
	"BitCoinConnect1214/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)


func PrepareJson(constant string,params...interface{}) []byte{
	rpcReq:= entity.RPCRequest{Id: time.Now().UnixNano(),Jsonrpc:"2.0",Method:constant,Params:params}
	//对结构体类型进行序列化
	rpcBytes,err:=json.Marshal(&rpcReq)
	if err != nil {
		fmt.Println(err.Error())
	}
	return rpcBytes
}

func SendPost(rpcBytes []byte)interface{}{
	var result entity.RPCResult
	client:=http.Client{}//实例化一个客户端
	//实例化一个请求
	request, err := http.NewRequest("POST",RPCURL,bytes.NewBuffer([]byte(rpcBytes)))//建立一个缓冲区，进行读操作和写操作
	if err != nil {
		fmt.Println(err.Error())

	}
	//给post请求添加请求头设置信息
	//key --> value
	request.Header.Add("Encoding","UTF-8")
	request.Header.Add("Content-Type","application/json")
	//权限认证设置
	request.Header.Add("Authorization","Basic "+utils.Base64Str( RPCUSER+":"+RPCPASSWORD))

	//使用客户端发送请求
	response,err:=client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	//通过response，获取响应的数据
	code := response.StatusCode
	//fmt.Println(code)
	if code == 200{
		//fmt.Println("请求成功")
		resultBytes,err:=ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
	//	fmt.Println("rpc的调用响应结果：",string(resultBytes))
		//json的反序列化

		err=json.Unmarshal(resultBytes,&result)
		if err != nil {
			fmt.Println(err.Error())
		}
		//反序列化正常，没有出现错误
		//fmt.Println("功能调用结果：",result.Result)
		return  result.Result
	    }else {
		fmt.Println("请求失败")
	}
	return  nil
}