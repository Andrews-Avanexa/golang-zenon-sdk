package nom

import (
	"github.com/Andrews-Avanexa/golang-zenon-sdk/src/model/primitives"
)

type AccountBlockConfirmationDetail struct {
	NumConfirmations  uint64          `json:"numConfirmations"`
	MomentumHeight    uint64          `json:"momentumHeight"`
	MomentumHash      primitives.Hash `json:"momentumHash"`
	MomentumTimestamp int64           `json:"momentumTimestamp"`
}

type AccountBlock struct {
	DescendantBlocks []*AccountBlock `json:"descendantBlocks"` // hash of DescendantBlocks is included in hash
	BasePlasma       uint64          `json:"basePlasma"`       // not included in hash, the smallest value of TotalPlasma required for block
	TotalPlasma      uint64          `json:"usedPlasma"`       // not included in hash, TotalPlasma = FusedPlasma + PowPlasma
	ChangesHash      primitives.Hash `json:"changesHash"`      // not included in hash
	Token            Token           `json:"token"`
}
type AccountBlockList struct {
	List  []*AccountBlock `json:"list"`
	Count int             `json:"count"`
	More  bool            `json:"more"`
}
