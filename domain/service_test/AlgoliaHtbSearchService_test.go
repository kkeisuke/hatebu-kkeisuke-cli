package service_test

import (
	"strings"
	"testing"

	"github.com/kkeisuke/hatebu-kkeisuke-cli/domain/entity"
	"github.com/kkeisuke/hatebu-kkeisuke-cli/domain/service"
)

const defaultPerPage = 1

func TestAlgoliaHtbSearchService_Search(t *testing.T) {
	tests := []struct {
		name       string
		flg        entity.HtbSearchFlagEntity
		wantOutput string
		wantErr    bool
	}{
		{
			name:       "正常系",
			flg:        entity.HtbSearchFlagEntity{Freeword: "golang", PerPage: defaultPerPage},
			wantOutput: "golang",
			wantErr:    false,
		},
		{
			name:       "空文字",
			flg:        entity.HtbSearchFlagEntity{Freeword: "", PerPage: defaultPerPage},
			wantOutput: "",
			wantErr:    false,
		},
		{
			name:       "空白・改行文字",
			flg:        entity.HtbSearchFlagEntity{Freeword: " \n ", PerPage: defaultPerPage},
			wantOutput: "",
			wantErr:    false,
		},
	}

	htb := &service.AlgoliaHtbSearchService{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotOutput, err := htb.Search(tt.flg)

			if (err != nil) != tt.wantErr {
				t.Errorf("AlgoliaHtbSearchService.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !strings.Contains(strings.ToLower(gotOutput), strings.ToLower(tt.wantOutput)) {
				t.Errorf("AlgoliaHtbSearchService.Search() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
