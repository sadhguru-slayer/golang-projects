package main

import (
	"net"
	"sync"
	"sync/atomic"
)

type ConnectionManager struct {
	mu      sync.RWMutex
    clients map[uint64]*Client
    nextID  atomic.Uint64
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		clients: make(map[uint64]*Client),
	}
}

func (m *ConnectionManager) NewClient(conn net.Conn) *Client {
    return &Client{
        ID:   m.nextID.Add(1),
        Conn: conn,
    }
}