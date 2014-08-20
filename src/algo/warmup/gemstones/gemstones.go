package main

import (
	"fmt"
	"strings"
)

// John has discovered various rocks. Each rock is composed of various elements, and each element is represented by a lowercase latin letter from 'a' to 'z'.
// An element can be present multiple times in a rock. An element is called a 'gem-element' if it occurs at least once in each of the rocks.
// Given the list of rocks with their compositions, display the number of gem-elements that exist in those rocks.

// Input Format
// The first line consists of N, the number of rocks.
// Each of the next N lines contain rocks' composition. Each composition consists of lowercase letters of English alphabet.

// Output Format
// Print the number gem-elements that exist in those rocks.

// Constraints
// 1 ≤ N ≤ 100
// Each composition consists of only small latin letters ('a'-'z').
// 1 ≤ Length of each composition ≤ 100

// Sample Input
// 3
// abcdde
// baccd
// eeabg

// Sample Output
// 2

// Explanation
// Only "a", "b" are the two kind of gem-elements, since these are the only characters that occur in each of the rocks' composition.

func ScanGem(readValue *string) error {
	_, err := fmt.Scanf("%s", readValue)
	return err
}

func ScanGemCount(readValue *int) error {
	_, err := fmt.Scanf("%d", readValue)
	return err
}

func main() {
	var (
		GemElements    string = "abcdefghijklmnopqrstuvwxyz"
		GemCount       int
		GemScanned     string
		Gems           []string = make([]string, 0, 0)
		err            error
		CurrentElement uint8
	)

	// Read values from stdin

	if err = ScanGemCount(&GemCount); err != nil {
		panic(err.Error())
	}
	//fmt.Println("Gem count:", GemCount)

	for i := 0; i < GemCount; i++ {
		if err = ScanGem(&GemScanned); err != nil {
			panic(err.Error())
		}
		Gems = append(Gems, GemScanned)
	}
	//fmt.Println("Gems:", Gems)

	// End of read values from stdin

	// For each gem, we parse all GemElements and we remove these if there not present in current gem
	for i := 0; i < len(Gems); i++ {
		var gem string = Gems[i]
		//fmt.Println("Analyzing gem:", gem)

		for j := len(GemElements) - 1; j >= 0; j-- {
			CurrentElement = GemElements[j]

			//fmt.Println("Checking gem element:", string(CurrentElement))
			if strings.Contains(gem, string(CurrentElement)) == false {
				//fmt.Println("Removing...")
				GemElements = strings.Replace(GemElements, string(CurrentElement), "", 1)
			}
		}
	}

	//fmt.Println("Common elements:", GemElements)
	fmt.Println(len(GemElements))
}
