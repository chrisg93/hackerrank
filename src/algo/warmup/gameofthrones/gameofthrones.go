package main

import "fmt"

func ValidCount(src map[uint8]int) bool {
	var (
		alreadyHasOdd bool = false
		value         int
	)

	for _, value = range src {
		if value%2 == 0 {
			continue
		} else {
			if alreadyHasOdd {
				return false
			} else {
				alreadyHasOdd = true
			}
		}
	}

	return true
}

func CanBePalindromic(src string) bool {
	var (
		charCountMap map[uint8]int = make(map[uint8]int, 0)
		curCount     int
		found        bool
		char         uint8
	)

	for i := 0; i < len(src); i++ {
		char = src[i]
		if curCount, found = charCountMap[char]; found {
			charCountMap[char] = curCount + 1
		} else {
			charCountMap[char] = 1
		}
	}

	if ValidCount(charCountMap) {
		return true
	} else {
		return false
	}
}

func ReadStr() (string, error) {
	var readErr error
	var readVal string

	_, readErr = fmt.Scanln(&readVal)
	return readVal, readErr
}

func main() {
	var (
		sourceStr string
		err       error
	)

	if sourceStr, err = ReadStr(); err != nil {
		panic(err.Error())
	}

	if CanBePalindromic(sourceStr) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
