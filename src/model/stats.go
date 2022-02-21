package model

import (
	"runtime"
	"strings"

	"github.com/inconshreveable/log15"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"

	"github.com/zenon-network/go-zenon/common"
	"github.com/zenon-network/go-zenon/metadata"
	"github.com/zenon-network/go-zenon/p2p"
	"github.com/zenon-network/go-zenon/protocol"
	"github.com/zenon-network/go-zenon/zenon"
)

type Peer struct {
	PublicKey string `json:"publicKey"`
	IP        string `json:"ip"`
}

func PeerfromJson(peer *p2p.Peer) (*Peer, error) {
	ip := peer.RemoteAddr().String()
	splits := strings.Split(ip, ":")
	return &Peer{
		PublicKey: peer.ID().String(),
		IP:        splits[0],
	}, nil
}

func selfToPeer(node *discover.Node) *Peer {
	return &Peer{
		PublicKey: node.ID.String(),
		IP:        "127.0.0.1",
	}
}

type NetworkInfoResponse struct {
	NumPeers int     `json:"numPeers"`
	Self     *Peer   `json:"self"`
	Peers    []*Peer `json:"peers"`
}

func (api *StatsApi) NetworkInfo() (*NetworkInfoResponse, error) {
	peersRaw := api.p2p.Peers()
	peers := make([]*Peer, 0, len(peersRaw))
	for _, raw := range peersRaw {
		peer, err := PeerfromJson(raw)
		if err != nil {
			return nil, err
		}
		peers = append(peers, peer)
	}

	return &NetworkInfoResponse{
		NumPeers: api.p2p.PeerCount(),
		Peers:    peers,
		Self:     selfToPeer(api.p2p.Self()),
	}, nil
}

type ProcessInfoResponse struct {
	Commit  string `json:"commit"`
	Version string `json:"version"`
}

func (api *StatsApi) ProcessInfo() (*ProcessInfoResponse, error) {
	return &ProcessInfoResponse{
		Version: metadata.Version,
		Commit:  metadata.GitCommit,
	}, nil
}

type OsInfoResponse struct {
	Os              string `json:"os"`
	Platform        string `json:"platform"`
	PlatformVersion string `json:"platformVersion"`
	KernelVersion   string `json:"kernelVersion"`
	MemoryTotal     uint64 `json:"memoryTotal"`
	MemoryFree      uint64 `json:"memoryFree"`
	NumCPU          int    `json:"numCPU"`
	NumGoroutine    int    `json:"numGoroutine"`
}

type StatsApi struct {
	z   zenon.Zenon
	p2p *p2p.Server
	log log15.Logger
}

func NewStatsApi(z zenon.Zenon, p2p *p2p.Server) *StatsApi {
	return &StatsApi{
		z:   z,
		p2p: p2p,
		log: common.RPCLogger.New("module", "net_api"),
	}
}

func (api *StatsApi) OsInfo() (*OsInfoResponse, error) {
	result := &OsInfoResponse{}
	stat, e := host.Info()
	if e == nil {
		result.Os = stat.OS
		result.Platform = stat.Platform
		result.PlatformVersion = stat.PlatformVersion
		result.KernelVersion = stat.KernelVersion
	}

	memO, e := mem.VirtualMemory()
	if e == nil {
		result.MemoryFree = memO.Available
		result.MemoryTotal = memO.Total
	}

	result.NumCPU = runtime.NumCPU()
	result.NumGoroutine = runtime.NumGoroutine()
	return result, nil
}

func (api *StatsApi) SyncInfo() (*protocol.SyncInfo, error) {
	return api.z.Broadcaster().SyncInfo(), nil
}
