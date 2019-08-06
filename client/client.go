package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func NewChainlink(c *Config) (*Chainlink, error) {
	cl := &Chainlink{Config: c}
	return cl, cl.setSessionCookie()
}

func (c *Chainlink) CreateSpec(spec string) (string, error) {
	specObj := make(map[string]interface{})
	specResp := Response{}
	if err := json.Unmarshal([]byte(spec), &specObj); err != nil {
		return "", err
	} else if resp, err := c.do(http.MethodPost, "/v2/specs", &specObj, &specResp);
		err != nil {
		return "", err
	} else if resp.StatusCode != 200 {
		return "", fmt.Errorf("unexpected response code, got %d, expected 200", resp.StatusCode)
	}
	return fmt.Sprint(specResp.Data["id"]), nil
}

func (c *Chainlink) ReadSpec(id string) (*Response, error) {
	specObj := &Response{}
	if resp, err := c.do(http.MethodGet, fmt.Sprintf("/v2/specs/%s", id), nil, specObj);
		err != nil {
		return specObj, err
	} else if resp.StatusCode != 200 {
		return specObj, fmt.Errorf("unexpected response code, got %d, expected 200", resp.StatusCode)
	}
	return specObj, nil
}

func (c *Chainlink) CreateBridge(name, url string) error {
	if resp, err := c.do(
		http.MethodPost,
		"/v2/bridge_types",
		BridgeTypeAttributes{Name: name, URL: url},
		nil);
	err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected response code, got %d, expected 200", resp.StatusCode)
	}
	return nil
}

func (c *Chainlink) ReadBridge(name string) (*BridgeType, error) {
	bt := BridgeType{}
	if resp, err := c.do(
		http.MethodGet,
		fmt.Sprintf("/v2/bridge_types/%s", name),
		nil,
		&bt);
	err != nil {

	} else if resp.StatusCode != 200 && resp.StatusCode != 404 {
		return nil, fmt.Errorf("unexpected response code, got %d, expected 200", resp.StatusCode)
	}
	return &bt, nil
}

func (c *Chainlink) DeleteBridge(name string) error {
	if resp, err := c.do(
		http.MethodDelete,
		fmt.Sprintf("/v2/bridge_types/%s", name),
		nil,
		nil);
	err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected response code, got %d, expected 200", resp.StatusCode)
	}
	return nil
}

func (c *Chainlink) ReadWallet() (string, error) {
	walletObj := &ResponseArray{}
	if resp, err := c.do(http.MethodGet, "/v2/user/balances", nil, &walletObj); err != nil {
		return "", err
	} else if resp.StatusCode != 200 {
		return "", fmt.Errorf("unexpected response code, got %d, expected 200", resp.StatusCode)
	} else if len(walletObj.Data) == 0 {
		return "", fmt.Errorf("unexpected response back from Chainlink, no wallets were given")
	}
	return fmt.Sprint(walletObj.Data[0]["id"]), nil
}

func (c *Chainlink) do(method, endpoint string, body interface{}, obj interface{}) (*http.Response, error) {
	b, err := json.Marshal(body)
	if body != nil && err != nil {
		return nil, err
	}

	client := http.Client{}
	req, err := http.NewRequest(
		method,
		fmt.Sprintf("%s%s", c.Config.URL, endpoint),
		bytes.NewBuffer(b),
	)
	if err != nil {
		return nil, err
	}
	req.AddCookie(c.Cookie)

	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	} else if obj == nil {
		return resp, err
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &obj)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (c *Chainlink) setSessionCookie() error {
	session := &Session{Email: c.Config.Email, Password: c.Config.Password}
	b, err := json.Marshal(session)
	if err != nil {
		return err
	}
	resp, err := http.Post(
		fmt.Sprintf("%s/sessions", c.Config.URL),
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