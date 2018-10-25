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

// TestBin is the compiled bytecode used for deploying new contracts.
const TestBin = `6060604052341561000f57600080fd5b604051610497380380610497833981016040528080519060200190919080518201919060200180519060200190919050506060604051908101604052808381526020018281526020018473ffffffffffffffffffffffffffffffffffffffff168152506000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000190805190602001906100cc929190610129565b506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050505050506101ce565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061016a57805160ff1916838001178555610198565b82800160010185558215610198579182015b8281111561019757825182559160200191906001019061017c565b5b5090506101a591906101a9565b5090565b6101cb91905b808211156101c75760008160009055506001016101af565b5090565b90565b6102ba806101dd6000396000f30060606040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063469e906714610051578063a3ed51061461015a575b600080fd5b341561005c57600080fd5b610088600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061019c565b60405180806020018481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281038252858181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156101495780601f1061011e57610100808354040283529160200191610149565b820191906000526020600020905b81548152906001019060200180831161012c57829003601f168201915b505094505050505060405180910390f35b341561016557600080fd5b61019a600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506101e5565b005b600060205280600052604060002060009150905080600001908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905083565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508273ffffffffffffffffffffffffffffffffffffffff168160020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610289578181600101819055505b5050505600a165627a7a72305820833cbae3cd045dab86e690c8f584a92354c770a86c0ce9ae8bcc14821e7cc9070029`

const keyStore  = `{"address":"575f174218e725a15f3e427fafc2fa844c6cfbec","crypto":{"cipher":"aes-128-ctr","ciphertext":"8eb38f492c9c1340c599bf1f40fbfaa5da663ca7b6b78f7bd0b6747597966f6c","cipherparams":{"iv":"f474a1c15ba676779542bc033174ecfb"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"10fc0a654e87d3b9659560a31f73e497775fc818c281788f6d0a150ce5d7b30c"},"mac":"e3f51e2963f2aaf9635f2af6daa4608329c0cb6cb6a38f22a19e8a3c80fff377"},"id":"130be720-33d3-483b-baa8-f6239a663a17","version":3}`

func main()  {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	rawDeployContract(client)
}

func rawDeployContract(client *ethclient.Client)  {
	byteCode := common.Hex2Bytes(TestBin)
	testAbi,_:=abi.JSON(strings.NewReader(TestABI))

	_addr := common.HexToAddress("0xc702445912337196d14b1311c10371c8a0643025") //参数1
	_name := "kv" //参数2
	_age := big.NewInt(23) ////参数3
	input,_ := testAbi.Pack("",_addr,_name,_age) //如果构造函数没有参数,直接testAbi.pack("")就行
	byteCode = append(byteCode,input...) //需要发送的data数据

	unlockedKey,_ := keystore.DecryptKey([]byte(keyStore),"123") //解锁私钥

	nonce,_:= client.NonceAt(context.Background(),unlockedKey.Address,nil) //获取私钥所属账户的nonce值

	tx := types.NewContractCreation(nonce,big.NewInt(0),big.NewInt(10000000),big.NewInt(10000000),byteCode) //创建交易

	signTx,_ := types.SignTx(tx,types.NewEIP155Signer(big.NewInt(15)),unlockedKey.PrivateKey) //使用私钥签名交易
	fmt.Println(signTx)

	err := client.SendTransaction(context.Background(),signTx)  //发送rawTransaction
	if err != nil {
		fmt.Println(err)
	}
}
