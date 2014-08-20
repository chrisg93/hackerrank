package main

import "fmt"

// James got hold of a love letter that his friend Harry has written for his girlfriend. Being the prankster that James is, he decides to meddle with it. He changes all the words in the letter into palindromes.
// While modifying the letters of the word, he follows 2 rules:
// (a) He always reduces the value of a letter, e.g. he changes 'd' to 'c', but he does not change 'c' to 'd'.
// (b) If he has to repeatedly reduce the value of a letter, he can do it until the letter becomes 'a'. Once a letter has been changed to 'a', it can no longer be changed.
// Each reduction in the value of any letter is counted as a single operation. Find the minimum number of operations he carries out to convert a given string into a palindrome.

// Input Format
// The first line contains an integer T, i.e., the number of test cases.
// The next T lines will contain a string each.

// Output Format
// A single line containing the number of minimum operations corresponding to each test case.

// Constraints
// 1 ≤ T ≤ 10
// 1 ≤ length of string ≤ 104
// All characters are lower cased english letters.

// Sample Input #00
// 3
// abc
// abcba
// abcd

// Sample Output #00
// 2
// 0
// 4

// Explanation
// For the first test case, ab*c* -> ab*b* -> ab*a*.
// For the second test case, abcba is a palindromic string.
// For the third test case, abc*d* -> abc*c* -> abc*b* -> abc*a* = ab*c*a -> ab*b*a.

func solve(source string) (operationCount uint) {
	var (
		strLength           int = len(source)
		iterator            int = 0
		firstChar, lastChar uint8
	)
	for ; iterator < strLength/2; iterator++ {
		firstChar = source[iterator]
		lastChar = source[strLength-iterator-1]

		for firstChar > lastChar {
			firstChar--
			operationCount++
		}
		for firstChar < lastChar {
			firstChar++
			operationCount++
		}
	}
	return
}

func main() {
	var (
		T          uint
		scanItems  int
		scannedStr string
		err        error
		results    []uint
	)

	results = make([]uint, 0, 0)

	if scanItems, err = fmt.Scanf("%d", &T); err != nil {
		fmt.Println(err.Error())
	} else if scanItems == 0 {
		fmt.Println("No data to scan")
	} else {
		//fmt.Println("T =", T)
	}

	for i := uint(0); i < T; i++ {
		if scanItems, err = fmt.Scanf("%s", &scannedStr); err != nil {
			fmt.Println(err.Error())
		} else if scanItems == 0 {
			fmt.Println("No data to scan")
		} else {
			//fmt.Printf("Test #%d = %s\n", i, scannedStr)
			results = append(results, solve(scannedStr))
		}
	}

	for i := 0; i < len(results); i++ {
		fmt.Println(results[i])
	}
}
