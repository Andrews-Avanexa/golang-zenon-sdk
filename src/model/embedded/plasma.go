package embedded

import (
	"math/big"

	"github.com/Andrews-Avanexa/golang-zenon-sdk/src/model/primitives"
)

type FusionEntry struct {
	QsrAmount        *big.Int           `json:"qsrAmount"`
	Beneficiary      primitives.Address `json:"beneficiary"`
	ExpirationHeight uint64             `json:"expirationHeight"`
	Id               primitives.Hash    `json:"id"`
}
type FusionEntryList struct {
	QsrAmount *big.Int       `json:"qsrAmount"`
	Count     int            `json:"count"`
	Fusions   []*FusionEntry `json:"list"`
}

type PlasmaInfo struct {
	CurrentPlasma uint64   `json:"currentPlasma"`
	MaxPlasma     uint64   `json:"maxPlasma"`
	QsrAmount     *big.Int `json:"qsrAmount"`
}

type GetRequiredParam struct {
	SelfAddr  primitives.Address  `json:"address"`
	BlockType uint64              `json:"blockType"`
	ToAddr    *primitives.Address `json:"toAddress"`
	Data      []byte              `json:"data"`
}

type GetRequiredResponse struct {
	AvailablePlasma    uint64   `json:"availablePlasma"`
	BasePlasma         uint64   `json:"basePlasma"`
	RequiredDifficulty *big.Int `json:"requiredDifficulty"`
}
