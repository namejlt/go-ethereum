package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

/**

使用模拟客户端来快速轻松地在本地测试交易

*/

func main() {
	// 生成一个随机的私钥
	privateKey, err := crypto.GenerateKey()
	// 如果生成私钥时发生错误，记录错误并退出程序
	if err != nil {
		log.Fatal(err)
	}

	// 使用私钥创建一个新的交易发送器
	auth := bind.NewKeyedTransactor(privateKey)

	// 设置初始余额为 10 ETH（以 wei 为单位）
	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10) // 10 eth in wei

	// 获取交易发送器的地址
	address := auth.From
	// 创建一个创世分配，将初始余额分配给交易发送器的地址
	genesisAlloc := map[common.Address]core.GenesisAccount{
		address: {
			Balance: balance,
		},
	}

	// 设置区块的 gas 限制为 4712388
	blockGasLimit := uint64(4712388)
	// 使用创世分配和区块 gas 限制创建一个新的模拟后端客户端
	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)

	// 获取交易发送器的地址
	fromAddress := auth.From
	// 从模拟客户端获取下一个待处理的 nonce 值
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	// 如果获取 nonce 值时发生错误，记录错误并退出程序
	if err != nil {
		log.Fatal(err)
	}

	// 设置交易的价值为 1 ETH（以 wei 为单位）
	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	// 设置交易的 gas 限制为 21000 单位
	gasLimit := uint64(21000) // in units
	// 从模拟客户端获取建议的 gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	// 如果获取 gas 价格时发生错误，记录错误并退出程序
	if err != nil {
		log.Fatal(err)
	}

	// 设置接收地址
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	// 交易数据为空
	var data []byte
	// 创建一个新的交易，包含 nonce、接收地址、价值、gas 限制、gas 价格和数据
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	// 设置链 ID 为 1337
	chainID := big.NewInt(1337)
	// 使用链 ID 对交易进行签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	// 如果签名时发生错误，记录错误并退出程序
	if err != nil {
		log.Fatal(err)
	}

	// 将签名后的交易发送到模拟客户端
	err = client.SendTransaction(context.Background(), signedTx)
	// 如果发送交易时发生错误，记录错误并退出程序
	if err != nil {
		log.Fatal(err)
	}

	// 打印交易哈希
	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex()) // tx sent: 0xec3ceb05642c61d33fa6c951b54080d1953ac8227be81e7b5e4e2cfed69eeb51

	// 提交交易到模拟客户端的区块链中
	client.Commit()

	// 从模拟客户端获取交易收据
	receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	// 如果获取交易收据时发生错误，记录错误并退出程序
	if err != nil {
		log.Fatal(err)
	}
	// 如果交易收据为空，记录错误并退出程序
	if receipt == nil {
		log.Fatal("receipt is nil. Forgot to commit?")
	}

	// 打印交易状态
	fmt.Printf("status: %v\n", receipt.Status) // status: 1
}
