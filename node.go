package kademliago

import "crypto/rand"

const IdByteSize = 20

type Node struct {
	Id        []byte
	UdpPort   string
	IpAddress string
}

func NewNode(udpPort string, ipAddress string) *Node {
	id := make([]byte, IdByteSize)
	rand.Read(id)
	return &Node{id, udpPort, ipAddress}
}
