package libs

import (
	"os"
)

type Config struct {
	ServerPort        string
	DatabasePort      string
	DatabaseUser      string
	DatabasePassword  string
	DatabaseHost      string
	DatabaseName      string
	PaymentNamespaces string
	MemberNamespaces  string
	ProductNamespaces string
}

// NewEnvConfig function to get info in current config file
// @param name string
// TODO : get env
func NewEnvConfig() (*Config, error) {
	env := &Config{
		ServerPort:        os.Getenv("ServerPort"),
		DatabasePort:      os.Getenv("DatabasePort"),
		DatabaseUser:      os.Getenv("DatabaseUser"),
		DatabasePassword:  os.Getenv("DatabasePassword"),
		DatabaseHost:      os.Getenv("DatabaseHost"),
		DatabaseName:      os.Getenv("DatabaseName"),
		PaymentNamespaces: os.Getenv("PaymentNamespaces"),
		MemberNamespaces:  os.Getenv("MemberNamespaces"),
		ProductNamespaces: os.Getenv("ProductNamespaces"),
	}
	return env, nil
}
