package main

import (
	"github.com/alexandervantrijffel/scrape/articledb"
	"github.com/alexandervantrijffel/scrape/config"
	"github.com/alexandervantrijffel/scrape/sourceycombinator"
)

func main() {
	config.InitMe()
	articledb.Connect()

	_, _ = sourceycombinator.RetrieveArticles()

	sourceycombinator.WatchFoundArticles()
}
