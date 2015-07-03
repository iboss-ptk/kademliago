package kademliago

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNode(t *testing.T) {
	udpPort := "8000"
	ipAddress := "127.144.3.4"
	node := NewNode(udpPort, ipAddress)
	assert.Equal(t, node.UdpPort, udpPort, "incorrect UDPport")
	assert.Equal(t, node.IpAddress, ipAddress, "incoorect IPadress")
	assert.Equal(t, len(node.Id), 20, "incoorect id length")
}
