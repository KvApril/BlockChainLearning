package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"contract_read/contract"
	"fmt"
	"math/big"
)

func main()  {
	//链接geth客户端
	client,err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	//contractTest合约地址
	address := common.HexToAddress("0xe887722896D640B787D1E531BfB1d1765a327DEE")
	//contractTest合约对象
	instance,err := store.NewStore(address,client)
	if err != nil {
		log.Fatal(err)
	}

	//调用合约中的接口: StringItem,该接口没有参数
	//接口返回的数据:  合约部署时,传递的_item参数
	//其他查询类的接口,都可以这样调用,比如ERC20中的TotalSupply(),BalanceOf(address)等
	stringItem,err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stringItem)


	//调用需要传递参数的接口,无需auth
	userAddress,err := instance.GetItem(nil,big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("userAddress: ",userAddress.Hex())
}
