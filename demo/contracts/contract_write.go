package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/demo/contracts/contracts/bank"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

/**

调用智能合约写函数，需要设置交易的一些参数，秘钥授权

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

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54") //合约地址
	instance, err := bank.NewBank(address, client)
	if err != nil {
		log.Fatal(err)
	}

	opt := bind.TransactOpts{}
	opt.From = fromAddress
	opt.Nonce = big.NewInt(int64(nonce))
	opt.Value = big.NewInt(0)     // in wei
	opt.GasLimit = uint64(300000) // in units
	opt.GasPrice = gasPrice
	opt.From = fromAddress

	amount := big.NewInt(2000000000000000000)

	tx, err := instance.Workearncoin(&opt, amount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

}
