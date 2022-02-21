package wallet

import "fmt"

const (
	coinType       = "73404"
	derivationPath = "m/44'/" + coinType + "'"
)

func getDerivationAccount(account string) string {
	return "m/44'/" + derivationPath + "'/" + fmt.Sprintf("%s", account) + "'"
}
