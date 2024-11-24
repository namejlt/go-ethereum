package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	localcommon "github.com/ethereum/go-ethereum/demo/common"
	"github.com/ethereum/go-ethereum/rlp"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

/**

签名交易后，rlp序列化数据

这个可以发送给客户端进行交易发送

*/

func main() {
	client, err := ethclient.Dial(localcommon.ServiceUrl)
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

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 以上是创建交易

	// 这里获取交易的序列化数据

	ts := types.Transactions{signedTx}

	// 将交易转换为 RLP 格式
	rlpEncoded, err := rlp.EncodeToBytes(ts)
	if err != nil {
		log.Fatal(err)
	}

	rawTxHex := hex.EncodeToString(rlpEncoded)

	fmt.Printf(rawTxHex) // f86...772
}
