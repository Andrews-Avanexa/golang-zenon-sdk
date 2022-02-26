package embedded

import (
	"math/big"

	"github.com/Andrews-Avanexa/golang-zenon-sdk/src/model/primitives"
)

type UncollectedReward struct {
	address   primitives.Address `json:"address"`
	znnAmount int                `json:"znnAmount"`
	qsrAmount int                `json:"qsrAmount"`
}

type RewardHistoryEntry struct {
	Epoch int64    `json:"epoch"`
	Znn   *big.Int `json:"znnAmount"`
	Qsr   *big.Int `json:"qsrAmount"`
}
type RewardHistoryList struct {
	Count int64                 `json:"count"`
	List  []*RewardHistoryEntry `json:"list"`
}

type VoteBreakdown struct {
	id    primitives.Hash `json:"id"`
	yes   int             `json:"yes"`
	no    int             `json:"no"`
	total int             `json:"total"`
}

type PillarVote struct {
	id   primitives.Hash `json:"id"`
	name string          `json:"name"`
	vote int             `json:"vote"`
}

type PillarVoteList struct {
	voteBreakdown VoteBreakdown `json:"breakdown"`
	count         int           `json:"count"`
	list          []PillarVote  `json:"list	"`
}
