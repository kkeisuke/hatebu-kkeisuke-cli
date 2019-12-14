package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kkeisuke/hatebu-kkeisuke-cli/api/algolia"
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
	var freeword string
	flag.StringVar(&freeword, "freeword", "", "search keyword")
	flag.StringVar(&freeword, "f", "", "search keyword")
	var perPage int
	flag.IntVar(&perPage, "perPage", DefaultPerPage, "per page")
	flag.IntVar(&perPage, "p", DefaultPerPage, "per page")
	flag.Parse()

	if freeword == "" {
		fmt.Fprintf(os.Stderr, "empty freeword \n\n htb -f <freeword>\n\n")
		return ExitCodeError
	}

	htb := &algolia.HtbSearch{}
	htb.Setup()
	htb.Opt.HitsPerPage = perPage

	results, err := htb.Search(freeword)

	if err != nil {
		return ExitCodeError
	}

	output, err := htb.ParseRawData(freeword, results)

	if err != nil {
		return ExitCodeError
	}

	fmt.Fprintf(os.Stdout, output)

	return ExitCodeOK
}
