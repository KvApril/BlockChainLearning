package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"contract_load/contract"
	"fmt"
)

func main(){
	//链接geth客户端
	client,err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	//contractTest智能合约部署后的地址
	address := common.HexToAddress("0xe887722896D640B787D1E531BfB1d1765a327DEE")
	//获取合约对象
	instance,err := store.NewStore(address,client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is load")
	_ = instance
}
