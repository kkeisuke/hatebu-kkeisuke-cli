package value

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/kkeisuke/hatebu-kkeisuke-cli/domain/entity"
)

var (
	yellow = "33"
	cyan   = "36"
)

/*
AlgoliaHtbSearchResultValue algolia はてブ検索結果
*/
type AlgoliaHtbSearchResultValue struct {
	SearchResults []entity.AlgoliaHtbSearchResultEntity
}

/*
ParseRawData 検索結果（1日分のブックマーク）から該当項目のみ抽出します
*/
func (val *AlgoliaHtbSearchResultValue) ParseRawData(freeword string) (result string, err error) {
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

	for _, searchResult := range val.SearchResults {
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
