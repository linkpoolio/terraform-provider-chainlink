package client

import (
	"fmt"
	"net/http"
	"strings"
)

type Config struct {
	URL      string
	Email    string
	Password string
}

type ResponseArray struct {
	Data []map[string]interface{}
}

type Response struct {
	Data map[string]interface{}
}

func NewResponse() Response {
	return Response{
		Data: map[string]interface{}{},
	}
}

type BridgeType struct {
	Data BridgeTypeData `json:"data"`
}

type BridgeTypeData struct {
	Attributes BridgeTypeAttributes `json:"attributes"`
}

type BridgeTypeAttributes struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Session struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type OCRKeys struct {
	Data []OCRKeyData `json:"data"`
}

type OCRKey struct {
	Data OCRKeyData `json:"data"`
}

type OCRKeyData struct {
	Attributes OCRKeyAttributes `json:"attributes"`
}

type OCRKeyAttributes struct {
	ID                    string `json:"id"`
	ConfigPublicKey       string `json:"configPublicKey"`
	OffChainPublicKey     string `json:"offChainPublicKey"`
	OnChainSigningAddress string `json:"onChainSigningAddress"`
}

type P2PKeys struct {
	Data []P2PKeyData `json:"data"`
}

type P2PKey struct {
	Data P2PKeyData `json:"data"`
}

type P2PKeyData struct {
	Attributes P2PKeyAttributes `json:"attributes"`
}

type P2PKeyAttributes struct {
	ID        int    `json:"id"`
	PeerID    string `json:"peerId"`
	PublicKey string `json:"publicKey"`
}

type Client interface {
	CreateBridge(addr, name, url string) error
	ReadBridge(id string) (*BridgeType, error)
	DeleteBridge(id string) error
}

type Chainlink struct {
	Client
	Config  *Config
	Cookies []*http.Cookie
}

type DataIdentifier interface {
	Id() string
}

type Matcher struct {
	DataIdentifier
	Object string
	Data   string
}

func NewMatcher(obj, data string) Matcher {
	return Matcher{
		Object: obj,
		Data:   data,
	}
}

func NewMatcherFromID(id string) Matcher {
	s := strings.Split(id, Delimiter())
	return Matcher{
		Object: s[0],
		Data:   s[1],
	}
}

func (i *Matcher) Id() string {
	return fmt.Sprintf("%s%s%s", i.Object, Delimiter(), i.Data)
}
