package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTCPTransport(t *testing.T) {
	ListenAddr := ":4000"
	tr := NewTCPTransport(ListenAddr)

	assert.Equal(t, ListenAddr, tr.ListenAddr)

	assert.Nil(t, tr.ListenAndAccept())
}
