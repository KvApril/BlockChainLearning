package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"context"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

const key  = `{"address":"575f174218e725a15f3e427fafc2fa844c6cfbec","crypto":{"cipher":"aes-128-ctr","ciphertext":"8eb38f492c9c1340c599bf1f40fbfaa5da663ca7b6b78f7bd0b6747597966f6c","cipherparams":{"iv":"f474a1c15ba676779542bc033174ecfb"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"10fc0a654e87d3b9659560a31f73e497775fc818c281788f6d0a150ce5d7b30c"},"mac":"e3f51e2963f2aaf9635f2af6daa4608329c0cb6cb6a38f22a19e8a3c80fff377"},"id":"130be720-33d3-483b-baa8-f6239a663a17","version":3}`

func main()  {
	client,err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	rawTransaction(client)
}

func rawTransaction(client *ethclient.Client)  {
	unlockedKey,err := keystore.DecryptKey([]byte(key),"123") //解锁私钥

	nonce,_ := client.NonceAt(context.Background(),unlockedKey.Address,nil) //获取私钥所代表账户发送交易nonce值

	if err != nil {
		fmt.Println("wrong passcode")
	}else {
		to := common.HexToAddress("0x1aa69bf9f75f15993b2d03d92597605ff9231505") //向某账户转账
		tx := types.NewTransaction(nonce,to,big.NewInt(12400000), big.NewInt(10000000), big.NewInt(0), nil) //创建交易
		signTx,err := types.SignTx(tx,types.NewEIP155Signer(big.NewInt(15)),unlockedKey.PrivateKey) //使用私钥签名交易
		fmt.Println("signTx: ",signTx)
		if err != nil {
			fmt.Println(err)
		}

		err = client.SendTransaction(context.Background(),signTx) //发送交易，实际调用的是eth_sendRawTransaction接口
		if err != nil {
			fmt.Println(err)
		}
	}

}
