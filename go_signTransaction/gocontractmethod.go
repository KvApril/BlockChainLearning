package main

import (
	"log"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"fmt"
)

// TestABI is the input ABI used to generate the binding from.
const TestABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"records\",\"outputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"age\",\"type\":\"uint256\"},{\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_age\",\"type\":\"uint256\"}],\"name\":\"updateAge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_age\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

const keyStore  = `{"address":"575f174218e725a15f3e427fafc2fa844c6cfbec","crypto":{"cipher":"aes-128-ctr","ciphertext":"8eb38f492c9c1340c599bf1f40fbfaa5da663ca7b6b78f7bd0b6747597966f6c","cipherparams":{"iv":"f474a1c15ba676779542bc033174ecfb"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"10fc0a654e87d3b9659560a31f73e497775fc818c281788f6d0a150ce5d7b30c"},"mac":"e3f51e2963f2aaf9635f2af6daa4608329c0cb6cb6a38f22a19e8a3c80fff377"},"id":"130be720-33d3-483b-baa8-f6239a663a17","version":3}`

func main()  {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	rawContractTransaction(client)
}

func rawContractTransaction(client *ethclient.Client)  {
	contractAddress := common.HexToAddress("0xb50b5a3202cb435915bde5b3335679f18cb72f9e")
	testAbi,_:=abi.JSON(strings.NewReader(TestABI))

	unlockedKey,_ := keystore.DecryptKey([]byte(keyStore),"123") //解锁私钥
	nonce,_:= client.NonceAt(context.Background(),unlockedKey.Address,nil) //获取私钥所属账户的nonce值

	_addr := common.HexToAddress("0xc702445912337196d14b1311c10371c8a0643025")
	_age  := big.NewInt(25)
	bytesData,_ := testAbi.Pack("updateAge",_addr,_age) //调用合约中的方法
	fmt.Println("bytesData: ",bytesData)

	tx := types.NewTransaction(nonce,contractAddress,big.NewInt(0),big.NewInt(10000000),big.NewInt(10000000),bytesData) //创建交易

	signTx,_ := types.SignTx(tx,types.NewEIP155Signer(big.NewInt(15)),unlockedKey.PrivateKey) //使用私钥签名交易
	fmt.Println(signTx)

	err := client.SendTransaction(context.Background(),signTx)  //发送rawTransaction
	if err != nil {
		fmt.Println(err)
	}
}
