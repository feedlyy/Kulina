package main

// Complete the sockMerchant function below.
func number1(n int32, ar []int32) int32 {
	var pair,i int32

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



func main() {
	//please note
}
