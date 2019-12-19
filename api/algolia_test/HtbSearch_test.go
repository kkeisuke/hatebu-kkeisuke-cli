package algolia_test

import (
	"testing"

	"github.com/kkeisuke/hatebu-kkeisuke-cli/api/algolia"
)

const defaultPerPage = 31

func TestHtbSearch_Search(t *testing.T) {
	tests := []struct {
		name              string
		freeword          string
		wantSearchResults int
		wantErr           bool
	}{
		{
			name:              "正常系",
			freeword:          "golang",
			wantSearchResults: defaultPerPage,
			wantErr:           false,
		},
		{
			name:              "空文字",
			freeword:          "",
			wantSearchResults: 0,
			wantErr:           false,
		},
		{
			name:              "空白・改行文字",
			freeword:          " \n ",
			wantSearchResults: 0,
			wantErr:           false,
		},
	}

	htbSearch := &algolia.HtbSearch{}
	htbSearch.Opt.HitsPerPage = defaultPerPage
	htbSearch.Setup()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotSearchResults, err := htbSearch.Search(tt.freeword)
			if (err != nil) != tt.wantErr {
				t.Errorf("HtbSearch.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotSearchResults) != tt.wantSearchResults {
				t.Errorf("HtbSearch.Search() = %v, want %v", len(gotSearchResults), tt.wantSearchResults)
			}
		})
	}
}
