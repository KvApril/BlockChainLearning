package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"context"
	"event_read/contract"
	"fmt"
)

func main()  {

	client,err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	//合约地址
	contractAddress := common.HexToAddress("0xe887722896D640B787D1E531BfB1d1765a327DEE")
	//设置查询范围
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(11515),
		ToBlock:   big.NewInt(11518),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	//创建log对象
	logs,err := client.FilterLogs(context.Background(),query)
	if err != nil {
		log.Fatal(err)
	}

	//合约abi
	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		log.Fatal("abi error: ",err)
	}

	for _,vLog := range logs  {
		//合约中的event事件,组装成结构体
		event := struct {
			Key          big.Int
			Value      	 common.Address
		}{}

		//对输出结果unpack
		err := contractAbi.Unpack(&event,"ItemSet",vLog.Data)
		if err != nil {
			log.Fatal("unpack error: ",err)
		}

		//打印结果
		fmt.Println(event.Key)
		fmt.Println(event.Value.Hex())
	}

}


