package value_test

import (
	"testing"

	"github.com/kkeisuke/hatebu-kkeisuke-cli/domain/entity"
	"github.com/kkeisuke/hatebu-kkeisuke-cli/domain/value"
)

func TestAlgoliaHtbSearchResultValue_ParseRawData(t *testing.T) {
	tests := []struct {
		name          string
		SearchResults []entity.AlgoliaHtbSearchResultEntity
		freeword      string
		wantResult    string
		wantErr       bool
	}{
		{
			name: "正常系",
			SearchResults: []entity.AlgoliaHtbSearchResultEntity{
				{
					Content:  "#### タイトル kkeisuke \nhttps://kkeisuke.com<br>\n2019/12/19 00:00:00<br>\nkkeisuke",
					ObjectID: "2019-12-19",
					Path:     "/posts/2019-12-19",
				},
			},
			freeword:   "kkeisuke",
			wantResult: "#### タイトル \x1b[33mkkeisuke\x1b[0m \n\x1b[36mhttps://kkeisuke.com<br>\x1b[0m\n2019/12/19 00:00:00<br>\n\x1b[33mkkeisuke\x1b[0m\n\n",
			wantErr:    false,
		},
		{
			name: "空文字",
			SearchResults: []entity.AlgoliaHtbSearchResultEntity{
				{
					Content:  "",
					ObjectID: "",
					Path:     "",
				},
			},
			freeword:   "",
			wantResult: "",
			wantErr:    false,
		},
		{
			name: "空白・改行文字",
			SearchResults: []entity.AlgoliaHtbSearchResultEntity{
				{
					Content:  "",
					ObjectID: "",
					Path:     "",
				},
			},
			freeword:   " \n ",
			wantResult: "",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val := &value.AlgoliaHtbSearchResultValue{
				SearchResults: tt.SearchResults,
			}
			gotResult, err := val.ParseRawData(tt.freeword)
			if (err != nil) != tt.wantErr {
				t.Errorf("AlgoliaHtbSearchResultValue.ParseRawData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("AlgoliaHtbSearchResultValue.ParseRawData() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
