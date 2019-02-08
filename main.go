package main

import (
	"github.com/alexandervantrijffel/goutil/logging"
	"github.com/alexandervantrijffel/scrape/config"
	"github.com/alexandervantrijffel/scrape/sourceycombinator"
)

func main() {
	logging.InitWith("scrape", config.DEBUG)
	sourceycombinator.RetrieveArticles()
}
