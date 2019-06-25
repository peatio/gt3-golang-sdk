package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Conf info.
var (
	Conf *Config
)

// Config struct.
type Config struct {
	// httpClinet
	HTTPClient *HTTPClient
	// host
	Host *Host
	// Secret
	Secret *Secret
}

// HTTPClient conf.
type HTTPClient struct {
	Dial      int64
	KeepAlive int64
}

// Host conf.
type Host struct {
	Geetest string
}

// Secret of Geetest
type Secret struct {
	CaptchaID  string
	PrivateKey string
}

// Init conf.
func Init() (err error) {
	bs, err := ioutil.ReadFile("config.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(bs, &Conf)

	geetestId := os.Getenv("GEETEST_ID")
	if geetestId == "" {
		panic("GEETEST_ID not set")
	}

	geetestKey := os.Getenv("GEETEST_KEY")
	if geetestKey == "" {
		panic("GEETEST_KEY not set")
	}

	Conf.Secret.CaptchaID = geetestId
	Conf.Secret.PrivateKey = geetestKey

	return
}
