package wallet

import (
	"encoding/json"
	"strings"
)

func generateMnemonic(strength int) string {
	return GenerateMnemonic(strength)
}

func stringJoin(words []string) string {
	s, _ := json.Marshal(words)
	return strings.Trim(string(s), "[]")
}

func ValidateMnemonic(words []string) bool {
	return IsMnemonicValid(stringJoin(words))
}

func isValidWord(word string) bool {
	return IsValidWord(word)
}
