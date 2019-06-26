package conf

import (
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
	host := getEnv("GEETEST_HOST",  "http://api.geetest.com")



	geetestId := os.Getenv("GEETEST_ID")

	if geetestId == "" {
		panic("GEETEST_ID not set")
	}

	geetestKey := os.Getenv("GEETEST_KEY")

	if geetestKey == "" {
		panic("GEETEST_KEY not set")
	}

	Conf = &Config{Host: &Host{Geetest: host}, HTTPClient: &HTTPClient{Dial: 1, KeepAlive: 1}, Secret: &Secret{CaptchaID: geetestId, PrivateKey: geetestKey}}

	return
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}