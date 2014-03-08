package main

import (
	"errors"
	"fmt"
	"strconv"
)

func readNextInt() (val int, err error) {
	var (
		parsedValue int64
		readValue   string
		readSize    int
	)

	readSize, err = fmt.Scan(&readValue)
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
			val = int(parsedValue)
		}
	}

	return
}

type ITestCase interface {
	ToString() string
	GetI() int
	GetJ() int
}

type IDataset interface {
	Load() error
	Display()
	Solve()
	PrintResult()
}

type TestCase struct {
	I int
	J int
}

func (c *TestCase) ToString() string {
	return fmt.Sprintf("%d %d", c.I, c.J)
}

func (c *TestCase) GetI() int {
	return c.I
}

func (c *TestCase) GetJ() int {
	return c.J
}

type Dataset struct {
	HighwayLength int
	CasesLength   int
	Width         []int
	Cases         []ITestCase
	Result        []int
}

func (d *Dataset) PrintResult() {
	for i := 0; i < len(d.Result); i++ {
		fmt.Println(d.Result[i])
	}
}

func (d *Dataset) Solve() {
	var (
		indexI, indexJ, max int
	)

	d.Result = make([]int, d.CasesLength, d.CasesLength)
	for i := 0; i < len(d.Result); i++ {
		max = 3
		indexI = d.Cases[i].GetI()
		indexJ = d.Cases[i].GetJ()
		for j := indexI; j <= indexJ; j++ {
			if d.Width[j] < max {
				max = d.Width[j]
			}
		}
		d.Result[i] = max
	}
}

func (d *Dataset) Load() (err error) {
	var (
		indexI int
		indexJ int
	)

	d.HighwayLength, err = readNextInt()
	if err != nil {
		return
	}

	d.CasesLength, err = readNextInt()
	if err != nil {
		return
	}

	d.Width = make([]int, d.HighwayLength, d.HighwayLength)
	for i := 0; i < len(d.Width); i++ {
		d.Width[i], err = readNextInt()
		if err != nil {
			return
		}
	}

	d.Cases = make([]ITestCase, d.CasesLength, d.CasesLength)
	for i := 0; i < len(d.Cases); i++ {
		indexI, err = readNextInt()
		if err != nil {
			return
		}
		indexJ, err = readNextInt()
		if err != nil {
			return
		}
		d.Cases[i] = NewTestCase(indexI, indexJ)
	}

	return
}

func (d *Dataset) Display() {
	fmt.Println("Dataset:")
	fmt.Println("  Highway length:", d.HighwayLength)
	fmt.Println("  Cases length:", d.CasesLength)
	fmt.Println("  Width:", d.Width)
	fmt.Println("  Cases:")
	for i := 0; i < len(d.Cases); i++ {
		fmt.Println("   -", d.Cases[i].ToString())
	}
	fmt.Println("  Highway:")
	for i := 0; i < d.HighwayLength; i++ {
		line := "    |"
		for j := 0; j < d.Width[i]; j++ {
			line = line + "-"
		}
		line = line + "|"

		fmt.Println(line)
	}
	if len(d.Result) > 0 {
		fmt.Println("  Result:")
		for i := 0; i < len(d.Result); i++ {
			fmt.Println("   -", d.Result[i])
		}
	}
}

func NewDataset() IDataset {
	return new(Dataset)
}

func NewTestCase(i, j int) ITestCase {
	return &TestCase{i, j}
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
