package embedded

import "math/big"

type IssueParam struct {
	TokenName   string
	TokenSymbol string
	TokenDomain string
	TotalSupply *big.Int
	MaxSupply   *big.Int
	Decimals    uint8
	IsMintable  bool
	IsBurnable  bool
	IsUtility   bool
}
