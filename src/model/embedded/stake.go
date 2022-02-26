package embedded

import (
	"math/big"

	"github.com/Andrews-Avanexa/golang-zenon-sdk/src/model/primitives"
)

type StakeEntry struct {
	Amount              *big.Int           `json:"amount"`
	WeightedAmount      *big.Int           `json:"weightedAmount"`
	StartTimestamp      int64              `json:"startTimestamp"`
	ExpirationTimestamp int64              `json:"expirationTimestamp"`
	Address             primitives.Address `json:"address"`
	Id                  primitives.Hash    `json:"id"`
}
type StakeList struct {
	TotalAmount         *big.Int      `json:"totalAmount"`
	TotalWeightedAmount *big.Int      `json:"totalWeightedAmount"`
	Count               int           `json:"count"`
	Entries             []*StakeEntry `json:"list"`
}
