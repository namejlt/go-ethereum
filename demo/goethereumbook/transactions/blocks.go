package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/demo/goethereumbook/common"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

/**

查询区块

*/

func main() {
	client, err := ethclient.Dial(common.ServiceUrl)
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil) //区块的头信息
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.String()) // 5671744

	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber) //完整区块信息
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 5671744
	fmt.Println(block.Time())                // 1527211625
	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println(len(block.Transactions()))   // 144

	count, err := client.TransactionCount(context.Background(), block.Hash()) //交易总数
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 144
}
