package websocket

import (
	"github.com/davyxu/cellnet"
)

// Peer间的共享数据
//  如果结构体没有命名的话，默认名字为结构体的名字
type wsPeer struct {
	cellnet.EventQueue

	cellnet.SessionManager

	*cellnet.PeerProfileImplement

	*cellnet.HandlerChainManagerImplement
}

func (self *wsPeer) Queue() cellnet.EventQueue {
	return self.EventQueue
}

func newPeer(queue cellnet.EventQueue, sm cellnet.SessionManager) *wsPeer {

	self := &wsPeer{
		EventQueue:                   queue,
		SessionManager:               sm,
		PeerProfileImplement:         cellnet.NewPeerProfile(),
		HandlerChainManagerImplement: cellnet.NewHandlerChainManager(),
	}

	self.SetChainSend(
		cellnet.NewHandlerChain(
			cellnet.StaticEncodePacketHandler(),
		),
	)

	return self
}
