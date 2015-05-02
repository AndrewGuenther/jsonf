package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
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
	decoder   *json.Decoder
	output    io.Writer
	keyWrap   string
	strWrap   string
	arrWrap   string
	seperator string
}

func HandleArgs(opts *JsonfOptions) {
	var inputFile string
	flag.StringVar(&inputFile, "i", "", "The input file. Defaults to stdin.")
	var outputFile string
	flag.StringVar(&outputFile, "o", "", "The output file. Defaults to stdout.")
	flag.StringVar(&opts.keyWrap, "k", "{}", "The runes used to wrap output keys.")
	flag.StringVar(&opts.strWrap, "s", "\"\"", "The runes used to wrap output strings.")
	flag.StringVar(&opts.arrWrap, "a", "[]", "The runes used to wrap output array indexes.")
	flag.StringVar(&opts.seperator, "sep", "->", "String used to seperate keys.")

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
	if len(opts.strWrap) != 2 {
		panic("-s must contain exactly two characters")
	}
	if len(opts.arrWrap) != 2 {
		panic("-a must contain exactly two characters")
	}
}

func Jsonf(opts JsonfOptions) {
	for {
		var json interface{}
		if err := opts.decoder.Decode(&json); err != nil {
			if err.Error() != "EOF" {
				panic(err)
			} else {
				return
			}
		}
		ProcessJsonValue(opts, "", json)
	}
}

func ProcessJsonValue(opts JsonfOptions, leadKey string, json interface{}) {
	switch json.(type) {
	case map[string]interface{}:
		for key, value := range json.(map[string]interface{}) {
			keyLead := fmt.Sprintf("%s%s%c%s%c", leadKey, opts.seperator, opts.keyWrap[0], key, opts.keyWrap[1])
			fmt.Fprintln(opts.output, keyLead)
			ProcessJsonValue(opts, keyLead, value)
		}
	case []interface{}:
		fmt.Fprintf(opts.output, "%s%s%c%c\n", leadKey, opts.seperator, opts.arrWrap[0], opts.arrWrap[1])
		for index, value := range json.([]interface{}) {
			elemLead := fmt.Sprintf("%s%s%c%d%c", leadKey, opts.seperator, opts.arrWrap[0], index, opts.arrWrap[1])
			ProcessJsonValue(opts, elemLead, value)
		}
	case string:
		fmt.Fprintf(opts.output, "%s%s%c%s%c\n", leadKey, opts.seperator, opts.strWrap[0], json, opts.strWrap[1])
	default:
		fmt.Fprintf(opts.output, "%s%s%v\n", leadKey, opts.seperator, json)
	}
}
