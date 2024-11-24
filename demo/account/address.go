package main

import (
	"fmt"
	common2 "github.com/ethereum/go-ethereum/demo/common"

	"github.com/ethereum/go-ethereum/common"
)

/**

获取区块链账户地址


*/

func main() {
	address := common.HexToAddress(common2.AccountAddress)

	fmt.Println(address.Hex())    // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F
	fmt.Println(address.String()) // 0x00000000000000000000000071c7656ec7ab88b098defb751b7401b5f6d8976f
	fmt.Println(address.Bytes())  // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111]
}
