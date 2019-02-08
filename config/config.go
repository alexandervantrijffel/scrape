package config

import (
	"github.com/alexandervantrijffel/goutil/errorcheck"
	"github.com/alexandervantrijffel/goutil/logging"
	"github.com/caarlos0/env"
)

type Config struct {
	STREAMSDBCREDENTIALS string `env:"STREAMSDBCREDENTIALS"`
}

var THECONFIG Config

func init() {
	THECONFIG = Config{
		""}
	errorcheck.CheckPanic(env.Parse(&THECONFIG))
	logging.InitWith("scrape", DEBUG)
	logging.Infof("Configuration loaded DEBUG %t", DEBUG)
}
