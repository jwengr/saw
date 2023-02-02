package config

import (
	"github.com/influxdata/telegraf"
)

type AWSConfiguration struct {
	Endpoint string
	Region   string
	Profile  string

	Log telegraf.Logger `toml:"-"`
}
