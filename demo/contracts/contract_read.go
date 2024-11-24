package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/demo/contracts/contracts/bank"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

/**

调用智能合约函数

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

	address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	instance, err := bank.NewBank(address, client)
	if err != nil {
		log.Fatal(err)
	}
	opt := bind.CallOpts{}
	opt.From = fromAddress

	version, err := instance.GetMyBalance(&opt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version) // "1.0"
}
