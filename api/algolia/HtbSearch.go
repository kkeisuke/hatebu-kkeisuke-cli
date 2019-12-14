package algolia

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

var (
	application string
	apiKey      string
	indexName   string
)

var (
	yellow = "33"
	cyan   = "36"
)

/*
HtbSearch はてなブックマーク Index
*/
type HtbSearch struct {
	Index *search.Index
	Opt   struct {
		HitsPerPage int
	}
}

/*
SearchResult 検索結果
*/
type SearchResult struct {
	Content  string
	ObjectID string
	Path     string
}

/*
Setup APIキーなどのセットアップを行います
*/
func (htbSearch *HtbSearch) Setup() {
	if application == "" {
		application = os.Getenv("ALGOLIA_APPLICATION")
	}
	if apiKey == "" {
		apiKey = os.Getenv("ALGOLIA_API_KEY")
	}
	if indexName == "" {
		indexName = os.Getenv("ALGOLIA_INDEX")
	}
	client := search.NewClient(application, apiKey)
	htbSearch.Index = client.InitIndex(indexName)
}

/*
Search 検索
*/
func (htbSearch *HtbSearch) Search(freeword string) (searchResults []SearchResult, err error) {
	if freeword == "" {
		return
	}

	res, err := htbSearch.Index.Search(freeword, opt.HitsPerPage(htbSearch.Opt.HitsPerPage))

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n\n")
		return
	}

	err = res.UnmarshalHits(&searchResults)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n\n")
	}

	return
}

/*
ParseRawData 検索結果（1日分のブックマーク）から該当項目のみ抽出します
*/
func (htbSearch *HtbSearch) ParseRawData(freeword string, searchResults []SearchResult) (result string, err error) {
	if freeword == "" {
		return
	}

	var results []string
	freewordColor := ("\x1b[" + yellow + "m") + "$1" + "\x1b[0m"
	freewordRegExp, err := regexp.Compile(`(?i)(` + freeword + `)`)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n\n")
		return
	}

	for _, searchResult := range searchResults {
		// 段落
		paragraphs := strings.Split(searchResult.Content, "\n\n")
		for _, paragraph := range paragraphs {
			// 段落にフリーワードが含まれるか
			if strings.Contains(strings.ToLower(paragraph), strings.ToLower(freeword)) {
				// 1行毎に分割
				lines := strings.Split(strings.TrimSpace(paragraph), "\n")
				// タイトル
				lines[0] = freewordRegExp.ReplaceAllString(lines[0], freewordColor)
				// URL
				lines[1], _ = url.QueryUnescape(lines[1])
				lines[1] = ("\x1b[" + cyan + "m") + lines[1] + "\x1b[0m"
				// コメント
				if len(lines) >= 4 {
					lines[3] = freewordRegExp.ReplaceAllString(lines[3], freewordColor)
				}
				// 後ろから入れる
				results = append([]string{strings.Join(lines, "\n")}, results...)
			}
		}
	}
	result = strings.Join(results, "\n\n") + "\n\n"

	return
}
