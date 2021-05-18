package client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

var job = `
{
  "initiators": [
    {
      "type": "runlog"
    }
  ],
  "tasks": [
    {
      "type": "httpget"
    },
    {
      "type": "jsonparse"
    },
    {
      "type": "multiply"
    },
    {
      "type": "ethuint256"
    },
    {
      "type": "ethtx"
    }
  ]
}`

func TestNodeClient_CreateReadDeleteBridgeType(t *testing.T) {
	c := newDefaultClient(t)
	n := fmt.Sprintf("adapter-%d", rand.Int())
	u := "http://adapter.com/"

	err := c.CreateBridge(n, u)
	assert.NoError(t, err)

	bT, err := c.ReadBridge(n)
	assert.NoError(t, err)

	assert.Equal(t, bT.Data.Attributes.Name, n)
	assert.Equal(t, bT.Data.Attributes.URL, u)

	err = c.DeleteBridge(n)
	assert.NoError(t, err)
}

func TestNodeClient_CreateReadDeleteSpec(t *testing.T) {
	c := newDefaultClient(t)

	id, err := c.CreateSpec(job)
	assert.NoError(t, err)

	spec, err := c.ReadSpec(id)
	assert.NoError(t, err)

	assert.Equal(t, spec.Data["id"], id)

	err = c.DeleteSpec(id)
	assert.NoError(t, err)
}

func newDefaultClient(t *testing.T) *Chainlink {
	cl, err := NewChainlink(&Config{
		Email:    "admin@node.local",
		Password: "twochains",
		URL: 	  "http://localhost:6688",
	})
	require.Nil(t, err)
	return cl
}