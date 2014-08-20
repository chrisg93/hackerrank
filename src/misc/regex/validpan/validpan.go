package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const PAN_REGEX = "[A-Z]{5}[0-9]{4}[A-Z]"

func ReadNextInt() (result int, err error) {
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
			result = int(parsedValue)
		}
	}

	return
}

func ReadNextLine() (result string, err error) {
	var readSize int

	readSize, err = fmt.Scan(&result)
	if readSize < 1 {
		err = errors.New("Can't read input")
	}

	return
}

type IPan interface {
	GetValue() string
	IsPan() bool
}

type Pan struct {
	value string
}

func (p *Pan) GetValue() string {
	return p.value
}

func (p *Pan) IsPan() bool {
	var (
		err     error
		matched bool
	)

	if matched, err = regexp.MatchString(PAN_REGEX, p.value); err != nil {
		panic(err.Error())
	}

	return matched
}

type IDataset interface {
	Display()
	Load() error
	Solve()
}

type Dataset struct {
	PansNumber int
	Pans       []IPan
}

func (d *Dataset) Display() {
	fmt.Println("Dataset:")
	fmt.Println("  Number of PANs:", d.PansNumber)
	fmt.Println("  PANs:")
	for i := 0; i < len(d.Pans); i++ {
		fmt.Println("   ", d.Pans[i].GetValue())
	}
}

func (d *Dataset) Load() (err error) {
	var panString string

	if d.PansNumber, err = ReadNextInt(); err != nil {
		return
	}

	d.Pans = make([]IPan, d.PansNumber, d.PansNumber)
	for i := 0; i < d.PansNumber; i++ {
		if panString, err = ReadNextLine(); err != nil {
			return
		}

		d.Pans[i] = &Pan{value: panString}
	}

	return
}

func (d *Dataset) Solve() {
	for i := 0; i < len(d.Pans); i++ {
		if d.Pans[i].IsPan() {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func main() {
	var (
		dataset IDataset = new(Dataset)
		err     error
	)

	if err = dataset.Load(); err != nil {
		fmt.Println(err.Error())
	} else {
		//dataset.Display()
		dataset.Solve()
	}
}
