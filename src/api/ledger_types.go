package api

import (
	"math/big"

	"github.com/Andrews-Avanexa/golang-zenon-sdk/src/model/primitives"
)

type DetailedMomentum struct {
	AccountBlocks []*AccountBlock `json:"blocks"`
	Momentum      *Momentum       `json:"momentum"`
}
type Momentum struct {
	Producer primitives.Address `json:"producer"`
}
type MomentumHeader struct {
	Hash      primitives.Hash `json:"hash"`
	Height    uint64          `json:"height"`
	Timestamp int64           `json:"timestamp"`
}
type AccountBlockConfirmationDetail struct {
	NumConfirmations  uint64          `json:"numConfirmations"`
	MomentumHeight    uint64          `json:"momentumHeight"`
	MomentumHash      primitives.Hash `json:"momentumHash"`
	MomentumTimestamp int64           `json:"momentumTimestamp"`
}
type AccountBlock struct {
	TokenInfo          *Token                          `json:"token"`
	ConfirmationDetail *AccountBlockConfirmationDetail `json:"confirmationDetail"`
	PairedAccountBlock *AccountBlock                   `json:"pairedAccountBlock"`
}
type AccountInfo struct {
	Address        primitives.Address                             `json:"address"`
	AccountHeight  uint64                                         `json:"accountHeight"`
	BalanceInfoMap map[primitives.ZenonTokenStandard]*BalanceInfo `json:"balanceInfoMap"`
}
type BalanceInfo struct {
	TokenInfo *Token   `json:"token"`
	Balance   *big.Int `json:"balance"`
}
type Token struct {
	TokenName          string                        `json:"name"`
	TokenSymbol        string                        `json:"symbol"`
	TokenDomain        string                        `json:"domain"`
	TotalSupply        *big.Int                      `json:"totalSupply"`
	Decimals           uint8                         `json:"decimals"`
	Owner              primitives.Address            `json:"owner"`
	ZenonTokenStandard primitives.ZenonTokenStandard `json:"tokenStandard"`
	MaxSupply          *big.Int                      `json:"maxSupply"`
	IsBurnable         bool                          `json:"isBurnable"`
	IsMintable         bool                          `json:"isMintable"`
	IsUtility          bool                          `json:"isUtility"`
}

type AccountBlockList struct {
	List  []*AccountBlock `json:"list"`
	Count int             `json:"count"`
	More  bool            `json:"more"`
}
type MomentumList struct {
	List  []*Momentum `json:"list"`
	Count int         `json:"count"`
}
type DetailedMomentumList struct {
	List  []*DetailedMomentum `json:"list"`
	Count int                 `json:"count"`
}
