package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over a TCP established connection.
type TCPPeer struct {
	conn     net.Conn
	outbound bool

	wg *sync.WaitGroup
}

type TCPTransport struct {
	ListenAddr string
	Listener   net.Listener

	Mu    sync.RWMutex
	Peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		ListenAddr: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.Listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}

	go t.StartAcceptLoop()

	return nil
}

func (t *TCPTransport) StartAcceptLoop() {
	for {
		conn, err := t.Listener.Accept()
		if err != nil {
			fmt.Printf("failed to accept tcp transport %s\n", err)
		}

		go t.HandleConnection(conn)
	}
}

func (t *TCPTransport) HandleConnection(conn net.Conn) {
	fmt.Printf("New connection from %s\n", conn)
}
