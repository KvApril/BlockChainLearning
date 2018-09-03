package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"event_read/contract"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/net/context"
)

func main()  {
	//使用websocket的形式链接客户端
	//注意启动geth的时候开启ws: -ws --wsaddr "0.0.0.0" --wsorigins "*" --wsport "8546" --wsapi "db,eth,net,web3,miner,personal,admin,shh" --shh
	//注意最后的--shh一定要加上,否则会连接有问题
	client,err := ethclient.Dial("ws://localhost:8546")
	if err != nil {
		log.Fatal(err)
	}
	//合约地址
	contractAddress := common.HexToAddress("0xe887722896D640B787D1E531BfB1d1765a327DEE")
	//设置查询条件
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			contractAddress,
		},
	}

	//创建一个通道
	logs := make(chan types.Log)
	sub,err := client.SubscribeFilterLogs(context.Background(),query,logs)
	if err != nil {
		log.Fatal("sub err: ",err)
	}

	//合约abi
	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		log.Fatal("abi error: ",err)
	}

	for{
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			//合约中的event事件,组装成结构体
			event := struct {
				Key          big.Int
				Value   	 common.Address
			}{}

			err := contractAbi.Unpack(&event,"SetUserAddress",vLog.Data)
			if err != nil {
				log.Fatal("unpack error: ",err)
			}

			//打印结果
			fmt.Println(event.Key)
			fmt.Println(event.Value.Hex())
		}
	}

}


