package wallet

import (
	"encoding/hex"
	"fmt"
)

var Mnemonic string
var Entropy string
var Seed string

type argError struct {
	arg  string
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%s - %s", e.arg, e.prob)
}

func SetMnemonic(mnemonic string) {
	if !IsMnemonicValid(mnemonic) {
		print(mnemonic, &argError{mnemonic, "Invalid Mnemonic"})
	}
	Mnemonic = mnemonic
	e, _ := EntropyFromMnemonic(mnemonic)
	Entropy = hex.EncodeToString(e)
	Seed = hex.EncodeToString(NewSeed(mnemonic, ""))
}

func SetSeed(seed string) {
	Seed = seed
}

func SetEntropy(entropy string) {
	e, _ := hex.DecodeString(entropy)
	m, _ := NewMnemonic(e)
	SetMnemonic(m)
}
