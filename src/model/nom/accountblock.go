package nom

import "github.com/Andrews-Avanexa/golang-zenon-sdk/src/common/types"

type AccountBlockConfirmationDetail struct {
	NumConfirmations  uint64     `json:"numConfirmations"`
	MomentumHeight    uint64     `json:"momentumHeight"`
	MomentumHash      types.Hash `json:"momentumHash"`
	MomentumTimestamp int64      `json:"momentumTimestamp"`
}
