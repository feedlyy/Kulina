package main

import "fmt"

func number1(n int32, ar []int32) int32 {
	var pair, i int32

	var mapping = make(map[int32]int32)

	for i = 0; i < int32(len(ar)); i++ {
		mapping[ar[i]] += 1
	}

	for _, item := range mapping {
		if item < 2 {
			continue
		} else if item >= 2 {
			pair += item / 2
		}
	}

	return pair
}

func number2(n int32, s string) int32 {
	var valley, seaLevel, i int32

	seaLevel = 0

	for i = 0; i < int32(len(s)); i++ {
		if s[i:i+1] == "D" {
			seaLevel -= 1
		} else if s[i:i+1] == "U" {
			seaLevel += 1
		}

		if seaLevel == 0 && s[i:i+1] == "U" {
			valley += 1
		}
	}

	return valley
}

func main() {
	/*please comment or uncomment for testing
	each of the question*/

	//answer number 1
	var length, i int32
	_, _ = fmt.Scanln(&length)

	var arr = make([]int32, length)
	for i = 0; i < length; i++ {
		_, _ = fmt.Scanln(&arr[i])
	}
	fmt.Println(number1(length, arr))

}
