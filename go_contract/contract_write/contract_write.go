package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/ethereum/go-ethereum/crypto"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
	"contract_write/contract"
	"context"
)

func main()  {
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
	fmt.Println("fromAddress=",fromAddress.Hex())

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
	auth.Value = big.NewInt(0) //不发送eth
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	//合约地址
	address := common.HexToAddress("0xe887722896D640B787D1E531BfB1d1765a327DEE")
	instance,err := store.NewStore(address,client)
	if err != nil {
		log.Fatal(err)
	}

	//调用合约接口: SetUserAddress(_userId,_userAddress)
	//_userId:      big.NewInt()
	//_userAddress: common.HexToAddress()
	//以上两个参数在传入前需要做一点处理
	tx,err := instance.SetItem(auth,big.NewInt(1),common.HexToAddress("0xb272f26e7d0253897b43dd19a39120b244b45a98"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}
