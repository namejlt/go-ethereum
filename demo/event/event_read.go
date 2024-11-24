package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/demo/contracts/contracts/bank"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(2394201),
		ToBlock:   big.NewInt(2394201),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query) // 返回的所有日志将是ABI编码
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(bank.BankABI))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
		fmt.Println(vLog.BlockNumber)     // 2394201
		fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4

		switch vLog.Topics[0].Hex() { //通过topic0获取事件日志类型
		// 按照指定函数去解析数据，核心在于判断函数的签名是否一致，一致的话就可以解析数据
		case contractAbi.Events["transfer"].ID.Hex():
			event := &struct {
				To     common.Address
				Amount *big.Int
			}{}
			err = contractAbi.UnpackIntoInterface(&event, "transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(event)

		}

	}

}
