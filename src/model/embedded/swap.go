package embedded

import "math/big"

type SwapAssetEntry struct {
	KeyIdHash string   `json:"keyIdHash"`
	Znn       *big.Int `json:"znn"`
	Qsr       *big.Int `json:"qsr"`
}
type SwapLegacyPillarEntry struct {
	KeyIdHash  string `json:"keyIdHash"`
	NumPillars int    `json:"numPillars"`
}
