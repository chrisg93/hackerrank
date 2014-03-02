package main

import (
	"errors"
	"fmt"
	"strconv"
)

type IDataset interface {
	Load() error
	Display()
	Resolve()
	PrintSolution()
}

type Dataset struct {
	Length int
	Values []int
	Result []string
}

func (d *Dataset) Load() (err error) {
	var (
		parsedValue int64
		readLength  int
		readValue   string
	)

	readLength, err = fmt.Scanln(&readValue)
	if err != nil {
		return
	} else if readLength < 1 {
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
		readLength, err = fmt.Scanln(&readValue)
		if err != nil {
			return
		} else if readLength < 1 {
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

func (d *Dataset) Display() {
	fmt.Println("Dataset:")
	fmt.Println("  Length:", d.Length)
	fmt.Println("  Values:")
	for i := 0; i < len(d.Values); i++ {
		fmt.Println("   -", d.Values[i])
	}
}

func getFiboList(maxValue int) (fiboList []int) {
	var (
		a         int = 0
		b         int = 1
		nextValue int
	)

	fiboList = make([]int, 0, 0)

	fiboList = append(fiboList, a)
	fiboList = append(fiboList, b)

	for {
		nextValue = a + b
		if nextValue > maxValue {
			break
		} else {
			fiboList = append(fiboList, nextValue)
			a, b = b, nextValue
		}
	}

	return
}

func (d *Dataset) Resolve() {
	var (
		FiboList []int
	)

	FiboList = getFiboList(10000000000)

	d.Result = make([]string, d.Length, d.Length)

	for i := 0; i < len(d.Values); i++ {
		for j := 0; j < len(FiboList); j++ {
			if d.Values[i] == FiboList[j] {
				d.Result[i] = "IsFibo"
			} else if d.Values[i] < FiboList[j] {
				break
			}
		}
		if d.Result[i] != "IsFibo" {
			d.Result[i] = "IsNotFibo"
		}
	}
}

func (d *Dataset) PrintSolution() {
	for i := 0; i < len(d.Result); i++ {
		fmt.Println(d.Result[i])
	}
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
		dataset.Resolve()
		dataset.PrintSolution()
	}
}
