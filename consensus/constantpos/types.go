package constantpos

import (
	"sync"

	libp2p "github.com/libp2p/go-libp2p-peer"
	"github.com/ninjadotorg/constant/wire"
)

type ChainInfo struct {
	CurrentCommittee        []string
	CandidateListMerkleHash string
	ChainsHeight            []int
}

type chainsHeight struct {
	Heights []int
	sync.Mutex
}

type blockSig struct {
	Validator string
	BlockSig  string
}

type swapSig struct {
	Validator string
	SwapSig   string
}

type serverInterface interface {
	// list functions callback which are assigned from Server struct
	GetPeerIDsFromPublicKey(string) []libp2p.ID
	PushMessageToAll(wire.Message) error
	PushMessageToPeer(wire.Message, libp2p.ID) error
	PushMessageGetChainState() error
}