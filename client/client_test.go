package client

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestNodeClient_CreateBridgeType(t *testing.T) {
	c := newDefaultClient()
	err := c.CreateBridgeType(
		defaultNodeAddress(),
		RandomString(10),
		fmt.Sprintf("http://%s.com/", RandomString(10)))
	assert.NoError(t, err)
}

func TestNodeClient_ReadBridgeType(t *testing.T) {
	c := newDefaultClient()
	n := RandomString(10)
	u := fmt.Sprintf("http://%s.com/", RandomString(10))
	m := NewMatcher(defaultNodeAddress(), n)

	err := c.CreateBridgeType(defaultNodeAddress(), n, u)
	assert.NoError(t, err)
	bT, err := c.ReadBridgeType(m.Id())
	assert.NoError(t, err)

	assert.Equal(t, bT.Name, n)
	assert.Equal(t, bT.Url, u)
}

func TestNodeClient_DeleteBridgeType(t *testing.T) {
	c := newDefaultClient()
	n := RandomString(10)
	u := fmt.Sprintf("http://%s.com/", RandomString(10))
	m := NewMatcher(defaultNodeAddress(), n)

	err := c.CreateBridgeType(defaultNodeAddress(), n, u)
	assert.NoError(t, err)
	err = c.DeleteBridgeType(m.Id())
	assert.NoError(t, err)
}

func newDefaultClient() *NodeClient {
	return NewNodeClient(Config{
		Username: "chainlink",
		Password: "twochains",
		Protocol: "http",
	})
}

func defaultNodeAddress() string {
	return "localhost:6688"
}