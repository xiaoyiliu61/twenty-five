package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
    /**
    1\准备要进行rpc通常时的json数据
    2、使用http链接的post请求、发送json数据
    3、接收http响应
    4、根据响应的结果，进行判断和处理
      状态码200正常
      非200不正常
     */
    //1、准备要发送的json数据
    /**
    {
      "id":编号
       "method":"功能方法或者命令"，
       “jsonrpc”：“rpc版本2.0”
       “params”：携带的参数，数组形式
      }
     */
     //json操作：序列化 、反序列化
	/* rpcReq:= entity.RPCRequest{Id: time.Now().UnixNano(),Jsonrpc:"2.0",Method: GETBLOCKCOUNT}
	 //对结构体类型进行序列化
	 rpcBytes,err:=json.Marshal(&rpcReq)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
    fmt.Println("准备好的json格式的数据",string(rpcBytes))*/
	 //2.发送一个post类型的请求
	 //client :客户端，客户端用于发送请求
	 count:=GetBlockCount()
	 fmt.Println("得到的区块总数是：",count)

	 difficult:=GetDifficulty()
	 fmt.Println("获取区块的难度是：",difficult)

	 bestHash:=GetBestBlockHash()
	 fmt.Println("第一个区块的hash值：",bestHash)

	 blockChainInfo:=GetBlockChainInfo()
	 fmt.Println("获取当前区块的信息：",blockChainInfo)

	 blockHash:=GetBlockHash(2)
	 fmt.Println("获取指定高度的区块hash值：",blockHash)

	 newAddress:=GetNewAddress("ly","legacy")
	 fmt.Println("获取新的地址：",newAddress)
}



