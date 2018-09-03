package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/ethereum/go-ethereum/crypto"
	"fmt"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"golang.org/x/net/context"

	"contractGo/contract"
)

func main(){
	//链接geth客户端
	client,err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	//账户私钥
	privateKey,err := crypto.HexToECDSA("ba8efd14a6bf98f315f51ab957f3efb9fc8fcf82e4551839d207f4e8a97a8406")
	if err != nil{
		log.Fatal(err)
	}

	//计算公钥
	publicKey := privateKey.Public()
	publicKeyECDSA,ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key ECDSA")
	}

	//计算发送的地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("fromAddress=",fromAddress)

	//nonce值
	nonce,err := client.PendingNonceAt(context.Background(),fromAddress)
	fmt.Println("nonce=",nonce)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice,err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//交易相关参数
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	//合约参数
	stringItem := "hola"
	address,tx,instance,err := store.DeployStore(auth,client,stringItem)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}