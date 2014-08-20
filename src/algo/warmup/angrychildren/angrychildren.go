package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dataset struct {
	Children      int
	Packets       []int
	PacketsLength int
}

func ReadInput() (input string, err error) {
	const BUFFER_LENGTH = 255

	var (
		stdin  *os.File
		buffer []byte
		readNb int
	)

	stdin = os.Stdin

	for {
		buffer = make([]byte, BUFFER_LENGTH, BUFFER_LENGTH)

		readNb, err = stdin.Read(buffer)

		if err != nil {
			return
		}

		input = input + string(buffer)

		if readNb < BUFFER_LENGTH {
			return
		}
	}

	return
}

func Extract(source string) (dataset Dataset, err error) {
	var (
		rows        []string
		parsedValue int64
	)

	rows = strings.Split(source, "\n")

	parsedValue, err = strconv.ParseInt(rows[0], 10, 64)
	if err != nil {
		return
	} else {
		dataset.PacketsLength = int(parsedValue)
	}

	parsedValue, err = strconv.ParseInt(rows[1], 10, 64)
	if err != nil {
		return
	} else {
		dataset.Children = int(parsedValue)
	}

	dataset.Packets = make([]int, dataset.PacketsLength, dataset.PacketsLength)
	for i := 0; i < dataset.PacketsLength; i++ {

		parsedValue, err = strconv.ParseInt(rows[i+2], 10, 64)
		if err != nil {
			return
		} else {
			dataset.Packets[i] = int(parsedValue)
		}
	}

	return
}

func Diff(values ...int) (diff int) {
	var (
		curVal, min, max int
	)

	min = 9e9

	for _, curVal = range values {
		if curVal > max {
			max = curVal
		}
		if curVal < min {
			min = curVal
		}
	}

	diff = max - min

	return
}

func Solve(dataset Dataset) (min int) {
	// http://stackoverflow.com/questions/127704/algorithm-to-return-all-combinations-of-k-elements-from-n
	var (
		length, diff, i, j, k int
	)

	length = dataset.PacketsLength
	min = 9e9

	// TODO: Example for three, need generalization
	for i = 0; i < length-2; i++ {
		for j = i + 1; j < length-1; j++ {
			for k = j + 1; k < length; k++ {

				diff = Diff(dataset.Packets[i], dataset.Packets[j], dataset.Packets[k])
				if diff < min {
					min = diff
				}
			}
		}
	}

	return
}

func main() {
	var (
		input   string
		err     error
		dataset Dataset
		result  int
	)

	input, err = ReadInput()
	if err != nil {
		fmt.Println(err.Error())

	} else {
		dataset, err = Extract(input)
		if err != nil {
			fmt.Println(err.Error())

		} else {
			result = Solve(dataset)
			fmt.Println(result)
		}
	}
}
