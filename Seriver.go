package main

func GetBlockCount()interface{}{
	rpcBytes:=PrepareJson(GETBLOCKCOUNT)
	return SendPost(rpcBytes)
}

func GetDifficulty() interface{}  {
	rpcBytes:=PrepareJson(GETDIFFICULTY)
	return SendPost(rpcBytes)
}

func GetBestBlockHash() interface{} {
	rpcBytes:=PrepareJson(GETBESTBLOCKHASH)
	return SendPost(rpcBytes)
}

func GetBlockChainInfo() interface{} {
	rpcBytes:=PrepareJson(GETBLOCKCHAININFO)
	return SendPost(rpcBytes)
}

func GetBlockHash(height int) interface{} {
	rpcBytes:=PrepareJson(GETBLOCKHASH,height)
	return SendPost(rpcBytes)
}

func GetNewAddress(label string,address_type string)interface{}{
	rpcBytes:=PrepareJson(GETNEWADDRESS,label,address_type)
	return SendPost(rpcBytes)
}