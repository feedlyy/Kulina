package main

import (
	"fmt"
	"strconv"
)

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

func number3(n int) {

	//make a dynamic zero string
	init := strconv.Itoa(n)
	zero := ""
	for x := 0; x < len(init)-1; x++ { //len - 1 because the x start from 0 (incursion)
		zero += "0"
	}

	//make index for zero string
	index := 0

	for _, item := range init {
		if item == 48 { //skip when item is 0 / 48 byte ascii
			index += 1
			continue
		}
		fmt.Println(string(item) + zero[index:])
		index += 1
	}
}

func main() {
	/*please comment or uncomment for testing
	each of the question*/

	//answer number 1
	//var length, i int32
	//_, _ = fmt.Scanln(&length)
	//
	//var arr = make([]int32, length)
	//for i = 0; i < length; i++ {
	//	_, _ = fmt.Scanln(&arr[i])
	//}
	//fmt.Println(number1(length, arr))

	//answer number 2
	var n int32
	var s string

	_, _ = fmt.Scanln(&n, &s)
	fmt.Println(number2(n, s))
}
