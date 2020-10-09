package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var ErrNotFound = errors.New("unexpected response code, got 404")

func NewChainlink(c *Config) (*Chainlink, error) {
	cl := &Chainlink{Config: c}
	return cl, cl.setSessionCookie()
}

func (c *Chainlink) CreateSpec(spec string) (string, error) {
	specResp := NewResponse()
	_, err := c.doRaw(http.MethodPost, "/v2/specs", []byte(spec), &specResp, http.StatusOK)
	return fmt.Sprint(specResp.Data["id"]), err
}

func (c *Chainlink) CreateSpecV2(spec string) (string, error) {
	specResp := NewResponse()
	_, err := c.doRaw(http.MethodPost, "/v2/specs_v2", []byte(spec), &specResp, http.StatusOK)
	return fmt.Sprint(specResp.Data["jobID"]), err
}

func (c *Chainlink) DeleteSpecV2(id string) error {
	_, err := c.do(http.MethodDelete, fmt.Sprintf("/v2/specs_v2/%s", id), nil, nil, http.StatusNoContent)
	return err
}

func (c *Chainlink) ReadSpec(id string) (*Response, error) {
	specObj := &Response{}
	_, err := c.do(http.MethodGet, fmt.Sprintf("/v2/specs/%s", id), nil, specObj, http.StatusOK)
	return specObj, err
}

func (c *Chainlink) DeleteSpec(id string) error {
	_, err := c.do(http.MethodDelete, fmt.Sprintf("/v2/specs/%s", id), nil, nil, http.StatusNoContent)
	return err
}

func (c *Chainlink) CreateBridge(name, url string) error {
	_, err := c.do(http.MethodPost, "/v2/bridge_types", BridgeTypeAttributes{Name: name, URL: url}, nil, http.StatusOK)
	return err
}

func (c *Chainlink) ReadBridge(name string) (*BridgeType, error) {
	bt := BridgeType{}
	_, err := c.do(http.MethodGet, fmt.Sprintf("/v2/bridge_types/%s", name), nil, &bt, http.StatusOK)
	return &bt, err
}

func (c *Chainlink) DeleteBridge(name string) error {
	_, err := c.do(http.MethodDelete, fmt.Sprintf("/v2/bridge_types/%s", name), nil, nil, http.StatusOK)
	return err
}

func (c *Chainlink) ReadWallet() (string, error) {
	walletObj := &ResponseArray{}
	if _, err := c.do(http.MethodGet, "/v2/user/balances", nil, &walletObj, http.StatusOK); err != nil {
		return "", err
	} else if len(walletObj.Data) == 0 {
		return "", fmt.Errorf("unexpected response back from Chainlink, no wallets were given")
	}
	return fmt.Sprint(walletObj.Data[0]["id"]), nil
}

func (c *Chainlink) doRaw(
	method,
	endpoint string,
	body []byte, obj interface{},
	expectedStatusCode int,
) (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest(
		method,
		fmt.Sprintf("%s%s", c.Config.URL, endpoint),
		bytes.NewBuffer(body),
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

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &obj)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 404 {
		return resp, ErrNotFound
	} else if resp.StatusCode != expectedStatusCode {
		return resp, fmt.Errorf("unexpected response code, got %d, expected 200", resp.StatusCode)
	}
	return resp, err
}

func (c *Chainlink) do(
	method,
	endpoint string,
	body interface{},
	obj interface{},
	expectedStatusCode int,
) (*http.Response, error) {
	b, err := json.Marshal(body)
	if body != nil && err != nil {
		return nil, err
	}
	return c.doRaw(method, endpoint, b, obj, expectedStatusCode)
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
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "clsession" {
			c.Cookie = cookie
			break
		}
	}
	if c.Cookie == nil {
		return fmt.Errorf("chainlink: session cookie wasn't returned on login")
	}
	return nil
}