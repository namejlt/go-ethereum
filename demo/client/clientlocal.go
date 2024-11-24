package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/demo/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

/**

连接区块链


Ganache 是一个用于以太坊（Ethereum）开发和测试的本地区块链模拟器。
它提供了一个易于使用的界面，帮助开发者快速搭建和测试以太坊智能合约和去中心化应用（DApps）。
Ganache 是 Truffle 框架的一部分，Truffle 是一个流行的以太坊开发框架，提供了智能合约编译、部署、测试和调试的工具。

这里用已弃用的ganache-cli来跑demo代码

npm install -g ganache-cli

ganache-cli

这个时候已经运行一个本地链
RPC Listening on 127.0.0.1:8545

*/

func main() {
	client, err := ethclient.Dial(common.ServiceUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections
}
