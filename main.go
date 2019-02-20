package main

import (
	"flag"

	"github.com/alexandervantrijffel/scrape/articledb"
	"github.com/alexandervantrijffel/scrape/config"
	"github.com/alexandervantrijffel/scrape/sourceycombinator"
)

func main() {
	watch := flag.Bool("w", false, "watch events")
	flag.Parse()

	config.InitMe()
	articledb.Connect()

	_, _ = sourceycombinator.RetrieveArticles()

	if *watch {
		sourceycombinator.WatchFoundArticles()
	}
}
