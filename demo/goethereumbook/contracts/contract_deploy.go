package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/demo/goethereumbook/contracts/contracts/bank"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

/**

通过go来发布部署智能合约

核心是发送一个交易请求，交易请求地址为0，数据为智能合约的数据

1. 编写智能合约：使用 Solidity 或其他智能合约语言编写智能合约。
2. 编译智能合约：使用 Solidity 编译器（如 solc）编译智能合约，生成 ABI 和字节码。
3. 生成 Go 代码：使用 abigen 工具从 ABI 和字节码生成 Go 代码。  可以到remix.ethereum.org部署后获取相关解析的代码
4. 部署智能合约：使用生成的 Go 代码部署智能合约到以太坊网络。


*/

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA) //部署人
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// 部署智能合约

	// 合约的初始化入参
	inputCoinAddress := common.HexToAddress(`0x71C7656EC7ab88b098defB751B7401B5f6d8976F`) // 合约地址
	inputBankAddress := common.HexToAddress(`0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d`) // 外部地址
	inputWorkCoin := big.NewInt(10000000000000000000)                                     // 10

	address, tx, instance, err := bank.DeployBank(auth, client, inputCoinAddress, inputBankAddress, inputWorkCoin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0

	_ = instance
}
