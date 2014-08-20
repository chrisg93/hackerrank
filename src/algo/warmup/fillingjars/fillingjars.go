package main

import (
	"errors"
	"fmt"
	"strconv"
)

type IDataset interface {
	Load() (err error)
	Display()
	Resolve() (result int)
}

type Dataset struct {
	JarsNumber      int
	OperationNumber int
	Operations      [][]int
	Jars            []int
}

func (d *Dataset) Load() (err error) {
	var (
		readValue int
	)

	if readValue, err = readNextInt(); err != nil {
		return
	}
	d.JarsNumber = readValue

	if readValue, err = readNextInt(); err != nil {
		return
	}
	d.OperationNumber = readValue

	d.Operations = make([][]int, d.OperationNumber, d.OperationNumber)
	for i := 0; i < d.OperationNumber; i++ {
		d.Operations[i] = make([]int, 3, 3)

		if readValue, err = readNextInt(); err != nil {
			return
		}
		d.Operations[i][0] = readValue - 1

		if readValue, err = readNextInt(); err != nil {
			return
		}
		d.Operations[i][1] = readValue - 1

		if readValue, err = readNextInt(); err != nil {
			return
		}
		d.Operations[i][2] = readValue
	}

	return
}

func (d *Dataset) Display() {
	fmt.Println("Dataset:")
	fmt.Println("  Jars number: ", d.JarsNumber)
	fmt.Println("  Operation number: ", d.OperationNumber)
	fmt.Println("  Operations:")
	for i := 0; i < d.OperationNumber; i++ {
		fmt.Println("   ", d.Operations[i][0], ",", d.Operations[i][1], ",", d.Operations[i][2])
	}
	fmt.Println("  Jars:", d.Jars)
	fmt.Println("  Jars sum:", d.jarsSum())
}

func (d *Dataset) Resolve() (result int) {
	d.Jars = make([]int, d.JarsNumber, d.JarsNumber)
	d.fillJars()

	result = d.jarsSum() / d.JarsNumber

	return
}

func (d *Dataset) fillJars() {
	for i := 0; i < d.OperationNumber; i++ {
		for j := d.Operations[i][0]; j <= d.Operations[i][1]; j++ {
			d.Jars[j] = d.Jars[j] + d.Operations[i][2]
		}
	}
}

func (d *Dataset) jarsSum() (sum int) {
	for i := 0; i < d.JarsNumber; i++ {
		sum = sum + d.Jars[i]
	}
	return
}

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

func main() {
	var (
		dataset IDataset
		err     error
	)

	dataset = new(Dataset)
	if err = dataset.Load(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(dataset.Resolve())
		//dataset.Display()
	}
}
