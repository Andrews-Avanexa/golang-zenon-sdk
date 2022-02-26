package embedded

import (
	"github.com/Andrews-Avanexa/golang-zenon-sdk/src/model/primitives"
)

type SentinelInfo struct {
	Owner                 primitives.Address `json:"owner"`
	RegistrationTimestamp int64              `json:"registrationTimestamp"`
	CanBeRevoked          bool               `json:"isRevocable"`
	RevokeCooldown        int64              `json:"revokeCooldown"`
	Active                bool               `json:"active"`
}
type SentinelInfoList struct {
	Count int             `json:"count"`
	List  []*SentinelInfo `json:"list"`
}
