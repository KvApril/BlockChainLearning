### 目录结构
```$xslt
tree
.
├── contract
│   ├── StoreA.go
│   ├── StoreA.sol
│   ├── StoreA_sol_StoreA.abi
│   └── StoreA_sol_StoreA.bin
├── event_read.go
└── README.md

```

### 主要步骤
1.编写一段合约代码  
2.使用solc编译合约代码  
3.使用abigen生成go代码  
4.部署合约到链上,得到合约部署地址  
5.创建event log查询条件    
6.获取指定合约地址和指定块区间的event log,循环输出匹配的结果


### 相关操作命令
1.使用solc编译合约代码，生成abi和bin
```$xslt
//会将结果输出到文件中
solcjs --abi --bin StoreA.sol

//将结果打印在终端
solc --abi --bin StoreA.sol
```

2.使用abigen生成go结构代码
```$xslt
//pkg: 指定生成的包名
//out: 指定生成的go文件名称
abigen --bin=StoreA_sol_StoreA.bin --abi=StoreA_sol_StoreA.abi --pkg=store --out=StoreA.go
```



### 运行方式
1.确保本地geth私链已经开启,或者ganache-cli开启,用的是默认端口8545  
2.参考[contract_deploy](https://github.com/KvApril/blockChainLearning/tree/master/go_contract/contract_deploy)将合约部署到链上,并拿到合约地址  
3.修改event_read.go中的合约地址  
4.运行go程序: go run event_read.go 

### 环境相关
```$xslt
$ solc --version
  solc, the solidity compiler commandline interface
  Version: 0.4.24+commit.e67f0147.Linux.g++
  
$ solcjs --version
  0.4.24+commit.e67f0147.Emscripten.clang
  
$ geth version
  Geth
  Version: 1.8.14-unstable
  Git Commit: 2695fa2213fe5010a80970bca1078834662d5972
  Architecture: amd64
  Protocol Versions: [63 62]
  Network Id: 1
  Go Version: go1.10
  Operating System: linux
  GOPATH=/home/kv/gopath
  GOROOT=/usr/local/go

```

### 其他
该程序模拟读取智能合约的event事件.