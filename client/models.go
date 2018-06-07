package client

import (
	"fmt"
	"strings"
)

type Config struct {
	Username string
	Password string
	Protocol string
}

type BridgeType struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Client interface {
	CreateBridgeType(addr, name, url string) error
	ReadBridgeType(id string) (*BridgeType, error)
	DeleteBridgeType(id string) error
}

type NodeClient struct {
	Client
	Config Config
}

type DataIdentifier interface {
	Id() string
}

type Matcher struct {
	DataIdentifier
	NodeAddress string
	Data string
}

func NewMatcher(nodeAddress, data string) Matcher {
	return Matcher{
		NodeAddress: nodeAddress,
		Data: data,
	}
}

func NewMatcherFromId(id string) Matcher {
	s := strings.Split(id, Delimiter())
	return Matcher{
		NodeAddress: s[0],
		Data: s[1],
	}
}


func (i *Matcher) Id() string {
	return fmt.Sprintf("%s%s%s", i.NodeAddress, Delimiter(), i.Data)
}
