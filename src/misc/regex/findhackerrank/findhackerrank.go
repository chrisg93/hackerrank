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

func main() {
	var (
		beginWith, endWith bool
		err                error
		lines              []string
		linesLen           int
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
			//fmt.Println(" ", lines[i])
			if beginWith, err = regexp.MatchString("^hackerrank", lines[i]); err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
			if endWith, err = regexp.MatchString("hackerrank$", lines[i]); err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
			if beginWith && endWith {
				fmt.Println("0")
			} else if beginWith {
				fmt.Println("1")
			} else if endWith {
				fmt.Println("2")
			} else {
				fmt.Println("-1")
			}
		}
	}
}
