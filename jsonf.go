package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln(r)
		}
	}()

	var opts JsonfOptions
	HandleArgs(&opts)
	Jsonf(opts)
}

type JsonfOptions struct {
	decoder     *json.Decoder
	output      io.Writer
	keyWrap     string
	arrWrap     string
	seperator   string
	includeKeys bool
}

func HandleArgs(opts *JsonfOptions) {
	var inputFile string
	flag.StringVar(&inputFile, "i", "", "The input file. Defaults to stdin.")
	var outputFile string
	flag.StringVar(&outputFile, "o", "", "The output file. Defaults to stdout.")
	flag.StringVar(&opts.keyWrap, "k", "{}", "The runes used to wrap output keys.")
	flag.StringVar(&opts.arrWrap, "a", "[]", "The runes used to wrap output array indexes.")
	flag.StringVar(&opts.seperator, "sep", ">", "String used to seperate keys.")
	flag.BoolVar(&opts.includeKeys, "keys", false, "If true, keys are printed on their own line.")

	flag.Parse()

	if len(inputFile) == 0 {
		opts.decoder = json.NewDecoder(os.Stdin)
	} else if input, err := os.Open(inputFile); err == nil {
		opts.decoder = json.NewDecoder(input)
	} else {
		panic(err)
	}

	if len(outputFile) == 0 {
		opts.output = os.Stdout
	} else if output, err := os.Open(outputFile); err == nil {
		opts.output = output
	} else {
		panic(err)
	}

	if len(opts.keyWrap) != 2 {
		panic("-k must contain exactly two characters")
	}
	if len(opts.arrWrap) != 2 {
		panic("-a must contain exactly two characters")
	}
}

func Jsonf(opts JsonfOptions) {
	for numReads := 0; ; numReads++ {
		var json interface{}
		if err := opts.decoder.Decode(&json); err != nil {
			if err.Error() != "EOF" {
				panic(err)
			} else {
				return
			}
		}
		ProcessJsonValue(opts, strconv.Itoa(numReads), json)
	}
}

func ProcessJsonValue(opts JsonfOptions, leadKey string, json interface{}) {
	switch json.(type) {
	case map[string]interface{}:
		for key, value := range json.(map[string]interface{}) {
			keyLead := fmt.Sprintf("%s%s%c%s%c", leadKey, opts.seperator, opts.keyWrap[0], key, opts.keyWrap[1])
			if opts.includeKeys {
				fmt.Fprintln(opts.output, keyLead)
			}
			ProcessJsonValue(opts, keyLead, value)
		}
	case []interface{}:
		if opts.includeKeys {
			fmt.Fprintf(opts.output, "%s%s%c%c\n", leadKey, opts.seperator, opts.arrWrap[0], opts.arrWrap[1])
		}
		for index, value := range json.([]interface{}) {
			elemLead := fmt.Sprintf("%s%s%c%d%c", leadKey, opts.seperator, opts.arrWrap[0], index, opts.arrWrap[1])
			ProcessJsonValue(opts, elemLead, value)
		}
	case string:
		fmt.Fprintf(opts.output, "%s%s%s\n", leadKey, opts.seperator, strconv.Quote(json.(string)))
	default:
		fmt.Fprintf(opts.output, "%s%s%v\n", leadKey, opts.seperator, json)
	}
}
