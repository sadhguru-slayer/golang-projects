package main

import (
	"net"
)

type User struct {
	
}

type Client struct {
	ID   uint64
    Conn net.Conn
    User *User
}

// IsAuthenticated method
func (c *Client) IsAuthenticated() bool{
	return c.User != nil
}

// Add, remove, count methods for ConnectionManager
func (m *ConnectionManager) Add(client *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.clients[client.ID] = client
}

func (m *ConnectionManager) Remove(clientId uint64) {
	m.mu.Lock()
    defer m.mu.Unlock()

	delete(m.clients, clientId)
}

func (m *ConnectionManager) Count() int {
	m.mu.RLock()
    defer m.mu.RUnlock()

	return len(m.clients)
}

func (m *ConnectionManager) Get(clientId uint64) (*Client, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	client, exists := m.clients[clientId]
	return client,exists
}