package nom

import (
	"math/big"

	"github.com/Andrews-Avanexa/golang-zenon-sdk/src/common"
	"github.com/Andrews-Avanexa/golang-zenon-sdk/src/model/primitives"
)

var (
	balanceKeyPrefix = []byte{3}
)

type BalanceInfo struct {
	TokenInfo *Token   `json:"token"`
	Balance   *big.Int `json:"balance"`
}
type AccountInfo struct {
	Address        primitives.Address                             `json:"address"`
	AccountHeight  uint64                                         `json:"accountHeight"`
	BalanceInfoMap map[primitives.ZenonTokenStandard]*BalanceInfo `json:"balanceInfoMap"`
}

func getBalanceKey(zts primitives.ZenonTokenStandard) []byte {
	return common.JoinBytes(balanceKeyPrefix, zts.Bytes())
}

func znn() {
	getBalanceKey(primitives.ZnnTokenStandard)
}

func qsr() {
	getBalanceKey(primitives.QsrTokenStandard)
}
