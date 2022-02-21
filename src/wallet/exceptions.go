package wallet

import (
	"github.com/Andrews-Avanexa/golang-zenon-sdk/src"
)

type InvalidKeyStorePath struct {
	Message string
}

func (IKSP *InvalidKeyStorePath) ToString() string {
	return IKSP.Message
}

func IncorrectPasswordException() {
	ZSE := src.ZnnSdkException{Message: "Incorrect password"}
	ZSE.ToString()
}
