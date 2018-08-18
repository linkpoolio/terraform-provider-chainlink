package client

import (
	"fmt"
	"bytes"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

func NewNodeClient(c *Config) *NodeClient {
	return &NodeClient{Config: c}
}

func (c *NodeClient) CreateBridgeType(addr, name, url string) error {
	err := c.setSessionCookie(c.Config.Protocol, addr)
	if err != nil {
		return err
	}
	bridgeType := BridgeTypeAttributes{Name: name, Url: url}
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
	req.AddCookie(c.Cookie)
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
	err := c.setSessionCookie(c.Config.Protocol, m.NodeAddress)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s://%s/v2/bridge_types/%s", c.Config.Protocol, m.NodeAddress, m.Data),
		nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(c.Cookie)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 404 {
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
	err := c.setSessionCookie(c.Config.Protocol, m.NodeAddress)
	if err != nil {
		return err
	}

	client := http.Client{}
	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s://%s/v2/bridge_types/%s", c.Config.Protocol, m.NodeAddress, m.Data),
		nil)
	if err != nil {
		return err
	}
	req.AddCookie(c.Cookie)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected response code, got %d, expected 200", resp.StatusCode)
	}

	return nil
}

func (c *NodeClient) setSessionCookie(protocol, hostname string) error {
	session := &Session{Email: c.Config.Email, Password: c.Config.Password}
	b, err := json.Marshal(session)
	if err != nil {
		return err
	}
	resp, err := http.Post(
		fmt.Sprintf("%s://%s/sessions", protocol, hostname),
		"application/json",
		bytes.NewReader(b),
	)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("status code of %d was returned when trying to get a session", resp.StatusCode)
	}
	if len(resp.Cookies()) == 0 {
		return fmt.Errorf("no cookie was returned after getting a session")
	}
	c.Cookie = resp.Cookies()[0]
	return nil
}