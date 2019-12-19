package service

import (
	"github.com/kkeisuke/hatebu-kkeisuke-cli/api/algolia"
	"github.com/kkeisuke/hatebu-kkeisuke-cli/domain/entity"
	"github.com/kkeisuke/hatebu-kkeisuke-cli/domain/value"
)

/*
AlgoliaHtbSearchService algolia はてブ検索
*/
type AlgoliaHtbSearchService struct {
}

/*
Search 検索
*/
func (*AlgoliaHtbSearchService) Search(flg entity.HtbSearchFlagEntity) (output string, err error) {

	htb := &algolia.HtbSearch{}
	htb.Setup()
	htb.Opt.HitsPerPage = flg.PerPage

	results, err := htb.Search(flg.Freeword)

	if err != nil {
		return
	}

	resultVal := value.AlgoliaHtbSearchResultValue{SearchResults: results}
	output, err = resultVal.ParseRawData(flg.Freeword)

	return
}
