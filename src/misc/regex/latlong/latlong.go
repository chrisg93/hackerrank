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

func Check(source string) bool {
	var (
		dataExtractorRegex *regexp.Regexp = regexp.MustCompile("\\(([\\+\\-]?\\d+\\.?\\d*),\\s([\\+\\-]?\\d+\\.?\\d*)\\)")
		extract            []string
		lat, long          float64
		err                error
		matched            bool
	)

	extract = dataExtractorRegex.FindStringSubmatch(source)

	if lat, err = strconv.ParseFloat(extract[1], 64); err != nil {
		panic(err.Error())
	}

	if long, err = strconv.ParseFloat(extract[2], 64); err != nil {
		panic(err.Error())
	}

	// fmt.Println("    Lat:", lat)
	// fmt.Println("    Long:", long)

	if matched, err = regexp.MatchString("\\.$", extract[1]); err != nil {
		panic(err.Error())
	} else if matched {
		return false
	}

	if matched, err = regexp.MatchString("\\.$", extract[2]); err != nil {
		panic(err.Error())
	} else if matched {
		return false
	}

	if matched, err = regexp.MatchString("^[\\+\\-]?0", extract[1]); err != nil {
		panic(err.Error())
	} else if matched {
		return false
	}

	if matched, err = regexp.MatchString("^[\\+\\-]?0", extract[2]); err != nil {
		panic(err.Error())
	} else if matched {
		return false
	}

	if lat < -90.0 || long < -180.0 || lat > 90.0 || long > 180.0 {
		return false
	}

	return true
}

func main() {
	var (
		err      error
		lines    []string
		linesLen int
	)

	if linesLen, err = ReadInt(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	} else {
		// fmt.Println("Lines number:", linesLen)
	}

	if lines, err = ReadLines(linesLen); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	} else {
		// fmt.Println("Lines:")
		for i := 0; i < len(lines); i++ {
			// fmt.Println(" ", lines[i])
			if Check(lines[i]) {
				fmt.Println("Valid")
			} else {
				fmt.Println("Invalid")
			}
		}
	}
}
