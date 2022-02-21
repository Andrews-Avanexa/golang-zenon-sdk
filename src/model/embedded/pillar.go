package embedded

type PillarInfo struct {
	Name string `json:"name"`
	Rank int    `json:"rank"`
	Type uint8  `json:"type"`

	StakeAddress          primitives.Address `json:"ownerAddress"`
	BlockProducingAddress primitives.Address `json:"producerAddress"`
	RewardWithdrawAddress primitives.Address `json:"withdrawAddress"`

	CanBeRevoked   bool  `json:"isRevocable"`
	RevokeCooldown int64 `json:"revokeCooldown"`
	RevokeTime     int64 `json:"revokeTimestamp"`

	GiveMomentumRewardPercentage uint8 `json:"giveMomentumRewardPercentage"`
	GiveDelegateRewardPercentage uint8 `json:"giveDelegateRewardPercentage"`

	CurrentStats *PillarStats `json:"currentStats"`
	Weight       *big.Int     `json:"weight"`
}

type PillarInfoList struct {
	Count uint32        `json:"count"`
	List  []*PillarInfo `json:"list"`
}

type PillarEpochStats struct {
	ProducedMomentums uint64 `json:"producedMomentums"`
	ExpectedMomentums uint64 `json:"expectedMomentums"`
}

type PillarEpochHistoryList struct {
	Count int64                            `json:"count"`
	List  []*definition.PillarEpochHistory `json:"list"`
}

// User delegation
type GetDelegatedPillarResponse struct {
	Name       string   `json:"name"`
	NodeStatus uint8    `json:"status"`
	Balance    *big.Int `json:"weight"`
}