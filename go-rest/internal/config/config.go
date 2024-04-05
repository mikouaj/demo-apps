package config

import (
	"fmt"
	"os"
	"strconv"
)

const (
	TLSEnvVar     = "LISTEN_TLS"
	TLSCertEnvVar = "TLS_CERT"
	TLSKeyEnvVar  = "TLS_KEY"
)

type Config struct {
	ListenTLS   bool
	TLSCertPath string
	TLSKeyPath  string
}

func LoadFromEnv() (*Config, error) {
	c := &Config{
		ListenTLS:   false,
		TLSCertPath: "server.crt",
		TLSKeyPath:  "server.key",
	}
	if v, ok := os.LookupEnv(TLSEnvVar); ok {
		if b, err := strconv.ParseBool(v); err == nil {
			c.ListenTLS = b
		} else {
			return nil, fmt.Errorf("invalid variable %s value: has to be bool", TLSEnvVar)
		}
	}
	if v, ok := os.LookupEnv(TLSCertEnvVar); ok {
		c.TLSCertPath = v
	}
	if v, ok := os.LookupEnv(TLSKeyEnvVar); ok {
		c.TLSKeyPath = v
	}
	return c, nil
}
