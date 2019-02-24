package sourceycombinator

import (
	"strings"
	"time"

	"github.com/alexandervantrijffel/goutil/errorcheck"
	"github.com/alexandervantrijffel/goutil/jsonutil"
	"github.com/alexandervantrijffel/goutil/stringutil"
	"github.com/alexandervantrijffel/scrape/scraping"
)

type ArticleList struct {
	Articles []struct {
		ContentLink string `json:"contentLink"`
		Title       string `json:"title"`
	} `json:"articles"`
	Subtexts []struct {
		Comments        string `json:"comments"`
		Score           string `json:"score"`
		YcombinatorLink string `json:"ycombinatorLink"`
	} `json:"subtexts"`
}

func RetrieveArticles() ([]Article, error) {
	result, err := scraping.Get(`
LET doc = DOCUMENT('https://news.ycombinator.com/news?p=0', false)
LET articles = (
  FOR article IN ELEMENTS(doc, '.athing')
		LET storyLink = ELEMENT(article, '.title .storylink')
		RETURN {
			title: storyLink.innerText,
			contentLink: storyLink.attributes.href,
		}
)

LET subtexts = (
  FOR subtext IN ELEMENTS(doc, '.subtext')
    LET ycombinatorLink = LAST(subtext.children)
    RETURN {
      score: INNER_TEXT(subtext, '.score'),
      ycombinatorLink: ycombinatorLink.attributes.href,
      comments: ycombinatorLink.innerText
    }
)

RETURN {
  articles: articles,
  subtexts: subtexts
}`)
	if errorcheck.CheckLogf(err, "Failed to scrape") != nil {
		return nil, err
	}
	return persistArticles(topArticles(jsonToArticles(result)))
}

func topArticles(articles []Article, prevError error) ([]Article, error) {
	if prevError != nil {
		return articles, prevError
	}
	minScore := 150
	var top []Article
	for _, a := range articles {
		if a.Score >= minScore {
			top = append(top, a)
		}
	}
	return top, nil
}

type Article struct {
	Title           string
	ContentLink     string
	Score           int
	Comments        int
	YcombinatorLink string
	FoundAt         time.Time
}

func jsonToArticles(json []byte) ([]Article, error) {
	var articles ArticleList
	if err := jsonutil.UnmarshalWithLogging(&articles, json); err != nil {
		return nil, err
	}
	var result []Article
	for n, art := range articles.Articles {
		subtext := articles.Subtexts[n]
		comments := strings.TrimSpace(remove(remove(subtext.Comments, "comments"), "comment"))
		if comments == "hide" {
			// ads don't have an ycombinator link or comments
			continue
		}
		commentCount := 0
		if comments != "discuss" {
			commentCount = stringutil.AtoiWithLogging(comments)
		}
		score := strings.TrimSpace(remove(subtext.Score, "points"))
		result = append(result, Article{
			Title:           art.Title,
			Score:           stringutil.AtoiWithLogging(score),
			ContentLink:     art.ContentLink,
			Comments:        commentCount,
			YcombinatorLink: subtext.YcombinatorLink,
			FoundAt:         time.Now().UTC(),
		})
	}
	return result, nil
}

func remove(source string, toRemove string) string {
	return strings.Replace(source, toRemove, "", -1)
}
