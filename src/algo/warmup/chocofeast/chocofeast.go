package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	Money    int
	Price    int
	Discount int
}

type Dataset struct {
	Count int
	Items []Item
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

func BuildSet(dataset Dataset, lines []string, i int) (item Item, err error) {
	var (
		chars       []string
		parsedValue int64
	)

	chars = strings.Split(lines[i+1], " ")

	parsedValue, err = strconv.ParseInt(chars[0], 10, 64)
	if err != nil {
		return
	} else {
		item.Money = int(parsedValue)
	}

	parsedValue, err = strconv.ParseInt(chars[1], 10, 64)
	if err != nil {
		return
	} else {
		item.Price = int(parsedValue)
	}

	parsedValue, err = strconv.ParseInt(chars[2], 10, 64)
	if err != nil {
		return
	} else {
		item.Discount = int(parsedValue)
	}

	return
}

func BuildDataset(source string) (dataset Dataset, err error) {
	var (
		lines       []string
		parsedValue int64
	)

	lines = strings.Split(source, "\n")

	parsedValue, err = strconv.ParseInt(lines[0], 10, 64)
	if err != nil {
		return
	} else {
		dataset.Count = int(parsedValue)
	}

	dataset.Items = make([]Item, dataset.Count, dataset.Count)
	for i := 0; i < int(dataset.Count); i++ {
		dataset.Items[i], err = BuildSet(dataset, lines, i)
		if err != nil {
			return
		}
	}

	return
}

func Solve(dataset Dataset, index int) (chcocolatesEaten int) {
	var (
		price          int = dataset.Items[index].Price
		discount       int = dataset.Items[index].Discount
		curMoney       int = dataset.Items[index].Money
		curDiscount    int = 0
		curChcocolates int = 0
	)

	for {

		if curMoney >= price {
			curMoney = curMoney - price
			curChcocolates = curChcocolates + 1
			continue
		}

		if curChcocolates > 0 {
			chcocolatesEaten = chcocolatesEaten + 1
			curChcocolates = curChcocolates - 1
			curDiscount = curDiscount + 1
			continue
		}

		if curDiscount >= discount {
			curChcocolates = curChcocolates + 1
			curDiscount = curDiscount - discount
			continue
		}

		return
	}
}

func main() {
	var (
		dataset Dataset
		err     error
		input   string
		results []int
		result  int
	)

	input, err = ReadInput()
	if err != nil {
		fmt.Println(err.Error())

	} else {

		dataset, err = BuildDataset(input)
		if err != nil {
			fmt.Println(err.Error())

		} else {
			results = make([]int, dataset.Count, dataset.Count)
			for i := 0; i < dataset.Count; i++ {
				results[i] = Solve(dataset, i)
			}

			for _, result = range results {
				fmt.Println(result)
			}
		}

	}
}
