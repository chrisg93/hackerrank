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

type IDataset interface {
	Display()
	GetCasesLength() int
	SetCasesLength(int)
	GetCase(int) ICase
	SetCase(int, ICase)
}

type ICase interface {
	SetCutCount(int)
	ToString() string
	Resolve() int
}

type Case struct {
	cutCount int
}

func (c *Case) Resolve() int {
	var (
		xCuts, yCuts int
	)

	for i := 0; i < c.cutCount; i++ {
		if xCuts > yCuts {
			yCuts = yCuts + 1
		} else {
			xCuts = xCuts + 1
		}
	}

	return xCuts * yCuts
}

func (c *Case) ToString() string {
	return strconv.FormatInt(int64(c.cutCount), 10)
}

func (c *Case) SetCutCount(newVal int) {
	c.cutCount = newVal
}

type Dataset struct {
	CasesLength int
	Cases       []ICase
}

func (d *Dataset) GetCase(index int) ICase {
	return d.Cases[index]
}

func (d *Dataset) SetCase(index int, c ICase) {
	d.Cases[index] = c
}

func (d *Dataset) SetCasesLength(newVal int) {
	d.CasesLength = newVal
	d.Cases = make([]ICase, d.CasesLength, d.CasesLength)
}

func (d *Dataset) GetCasesLength() int {
	return d.CasesLength
}

func (d *Dataset) Display() {
	fmt.Println("Dataset:")
	fmt.Println("  N° cases:", d.CasesLength)
	fmt.Println("  Cases:")
	for i := 0; i < len(d.Cases); i++ {
		fmt.Println("   ", d.Cases[i].ToString())
	}
}

func Load() (dataset IDataset, err error) {
	var (
		readedValue int
		c           ICase
	)

	dataset = new(Dataset)

	readedValue, err = readNextInt()
	if err != nil {
		return
	}

	dataset.SetCasesLength(readedValue)

	for i := 0; i < dataset.GetCasesLength(); i++ {
		readedValue, err = readNextInt()
		if err != nil {
			return
		}

		c = new(Case)
		c.SetCutCount(readedValue)
		dataset.SetCase(i, c)
	}

	return
}

func main() {
	var (
		dataset IDataset
		err     error
	)

	dataset, err = Load()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		//dataset.Display()
		for i := 0; i < dataset.GetCasesLength(); i++ {
			fmt.Println(dataset.GetCase(i).Resolve())
		}
	}
}
