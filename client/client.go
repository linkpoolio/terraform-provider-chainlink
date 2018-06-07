package client

import (
	"fmt"
	"bytes"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

func NewNodeClient(c Config) *NodeClient {
	return &NodeClient{Config: c}
}

func (c *NodeClient) CreateBridgeType(addr, name, url string) error {
	bridgeType := BridgeType{Name: name, Url: url}
	b, err := json.Marshal(bridgeType)
	if err != nil {
		return err
	}

	client := http.Client{}
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s://%s/v2/bridge_types", c.Config.Protocol, addr),
		bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.Config.Username, c.Config.Password)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected response code, got %d, expected 200", resp.StatusCode)
	}
	return nil
}

func (c *NodeClient) ReadBridgeType(id string) (*BridgeType, error) {
	m := NewMatcherFromId(id)

	client := http.Client{}
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s://%s/v2/bridge_types/%s", c.Config.Protocol, m.NodeAddress, m.Data),
		nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Config.Username, c.Config.Password)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected response code, got %d, expected 200", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bT := BridgeType{}
	err = json.Unmarshal(b, &bT)
	if err != nil {
		return nil, err
	}

	return &bT, nil
}

func (c *NodeClient) DeleteBridgeType(id string) error {
	m := NewMatcherFromId(id)

	client := http.Client{}
	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s://%s/v2/bridge_types/%s", c.Config.Protocol, m.NodeAddress, m.Data),
		nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.Config.Username, c.Config.Password)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected response code, got %d, expected 200", resp.StatusCode)
	}

	return nil
}