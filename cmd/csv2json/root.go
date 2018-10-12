package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/willchertoff/csv2json/pkg/converter"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	flag.Parse()

	var data []byte
	var err error
	var separator string

	switch {
	case flag.NArg() < 1:
		data, err = ioutil.ReadAll(os.Stdin)
		separator = "comma"
		check(err)
		break
	case flag.NArg() == 1:
		data, err = ioutil.ReadFile(flag.Arg(0))
		separator = "comma"
		check(err)
		break
	case flag.NArg() > 1:
		data, err = ioutil.ReadFile(flag.Arg(0))
		separator = flag.Arg(1)
		check(err)
		break
	default:
		fmt.Printf("input must be from stdin or file\n")
		os.Exit(1)
	}

	jsonBytes := converter.ConvertCSVToJSON(data, Separator(separator))

	fmt.Print(string(jsonBytes))
}

func Separator(arg string) rune {
	var sep rune
	switch arg {
	case "comma":
		sep = ','
	case "tab":
		sep = '\t'
	case "":
		sep = ','
	default:
		panic("unrecognized separtator, accepted values are: tab, comma")
	}
	return sep
}
