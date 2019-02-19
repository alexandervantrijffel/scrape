package config

import (
	"strings"

	"github.com/alexandervantrijffel/goutil/errorcheck"
	"github.com/alexandervantrijffel/goutil/logging"
	"github.com/caarlos0/env"
)

type Config struct {
	STREAMSDBCREDENTIALS string `env:"STREAMSDBCREDENTIALS"`
}

var THECONFIG Config

func InitMe() {
	THECONFIG = Config{}
	errorcheck.CheckPanic(env.Parse(&THECONFIG))
	logging.InitWith("scrape", DEBUG)
	logging.Infof("Configuration loaded DEBUG %t", DEBUG)
	if len(strings.TrimSpace(THECONFIG.STREAMSDBCREDENTIALS)) == 0 {
		logging.Fatal("Please provide credentials in the STREAMSDBCREDENTIALS environment variable. Format: user:password")
	}
}
