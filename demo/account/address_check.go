package main

import (
	"context"
	"fmt"
	common2 "github.com/ethereum/go-ethereum/demo/common"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

/**

检测地址有效性

正则判断地址
是否存储字节码判断是否合约地址

*/

func main() {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$") //字母仅a到f

	fmt.Printf("is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // is valid: true
	fmt.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is valid: false

	client, err := ethclient.Dial(common2.ServiceUrl)
	if err != nil {
		log.Fatal(err)
	}

	// 合约地址存储了字节码

	// 0x Protocol Token (ZRX) smart contract address
	address := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract) // is contract: true

	// a random user account address
	address = common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
	bytecode, err = client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract = len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract) // is contract: false
}
