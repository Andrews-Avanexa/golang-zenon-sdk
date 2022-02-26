package nom

import (
	"github.com/Andrews-Avanexa/golang-zenon-sdk/src/model/primitives"
)

type AccountHeader struct {
	Address    primitives.Address `json:"address"`
	HashHeight primitives.HashHeight
}
