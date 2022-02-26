package common

import (
	"encoding/binary"
	"math/big"

	"github.com/Andrews-Avanexa/golang-zenon-sdk/src/utils"
)

func JoinBytes(data ...[]byte) []byte {
	var newData []byte
	for _, d := range data {
		newData = append(newData, d...)
	}
	return newData
}

func Uint64ToBytes(height uint64) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, height)
	return bytes
}
func BytesToUint64(bytes []byte) uint64 {
	return binary.BigEndian.Uint64(bytes)
}

func BigIntToBytes(int *big.Int) []byte {
	if int == nil {
		return utils.LeftPadBytes(Big0.Bytes(), 32)
	} else {
		return utils.LeftPadBytes(int.Bytes(), 32)
	}
}
func BytesToBigInt(bytes []byte) *big.Int {
	if len(bytes) == 0 {
		return big.NewInt(0)
	} else {
		return new(big.Int).SetBytes(bytes)
	}
}
