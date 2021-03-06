package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func IsValid(str string) (valid bool, err error) {
	valid, err = regexp.MatchString("^[a-z]{0,3}\\d{2,8}[A-Z]{3}", str)

	return
}

func main() {
	var (
		valid       bool
		err         error
		linesNumber int
		lines       []string
	)

	if linesNumber, err = ReadInt(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	} else {
		// fmt.Println("Lines number:", linesNumber)

		if lines, err = ReadLines(linesNumber); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		} else {
			for i := 0; i < len(lines); i++ {
				// fmt.Println(lines[i])
				if valid, err = IsValid(lines[i]); err != nil {
					fmt.Fprintln(os.Stderr, err.Error())
				} else if valid {
					fmt.Println("VALID")
				} else {
					fmt.Println("INVALID")
				}
			}
		}
	}
}
