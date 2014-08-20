package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadInt() (result int, err error) {
	var (
		parsedValue int64
		readValue   string
		readSize    int
	)

	if readSize, err = fmt.Scan(&readValue); err == nil {
		if readSize <= 0 {
			err = errors.New("No data to read")
		} else {
			if parsedValue, err = strconv.ParseInt(readValue, 10, 64); err == nil {
				result = int(parsedValue)
			}
		}
	}

	return
}

func ReadLines(limit int) (result []string, err error) {
	var (
		scanner = bufio.NewScanner(os.Stdin)
	)

	result = make([]string, limit, limit)

	for i := 0; i < limit; i++ {
		scanner.Scan()
		result[i] = scanner.Text()
		if err = scanner.Err(); err != nil {
			return
		}
	}

	return
}

func main() {
	var (
		err                    error
		lines                  []string
		linesLen, patternFound int
		loweredLine            string
		searchRegExpt          *regexp.Regexp = regexp.MustCompile("hackerrank")
	)

	if linesLen, err = ReadInt(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	} else {
		//fmt.Println("Lines number:", linesLen)
	}

	if lines, err = ReadLines(linesLen); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	} else {
		//fmt.Println("Lines:")
		for i := 0; i < len(lines); i++ {
			loweredLine = strings.ToLower(lines[i])
			//fmt.Println(" ", loweredLine)
			patternFound += len(searchRegExpt.FindAllString(loweredLine, -1))
		}
		fmt.Println(patternFound)
	}
}
