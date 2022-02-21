package primitives

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/crypto/sha3"
)

const (
	prefix        = "z"
	AddressLength = 40
	userByte      = byte(0)
	contractByte  = byte(1)
	coreSize      = 20
)

type Address [coreSize]byte

var (
	PlasmaAddress           = ParseAddress("z1qxemdeddedxplasmaxxxxxxxxxxxxxxxxsctrp")
	PillarAddress           = ParseAddress("z1qxemdeddedxpyllarxxxxxxxxxxxxxxxsy3fmg")
	TokenAddress            = ParseAddress("z1qxemdeddedxt0kenxxxxxxxxxxxxxxxxh9amk0")
	SentinelAddress         = ParseAddress("z1qxemdeddedxsentynelxxxxxxxxxxxxxwy0r2r")
	SwapAddress             = ParseAddress("z1qxemdeddedxswapxxxxxxxxxxxxxxxxxxl4yww")
	StakeAddress            = ParseAddress("z1qxemdeddedxstakexxxxxxxxxxxxxxxxjv8v62")
	AcceleratorAddress      = ParseAddress("z1qxemdeddedxaccelerat0rxxxxxxxxxxp4tk22")
	EmbeddedContractAddress = []Address{PlasmaAddress, PillarAddress, TokenAddress, SentinelAddress, SwapAddress, StakeAddress, AcceleratorAddress}
)

func (addr *Address) SetBytes(b []byte) error {
	if length := len(b); length != coreSize {
		return fmt.Errorf("error address size  %v", length)
	}
	copy(addr[:], b)
	return nil
}

func BytesToAddress(b []byte) (Address, error) {
	var a Address
	err := a.SetBytes(b)
	return a, err
}

func ParseBech32(address string) (string, []byte, error) {
	hrp, decoded, err := Decode(address)
	if err != nil {
		return "", nil, err
	}
	core, err := convertBech32Bits(decoded, 5, 8, true)
	if err != nil {
		return "", nil, fmt.Errorf("unable to convert address from 5-bit to 8-bit formatting")
	}
	return hrp, core, nil
}

func ParseAddress(address string) Address {
	hrp, core, err := ParseBech32(address)
	if err != nil {
		return Address{}
	}

	if hrp != prefix {
		return Address{}
	}

	var addr Address
	err = addr.SetBytes(core)
	if err != nil {
		panic(err)
	}
	return addr
}

func (addr Address) toString() string {
	fiveBits, err := convertBech32Bits(addr[:], 8, 5, true)
	if err != nil {
		return ""
	}
	s, err := Encode(prefix, fiveBits)
	if err != nil {
		return ""
	}
	if err != nil {
		panic(err)
	}
	return s
}

func (addr Address) toShortString() string {
	var longString string = addr.toString()
	return longString[0:7] + "..." + longString[len(longString)-6:]
}

func (addr Address) equals() bool {
	return bytes.Equal(addr.Bytes(), Address{}.Bytes())
}

func IsEmbedded(addr Address) bool {
	return addr[0] == contractByte
}

func isValid(address string) bool {
	addr := ParseAddress(address)
	return addr.toString() == address
}

func FromPublicKey(pubKey []byte) Address {
	hash := sha3.Sum256(pubKey)
	var addr Address
	err := addr.SetBytes(append([]byte{userByte}, hash[:contractByte]...))
	if err != nil {
		panic(err)
	}
	return addr
}

func (addr Address) Bytes() []byte { return addr[:] }

func (addr Address) compareTo(otherAddress Address) int {
	return strings.Compare(addr.toString(), otherAddress.toString())
}
