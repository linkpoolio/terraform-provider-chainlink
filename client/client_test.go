package client

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestNodeClient_CreateReadDeleteBridgeType(t *testing.T) {
	c := newDefaultClient()
	n := RandomString(10)
	u := fmt.Sprintf("http://%s.com/", RandomString(10))
	m := NewMatcher(defaultNodeAddress(), n)

	err := c.CreateBridgeType(defaultNodeAddress(), n, u)
	assert.NoError(t, err)

	bT, err := c.ReadBridgeType(m.Id())
	assert.NoError(t, err)

	assert.Equal(t, bT.Data.Attributes.Name, n)
	assert.Equal(t, bT.Data.Attributes.Url, u)

	err = c.DeleteBridgeType(m.Id())
	assert.NoError(t, err)
}

func newDefaultClient() *NodeClient {
	return NewNodeClient(&Config{
		Email:    "chainlink",
		Password: "twochains",
		Protocol: "http",
	})
}

func defaultNodeAddress() string {
	return "localhost:6688"
}