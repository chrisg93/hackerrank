package main

import (
	"fmt"
)

// Given two integers: L and R,
// find the maximal values of A xor B given, L ≤ A ≤ B ≤ R

// Input Format
// The input contains two lines, L is present in the first line.
// R in the second line.

// Constraints
// 1 ≤ L ≤ R ≤ 10^3

// Output Format
// The maximal value as mentioned in the problem statement.

// Sample Input
// 1
// 10
// Sample Output
// 15

// Explanation
// The maximum value can be obtained for A = 5 and B = 10,
// 1010 xor 0101 = 1111 hence 15.

func ScanInt(dest *int) error {
	var err error
	_, err = fmt.Scanf("%d", dest)
	return err
}

func Solve(L, R int) int {
	var (
		max int
	)

	for i := L; i <= R; i++ {
		for j := i; j <= R; j++ {
			if i^j > max {
				max = i ^ j
			}
		}
	}
	return max
}

func main() {
	var (
		L, R, result int
		err          error
	)

	// Scan L & R

	if err = ScanInt(&L); err != nil {
		panic(err.Error())
	}
	if err = ScanInt(&R); err != nil {
		panic(err.Error())
	}
	// fmt.Println("L:", L)
	// fmt.Println("R:", R)

	// End of scan L & R

	result = Solve(L, R)
	fmt.Println(result)
}
