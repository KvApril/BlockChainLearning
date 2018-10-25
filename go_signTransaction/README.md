### 目录结构
```cgo
.
├── contracts
│   ├── test.go
│   ├── test.sol
│   ├── test_sol_test.abi
│   └── test_sol_test.bin
├── gocontract.go
├── gocontractmethod.go
└── gosign001.go

```
### 目录说明
1. contracts: 合约代码以及编译得到的abi和bin
2. gosign001.go: 发送离线签名的转账交易
3. gocontract.go: 发送离线签名的创建合约的交易
4. gocontractmethod.go: 发送离线签名调用合约方法交易

### 具体
gosign001.go signTx:输出:
```cgo
signTx:  
	TX(54bd176928ffd4e9490cefe100670394d1e483f0cde0e024fa14e726a25c7f3b)
	Contract: false
	From:     575f174218e725a15f3e427fafc2fa844c6cfbec
	To:       1aa69bf9f75f15993b2d03d92597605ff9231505
	Nonce:    169
	GasPrice: 0x0
	GasLimit  0x989680
	Value:    0xbd3580
	Data:     0x
	V:        0x41
	R:        0x756bb568ddf2005dbd08841fa491ec0e76b25500516df684e2f6ec794c212349
	S:        0x5e11a9e265d65f408f3db495758ee6e46961f1ddda8c0160dbf93aa83b7eccab
	Hex:      f86481a98083989680941aa69bf9f75f15993b2d03d92597605ff923150583bd35808041a0756bb568ddf2005dbd08841fa491ec0e76b25500516df684e2f6ec794c212349a05e11a9e265d65f408f3db495758ee6e46961f1ddda8c0160dbf93aa83b7eccab
known transaction: 54bd176928ffd4e9490cefe100670394d1e483f0cde0e024fa14e726a25c7f3b

签名的code = Hex = f86481a98083989680941aa69bf9f75f15993b2d03d92597605ff923150583bd35808041a0756bb568ddf2005dbd08841fa491ec0e76b25500516df684e2f6ec794c212349a05e11a9e265d65f408f3db495758ee6e46961f1ddda8c0160dbf93aa83b7eccab

可以使用rlpdump来验证一下签名;rlpdump代码位置github.com/ethereum/go-ethereum/cmd/rlpdump/

使用方式:
$ ./rlpdump -hex f86481a98083989680941aa69bf9f75f15993b2d03d92597605ff923150583bd35808041a0756bb568ddf2005dbd08841fa491ec0e76b25500516df684e2f6ec794c212349a05e11a9e265d65f408f3db495758ee6e46961f1ddda8c0160dbf93aa83b7eccab
  [
    a9,
    "",
    989680,
    1aa69bf9f75f15993b2d03d92597605ff9231505,
    bd3580,
    "",
    "A",
    756bb568ddf2005dbd08841fa491ec0e76b25500516df684e2f6ec794c212349,
    5e11a9e265d65f408f3db495758ee6e46961f1ddda8c0160dbf93aa83b7eccab,
  ]

可以在geth console中使用: eth.sendRawTransaction("0xf86481a98083989680941aa69bf9f75f15993b2d03d92597605ff923150583bd35808041a0756bb568ddf2005dbd08841fa491ec0e76b25500516df684e2f6ec794c212349a05e11a9e265d65f408f3db495758ee6e46961f1ddda8c0160dbf93aa83b7eccab")来发送该交易
```