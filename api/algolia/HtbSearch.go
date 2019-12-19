package algolia

import (
	"fmt"
	"os"
	"strings"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/kkeisuke/hatebu-kkeisuke-cli/domain/entity"
)

var (
	application string
	apiKey      string
	indexName   string
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
func (htbSearch *HtbSearch) Search(freeword string) (searchResults []entity.AlgoliaHtbSearchResultEntity, err error) {
	if strings.TrimSpace(freeword) == "" {
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
