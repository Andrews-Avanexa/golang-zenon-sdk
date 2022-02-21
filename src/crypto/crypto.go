package crypto

import (
	crypt "crypto"
)

func _ed25519HashFunc(m []uint8) []uint8 {
	var sink = crypt.BLAKE2b_512.HashFunc().New()
	hash := sink.Sum(m)
	return hash
}

func getPublicKey(privateKey []byte) []int {
	return GenerateKey()
}
