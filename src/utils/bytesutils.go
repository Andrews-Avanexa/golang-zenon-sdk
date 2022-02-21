package utils

import (
	"encoding/base64"
	"encoding/hex"
	"math/big"
	"reflect"
	"strconv"
)

func Arraycopy(src []int, startPos int, dest []int, destPos int, len int) []int {
	for i := 0; i < len; i++ {
		dest[destPos+i] = src[startPos+i]
	}
	return dest
}

func DecodeBigInt(bytes []int64) *big.Int {
	result := big.NewInt(0)
	for i := 0; i < len(bytes); i++ {
		result = result.Mul(result, big.NewInt(256))
		result = result.Add(result, big.NewInt(bytes[i]))
	}
	return result
}

func EncodeBigInt(number *big.Int) []int {

	size := (number.BitLen() + 7) >> 3
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = 0
	}
	var _byteMask = big.NewInt(0xff)
	for i := 0; i < size; i++ {
		var limit big.Int
		limit.And(number, _byteMask)
		byteToInt, _ := strconv.Atoi(string(limit.Bytes()))
		result[size-i-1] = byteToInt
		number := big.NewInt(int64(byteToInt >> 8))
	}
	return result
}

func Min(v1 int, v2 int) int {
	var min int = v1
	if min > v2 {
		min = v2
	}
	return min
}

func BigIntToBytes(b *big.Int, numBytes int) []int {
	var bytes = make([]int, numBytes)
	for i := 0; i < numBytes; i++ {
		bytes[i] = 0
	}
	var biBytes = EncodeBigInt(b)
	var start int = 1
	if len(biBytes) == numBytes+1 {
		start = 0
	}
	var length int = Min(len(biBytes), numBytes)
	var dest int = numBytes - length
	Arraycopy(biBytes, start, bytes, dest, length)
	return bytes
}
func BigIntToBytesSigned(b *big.Int, numBytes int) []int {
	var bytes = make([]int, numBytes)
	var biBytes = EncodeBigInt(b)
	var start int = 1
	if len(biBytes) == numBytes+1 {
		start = 0
	}
	var length int = Min(len(biBytes), numBytes)
	var dest int = numBytes - length
	Arraycopy(biBytes, start, bytes, dest, length)
	return bytes
}

func BytesToBigInt(bb []int64) *big.Int {
	if len(bb) > 0 {
		return DecodeBigInt(bb)
	} else {
		return big.NewInt(0)
	}
}

func Merge(arrays [][]int) []int {
	count := 0
	for _, array := range arrays {
		count += len(array)
	}
	if count == 0 {
		return []int{}
	}
	mergedArray := make([]int, count)
	start := 0
	for _, array := range arrays {
		if len(array) > 0 {
			ty := reflect.TypeOf(array).Kind()
			if ty == reflect.Array {
				Arraycopy(array, 0, mergedArray, start, len(array))
				start = start + len(array)
			}
		}
	}
	return mergedArray
}

func IntToBytes(integer int) []int {
	var bytes = make([]int, 4)
	bytes[0] = integer >> 24
	bytes[1] = integer >> 16
	bytes[2] = integer >> 8
	bytes[3] = integer
	return bytes
}

func LongToBytes(longValue int) []int {
	var buffer = make([]int, 8)
	for i := 0; i < 8; i++ {
		var offset = 64 - (i+1)*8
		buffer[i] = (longValue >> offset) & 0xff
	}
	return buffer
}

func Base64ToBytes(base64Str string) []byte {
	if len(base64Str) > 0 {
		base64Decode, _ := base64.StdEncoding.DecodeString(base64Str)
		return base64Decode
	}
	return nil
}

func BytesToBase64(bytes []byte) string {
	base64Encode := base64.StdEncoding.EncodeToString(bytes)
	return base64Encode
}

func BytesToHex(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

func LeftPadBytes(bytes []int, size int) []int {

	if len(bytes) >= size {
		return bytes
	}
	var result = make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = 0
	}
	result = Arraycopy(bytes, 0, result, size-len(bytes), len(bytes))
	return result
}
