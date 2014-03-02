package main

import (
	"errors"
	"fmt"
	"strconv"
)

type IDataset interface {
	Load() error
	Display()
	Solve()
	PrintResult()
}

type Dataset struct {
	Length int
	Values []int
	Result []int
}

func (d *Dataset) PrintResult() {
	for i := 0; i < len(d.Result); i++ {
		fmt.Println(d.Result[i])
	}
}

func (d *Dataset) solveCase(value int) (result int) {
	result = 1

	for i := 1; i <= value; i++ {
		if i%2 == 0 {
			result = result + 1
		} else {
			result = result * 2
		}
	}

	return
}

func (d *Dataset) Solve() {
	d.Result = make([]int, d.Length, d.Length)

	for i := 0; i < d.Length; i++ {
		d.Result[i] = d.solveCase(d.Values[i])
	}
}

func (d *Dataset) Display() {
	fmt.Println("Dataset:")
	fmt.Println("  Length:", d.Length)
	fmt.Println("  Values:")
	for i := 0; i < len(d.Values); i++ {
		fmt.Println("   -", d.Values[i])
	}
}

func (d *Dataset) Load() (err error) {
	var (
		readValue   string
		readSize    int
		parsedValue int64
	)

	readSize, err = fmt.Scanln(&readValue)
	if err != nil {
		return
	} else if readSize < 1 {
		err = errors.New("Can't read input")
		return
	} else {
		parsedValue, err = strconv.ParseInt(readValue, 10, 64)
		if err != nil {
			return
		} else {
			d.Length = int(parsedValue)
		}
	}

	d.Values = make([]int, d.Length, d.Length)

	for i := 0; i < len(d.Values); i++ {
		readSize, err = fmt.Scanln(&readValue)
		if err != nil {
			return
		} else if readSize < 1 {
			err = errors.New("Can't read input")
			return
		} else {
			parsedValue, err = strconv.ParseInt(readValue, 10, 64)
			if err != nil {
				return
			} else {
				d.Values[i] = int(parsedValue)
			}
		}
	}

	return
}

func NewDataset() IDataset {
	return new(Dataset)
}

func main() {
	var (
		dataset IDataset
		err     error
	)

	dataset = NewDataset()

	err = dataset.Load()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		dataset.Solve()
		dataset.PrintResult()
	}
}
