package embedded

import (
	"math/big"
	"sort"

	"github.com/Andrews-Avanexa/golang-zenon-sdk/src/model/primitives"

	"github.com/zenon-network/go-zenon/common"
	"github.com/zenon-network/go-zenon/rpc/api"
	"github.com/zenon-network/go-zenon/vm/constants"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"github.com/zenon-network/go-zenon/vm/embedded/implementation"
)

func (a *PillarApi) GetDepositedQsr(address Address) (*big.Int, error) {
	return getDepositedQsr(a.chain, primitives.PillarAddress, address)
}
func (a *PillarApi) GetUncollectedReward(address primitives.Address) (*definition.RewardDeposit, error) {
	return getUncollectedReward(a.chain, primitives.PillarAddress, address)
}
func (a *PillarApi) GetFrontierRewardByPage(address primitives.Address, pageIndex, pageSize uint32) (*RewardHistoryList, error) {
	if pageSize > api.RpcMaxPageSize {
		return nil, api.ErrPageSizeParamTooBig
	}
	return getFrontierRewardByPage(a.chain, primitives.PillarAddress, address, pageIndex, pageSize)
}

func (a *PillarApi) GetQsrRegistrationCost() (*big.Int, error) {
	_, context, err := api.GetFrontierContext(a.chain, primitives.PillarAddress)
	if err != nil {
		return nil, err
	}

	currentQsrCost, err := implementation.GetQsrCostForNextPillar(context)
	if err != nil {
		return nil, err
	}

	return currentQsrCost, nil
}

type PillarInfoByWeight []*PillarInfo

func (a PillarInfoByWeight) Len() int      { return len(a) }
func (a PillarInfoByWeight) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a PillarInfoByWeight) Less(i, j int) bool {
	r := a[j].Weight.Cmp(a[i].Weight)
	if r == 0 {
		return a[i].Name < a[j].Name
	} else {
		return r < 0
	}
}

func (a *PillarApi) GetAll(pageIndex, pageSize uint32) (*PillarInfoList, error) {
	if pageSize > api.RpcMaxPageSize {
		return nil, api.ErrPageSizeParamTooBig
	}

	m, context, err := api.GetFrontierContext(a.chain, primitives.PillarAddress)
	if err != nil {
		return nil, err
	}

	// pillars
	candidateList, err := definition.GetPillarsList(context.Storage(), true, definition.AnyPillarType)
	if err != nil {
		return nil, err
	}

	targetList := make([]*PillarInfo, len(candidateList))

	for index, pillar := range candidateList {
		// canBeRevoked
		canBeRevoked, revokeCooldown := implementation.PillarGetRevokeStatus(pillar, m)

		targetList[index] = &PillarInfo{
			Name:                         pillar.Name,
			Type:                         pillar.PillarType,
			StakeAddress:                 pillar.StakeAddress,
			BlockProducingAddress:        pillar.BlockProducingAddress,
			RewardWithdrawAddress:        pillar.RewardWithdrawAddress,
			RevokeTime:                   pillar.RevokeTime,
			GiveMomentumRewardPercentage: pillar.GiveBlockRewardPercentage,
			GiveDelegateRewardPercentage: pillar.GiveDelegateRewardPercentage,
			CanBeRevoked:                 canBeRevoked,
			RevokeCooldown:               revokeCooldown,
			CurrentStats: &PillarEpochStats{
				ProducedMomentums: 0,
				ExpectedMomentums: 0,
			},
			Weight: common.Big0,
		}
	}

	// feed information from rpc consensus cache
	weights, stats := a.consensusCache.Get()
	if weights != nil {
		for _, pillar := range targetList {
			weight, ok := weights[pillar.Name]
			if ok == false {
				pillar.Weight = big.NewInt(0)
			} else {
				pillar.Weight = (&big.Int{}).Set(weight)
			}
		}
	}

	if stats != nil {
		for _, pillar := range targetList {
			pillarStat, ok := stats.Pillars[pillar.Name]
			if ok == true {
				pillar.CurrentStats.ProducedMomentums = pillarStat.BlockNum
				pillar.CurrentStats.ExpectedMomentums = pillarStat.ExceptedBlockNum
			}
		}
	}

	sort.Sort(PillarInfoByWeight(targetList))
	for i := range targetList {
		targetList[i].Rank = i
	}

	start, end := api.GetRange(pageIndex, pageSize, uint32(len(targetList)))

	return &PillarInfoList{
		Count: uint32(len(targetList)),
		List:  targetList[start:end],
	}, nil
}

func (a *PillarApi) GetByOwner(stakeAddress primitives.Address) ([]*PillarInfo, error) {
	list, err := a.GetAll(0, api.RpcMaxPageSize)
	if err != nil {
		return nil, err
	}
	targetList := make([]*PillarInfo, 0)
	for _, pillar := range list.List {
		if pillar.StakeAddress == stakeAddress {
			targetList = append(targetList, pillar)
		}
	}

	return targetList, nil
}

func (a *PillarApi) GetByName(name string) (*PillarInfo, error) {
	list, err := a.GetAll(0, api.RpcMaxPageSize)
	if err != nil {
		return nil, err
	}
	for _, pillar := range list.List {
		if pillar.Name == name {
			return pillar, nil
		}
	}

	return nil, nil
}

func (a *PillarApi) CheckNameAvailability(name string) (bool, error) {
	_, context, err := api.GetFrontierContext(a.chain, primitives.PillarAddress)
	if err != nil {
		return false, err
	}

	// pillars
	pillars, err := definition.GetPillarsList(context.Storage(), false, definition.AnyPillarType)
	if err != nil {
		return false, err
	}

	for _, pillar := range pillars {
		if pillar.Name == name {
			return false, nil
		}
	}
	return true, nil
}

func (a *PillarApi) GetDelegatedPillar(addr primitives.Address) (*GetDelegatedPillarResponse, error) {
	_, context, err := api.GetFrontierContext(a.chain, primitives.PillarAddress)
	if err != nil {
		return nil, err
	}
	delegationInfo, err := definition.GetDelegationInfo(context.Storage(), addr)
	if err == constants.ErrDataNonExistent {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if delegationInfo != nil {
		balance, err := a.chain.GetFrontierMomentumStore().GetAccountStore(addr).GetBalance(primitives.ZnnTokenStandard)
		if err != nil {
			return nil, err
		}
		status := PillarInActive
		if pillar, err := definition.GetPillarInfo(context.Storage(), delegationInfo.Name); err == constants.ErrDataNonExistent {
		} else if err == nil {
			if pillar.RevokeTime == 0 {
				status = PillarActive
			}
		} else {
			return nil, err
		}

		return &GetDelegatedPillarResponse{
			Name:       delegationInfo.Name,
			NodeStatus: status,
			Balance:    balance}, nil

	}
	return nil, nil
}

func (a *PillarApi) GetPillarEpochHistory(pillarName string, pageIndex, pageSize uint32) (*PillarEpochHistoryList, error) {
	if pageSize > api.RpcMaxPageSize {
		return nil, api.ErrPageSizeParamTooBig
	}

	_, context, err := api.GetFrontierContext(a.chain, primitives.PillarAddress)
	if err != nil {
		return nil, err
	}

	// get latest epoch
	lastEpoch, err := definition.GetLastEpochUpdate(context.Storage())
	if err != nil {
		return nil, err
	}

	epoch := lastEpoch.LastEpoch - int64(pageIndex*pageSize)

	result := &PillarEpochHistoryList{
		Count: lastEpoch.LastEpoch + 1,
		List:  make([]*definition.PillarEpochHistory, 0, pageSize),
	}
	for i := 0; i < int(pageSize); i += 1 {
		if epoch < 0 {
			break
		}
		if pillars, err := definition.GetPillarEpochHistoryList(context.Storage(), uint64(epoch)); err == nil {
			found := false
			for _, pillar := range pillars {
				if pillar.Name == pillarName {
					result.List = append(result.List, pillar)
					found = true
					break
				}
			}
			if !found {
				result.List = append(result.List, &definition.PillarEpochHistory{
					Name:                         pillarName,
					Epoch:                        uint64(epoch),
					GiveDelegateRewardPercentage: 0,
					GiveBlockRewardPercentage:    0,
					ProducedBlockNum:             0,
					ExpectedBlockNum:             0,
					Weight:                       common.Big0,
				})
			}
		} else {
			return nil, err
		}
		epoch -= 1
	}

	return result, err
}

func (a *PillarApi) GetPillarsHistoryByEpoch(epoch uint64, pageIndex, pageSize uint32) (*PillarEpochHistoryList, error) {
	if pageSize > api.RpcMaxPageSize {
		return nil, api.ErrPageSizeParamTooBig
	}

	_, context, err := api.GetFrontierContext(a.chain, primitives.PillarAddress)
	if err != nil {
		return nil, err
	}

	pillars, err := definition.GetPillarEpochHistoryList(context.Storage(), epoch)
	if err != nil {
		return nil, err
	}

	start, end := api.GetRange(pageIndex, pageSize, uint32(len(pillars)))

	return &PillarEpochHistoryList{
		Count: int64(len(pillars)),
		List:  pillars[start:end],
	}, nil
}
