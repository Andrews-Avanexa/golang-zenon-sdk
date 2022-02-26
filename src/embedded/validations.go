package embedded

import (
	"math/big"
	"regexp"

	"github.com/Andrews-Avanexa/golang-zenon-sdk/src/common"
	"github.com/pkg/errors"
)

const (
	Decimals = 100000000
	/// === Token constants ===
	TokenIssueAmount     = big.NewInt(1 * Decimals)
	TokenNameLengthMax   = 40  // Maximum length of a token name
	TokenSymbolLengthMax = 10  // Maximum length of a token symbol
	TokenDomainLengthMax = 128 // Maximum length of a token domain
	TokenMaxSupplyBig    = common.BigP255m1
	TokenMaxDecimals     = 18
	// Token
	ErrIDNotUnique        = errors.New("there is another token with the same id")
	ErrTokenInvalidText   = errors.New("invalid token name/symbol/domain/decimals")
	ErrTokenInvalidAmount = errors.New("invalid token total/max supply")
)

func checkToken(param IssueParam) error {
	// Valid names
	if len(param.TokenName) == 0 || len(param.TokenName) > constants.TokenNameLengthMax {
		return ErrTokenInvalidText
	}
	if len(param.TokenSymbol) == 0 || len(param.TokenSymbol) > constants.TokenSymbolLengthMax {
		return ErrTokenInvalidText
	}
	if len(param.TokenDomain) > constants.TokenDomainLengthMax {
		return ErrTokenInvalidText
	}

	if ok, _ := regexp.MatchString("^([a-zA-Z0-9]+[-._]?)*[a-zA-Z0-9]$", param.TokenName); !ok {
		return ErrTokenInvalidText
	}
	if ok, _ := regexp.MatchString("^[A-Z0-9]+$", param.TokenSymbol); !ok {
		return ErrTokenInvalidText
	}
	if ok, _ := regexp.MatchString("^([A-Za-z0-9][A-Za-z0-9-]{0,61}[A-Za-z0-9]\\.)+[A-Za-z]{2,}$", param.TokenDomain); ok == false && len(param.TokenDomain) != 0 {
		return ErrTokenInvalidText
	}

	if param.TokenSymbol == "ZNN" || param.TokenSymbol == "QSR" {
		return ErrTokenInvalidText
	}

	if param.Decimals > uint8(constants.TokenMaxDecimals) {
		return ErrTokenInvalidText
	}

	// 0 or too big
	if param.MaxSupply.Cmp(constants.TokenMaxSupplyBig) > 0 {
		return ErrTokenInvalidAmount
	}
	if param.MaxSupply.Cmp(common.Big0) == 0 {
		return ErrTokenInvalidAmount
	}

	// total supply is less and equal in case of non-mintable coins
	if param.MaxSupply.Cmp(param.TotalSupply) == -1 {
		return ErrTokenInvalidAmount
	}
	if param.IsMintable == false && param.MaxSupply.Cmp(param.TotalSupply) != 0 {
		return ErrTokenInvalidAmount
	}
	return nil
}
