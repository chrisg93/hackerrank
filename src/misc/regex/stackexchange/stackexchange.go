package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const ID_EXTRACTOR_REGEX = "question-summary-(\\d+)"
const TITLE_EXTRACTOR_REGEX = "class=\"question-hyperlink\">([\\w\\s]+)</a>"
const TIME_AGO_EXTRACTOR_REGEX = "class=\"relativetime\">([\\w\\s]+)</span>"
const CLEANER_REGEX = "[\\r\\n\\t\\s]+"

func ReadInput() (readInput string, err error) {
	var scanner *bufio.Scanner

	scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		readInput = readInput + scanner.Text()
		if err = scanner.Err(); err != nil {
			return
		}
	}

	return
}

func main() {
	var (
		err                   error
		readInput             string
		ids, titles, timesAgo []string
		regexResult           [][]string
		re                    *regexp.Regexp
	)

	ids = make([]string, 0, 0)

	if readInput, err = ReadInput(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	re = regexp.MustCompile(CLEANER_REGEX)
	readInput = re.ReplaceAllString(readInput, " ")

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println(readInput)

	// fmt.Println("Read input:")
	// fmt.Println(readInput)

	re = regexp.MustCompile(ID_EXTRACTOR_REGEX)
	regexResult = re.FindAllStringSubmatch(readInput, -1)
	for i := 0; i < len(regexResult); i++ {
		if len(regexResult[i]) != 2 {
			fmt.Fprintln(os.Stderr, "ERROR PARSING IDS")
			return
		}
		ids = append(ids, regexResult[i][1])
	}

	re = regexp.MustCompile(TITLE_EXTRACTOR_REGEX)
	regexResult = re.FindAllStringSubmatch(readInput, -1)
	for i := 0; i < len(regexResult); i++ {
		if len(regexResult[i]) != 2 {
			fmt.Fprintln(os.Stderr, "ERROR PARSING TITLES")
			return
		}
		titles = append(titles, regexResult[i][1])
	}

	re = regexp.MustCompile(TIME_AGO_EXTRACTOR_REGEX)
	regexResult = re.FindAllStringSubmatch(readInput, -1)
	for i := 0; i < len(regexResult); i++ {
		if len(regexResult[i]) != 2 {
			fmt.Fprintln(os.Stderr, "ERROR PARSING TIMES AGO")
			return
		}
		timesAgo = append(timesAgo, regexResult[i][1])
	}

	// fmt.Println("Ids:", ids)
	// fmt.Println("Titles:", titles)
	// fmt.Println("Times ago:", timesAgo)

	if len(ids) != len(titles) || len(titles) != len(timesAgo) {
		fmt.Fprintln(os.Stderr, "INCORRECT NUMBER OF RESULT")
	} else {
		for i := 0; i < len(ids); i++ {
			fmt.Printf("%s;%s;%s\n", ids[i], titles[i], timesAgo[i])
		}
	}
}
