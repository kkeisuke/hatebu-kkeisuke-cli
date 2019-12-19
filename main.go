package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/kkeisuke/hatebu-kkeisuke-cli/domain/entity"
	"github.com/kkeisuke/hatebu-kkeisuke-cli/domain/service"
)

const (
	// ExitCodeOK 0
	ExitCodeOK int = iota
	// ExitCodeError 1
	ExitCodeError
)

// DefaultPerPage 検索件数 初期値
const DefaultPerPage = 31

func main() {
	os.Exit(Run(os.Args))
}

/*
Run 実行
*/
func Run(args []string) int {
	flgEntity := entity.HtbSearchFlagEntity{}
	flag.StringVar(&flgEntity.Freeword, "freeword", "", "search keyword")
	flag.StringVar(&flgEntity.Freeword, "f", "", "search keyword")
	flag.IntVar(&flgEntity.PerPage, "perPage", DefaultPerPage, "per page")
	flag.IntVar(&flgEntity.PerPage, "p", DefaultPerPage, "per page")
	flag.Parse()

	if strings.TrimSpace(flgEntity.Freeword) == "" {
		fmt.Fprintf(os.Stderr, "empty freeword \n\n htb -f <freeword>\n\n")
		return ExitCodeError
	}

	htb := service.AlgoliaHtbSearchService{}
	output, err := htb.Search(flgEntity)

	if err != nil {
		return ExitCodeError
	}

	fmt.Fprintf(os.Stdout, output)

	return ExitCodeOK
}
