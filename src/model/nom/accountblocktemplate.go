package nom

import (
	"crypto/ed25519"
	"math/big"

	"github.com/Andrews-Avanexa/golang-zenon-sdk/src/model/primitives"
)

type Nonce struct {
	Data [8]byte
}
type BlockTypeEnum int

const (
	BlockTypeGenesisReceive = 1 // receive

	BlockTypeUserSend    = 2 // send
	BlockTypeUserReceive = 3 // receive

	BlockTypeContractSend    = 4 // send
	BlockTypeContractReceive = 5 // receive
)

type AccountBlockTemplate struct {
	Version         uint64 `json:"version"`
	ChainIdentifier uint64 `json:"chainIdentifier"`
	BlockType       uint64 `json:"blockType"`

	Hash                 primitives.Hash       `json:"hash"`
	PreviousHash         primitives.Hash       `json:"previousHash"`
	Height               uint64                `json:"height"`
	MomentumAcknowledged primitives.HashHeight `json:"momentumAcknowledged"`

	Address primitives.Address `json:"address"`

	// Send information
	ToAddress     primitives.Address            `json:"toAddress"`
	Amount        *big.Int                      `json:"amount"`
	TokenStandard primitives.ZenonTokenStandard `json:"tokenStandard"`

	// Receive information
	FromBlockHash primitives.Hash `json:"fromBlockHash"`

	// Batch information
	DescendantBlocks []*AccountBlock `json:"descendantBlocks"` // hash of DescendantBlocks is included in hash

	Data []byte `json:"data"` // hash of Data is included in hash

	FusedPlasma uint64 `json:"fusedPlasma"`
	Difficulty  uint64 `json:"difficulty"`
	Nonce       Nonce  `json:"nonce"`
	BasePlasma  uint64 `json:"basePlasma"` // not included in hash, the smallest value of TotalPlasma required for block
	TotalPlasma uint64 `json:"usedPlasma"` // not included in hash, TotalPlasma = FusedPlasma + PowPlasma

	ChangesHash primitives.Hash `json:"changesHash"` // not included in hash

	PublicKey ed25519.PublicKey `json:"publicKey"` // not included in hash
	Signature []byte            `json:"signature"` // not included in hash
}

func NewAccountBlockTemplate(blockType uint64, toAddress primitives.Address, amount *big.Int, tokenStandard primitives.ZenonTokenStandard, fromBlockHash primitives.Hash, data []byte) *AccountBlockTemplate {
	abt := new(AccountBlockTemplate)
	abt.Version = 1
	abt.ChainIdentifier = 6
	abt.Hash = primitives.ZeroHash
	abt.PreviousHash = primitives.ZeroHash
	abt.Height = 0
	abt.MomentumAcknowledged = primitives.ZeroHashHeight
	abt.Address = primitives.emptyAddress
	if toAddress == nil {
		abt.ToAddress = primitives.emptyAddress
	} else {
		abt.ToAddress = toAddress
	}
	if amount == nil {
		abt.Amount = big.NewInt(0)
	} else {
		abt.Amount = amount
	}
	if tokenStandard == nil {
		abt.TokenStandard = TokenStandard.parse(emptyTokenStandard)
	} else {
		abt.TokenStandard = tokenStandard
	}
	if fromBlockHash == nil {
		abt.FromBlockHash = emptyHash
	} else {
		abt.FromBlockHash = fromBlockHash
	}
	if data == nil {
		abt.Data = byte(0)
	} else {
		abt.Data = data
	}
	abt.FusedPlasma = 0
	abt.Difficulty = 0
	abt.Nonce = byte(0)
	abt.PublicKey = byte(0)
	abt.Signature = byte(0)
	return abt
}
