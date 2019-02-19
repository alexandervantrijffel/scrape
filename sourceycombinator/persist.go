package sourceycombinator

import (
	"fmt"
	"strings"

	"github.com/alexandervantrijffel/goutil/errorcheck"
	"github.com/alexandervantrijffel/goutil/jsonutil"
	"github.com/alexandervantrijffel/goutil/logging"
	"github.com/alexandervantrijffel/scrape/articledb"
	"github.com/alexandervantrijffel/scrape/config"
	"github.com/pkg/errors"
	sdb "github.com/streamsdb/driver/go/sdb"
)

func persistArticles(articles []Article, prevError error) ([]Article, error) {
	if prevError != nil {
		return nil, prevError
	}

	db := articledb.THEDB

	prefix := ""
	if config.DEBUG {
		prefix = "dev2-"
	}
	articlesStream := ArticlesStreamName()

	for _, a := range articles {
		a.ContentLink = strings.TrimRight(a.ContentLink, "/?")
		streamName := fmt.Sprintf("%sarticle-%s", prefix, a.ContentLink)
		slice, err := db.Read(streamName, 0, 1)
		_ = errorcheck.CheckLogf(err, "Failed to read from stream %s", streamName)
		if slice.Head != 0 { // existing stream?
			logging.Debugf("Skipping known article %s", a.ContentLink)
			continue
		}
		found := ArticleFound{Article: a}
		msg := sdb.MessageInput{Value: jsonutil.MarshalWithLogging(found)}

		logging.Debugf("Appending to stream %s %+v", streamName, found)
		_, err = db.Append(streamName, msg)
		_ = errorcheck.CheckLogf(err, "Failed to append article to stream %s. Data: %+v", streamName, found)

		logging.Debugf("Appending to stream %s", articlesStream)
		_, err = db.Append(ArticlesStreamName(), msg)
		_ = errorcheck.CheckLogf(err, "Failed to append article to stream %s. Data: %+v", streamName, found)
	}
	return articles, nil
}

type ArticleFound struct {
	Article Article
}

func ArticlesStreamName() string {
	prefix := ""
	if config.DEBUG {
		prefix = "dev2-"
	}
	return prefix + "articles-ycombinator"
}

func WatchFoundArticles() {
	errs := make(chan error)

	go func() {
		watch := articledb.THEDB.Watch(ArticlesStreamName(), 0, 10)
		for slice := range watch.Slices {
			for _, msg := range slice.Messages {
				println("received: ", string(msg.Value))
			}
		}

		errs <- errors.Wrap(watch.Err(), "watch error")
	}()

	logging.Fatal((<-errs).Error())
}
