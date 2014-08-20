package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const PHONE_REGEX = "([0-9]{1,3}).([0-9]{1,3}).([0-9]{4,10})"

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

func ReadLines(limit int) (result []string, err error) {
	var scanner = bufio.NewScanner(os.Stdin)

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

type IPhoneNumber interface {
	Analyse()
	GetValue() string
	PrintInfos()
}

type PhoneNumber struct {
	value, countryCode, areaCode, number string
}

func (p *PhoneNumber) Analyse() {
	var (
		phoneRegExp *regexp.Regexp
		extracted   []string
	)

	phoneRegExp = regexp.MustCompile(PHONE_REGEX)
	extracted = phoneRegExp.FindStringSubmatch(p.value)

	p.countryCode = extracted[1]
	p.areaCode = extracted[2]
	p.number = extracted[3]
}

func (p *PhoneNumber) GetValue() string {
	return p.value
}

func (p *PhoneNumber) PrintInfos() {
	fmt.Printf("CountryCode=%s,LocalAreaCode=%s,Number=%s\n", p.countryCode, p.areaCode, p.number)
}

type IDataset interface {
	Display()
	Load() error
	Solve()
}

type Dataset struct {
	PhoneNumbersNumber int
	PhoneNumbers       []IPhoneNumber
}

func (d *Dataset) Display() {
	fmt.Println("Dataset:")
	fmt.Println("  # Phone numbers:", d.PhoneNumbersNumber)
	fmt.Println("  Phone numbers:")
	for i := 0; i < len(d.PhoneNumbers); i++ {
		fmt.Println("   ", d.PhoneNumbers[i].GetValue())
	}
}

func (d *Dataset) Load() (err error) {
	var phoneStrings []string

	if d.PhoneNumbersNumber, err = ReadNextInt(); err != nil {
		return
	}

	d.PhoneNumbers = make([]IPhoneNumber, d.PhoneNumbersNumber, d.PhoneNumbersNumber)

	if phoneStrings, err = ReadLines(d.PhoneNumbersNumber); err != nil {
		return
	}

	for i := 0; i < d.PhoneNumbersNumber; i++ {
		d.PhoneNumbers[i] = &PhoneNumber{value: phoneStrings[i]}
	}

	return
}

func (d *Dataset) Solve() {
	for i := 0; i < len(d.PhoneNumbers); i++ {
		d.PhoneNumbers[i].Analyse()
		d.PhoneNumbers[i].PrintInfos()
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
