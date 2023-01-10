package myhackerRank

import "fmt"

func BirthdayCakeCandles(candles []int32) int32 {
	var x int32
	var y int32

	for i := 0; i < len(candles); i++ {
		for j := i; j < len(candles); j++ {
			if candles[i] == candles[j] {
				fmt.Printf("A= %v ,B= %v \n", candles[i], candles[j])
				x++
			}
		}

		if y < x {
			y = x
		}
		x = 0

	}

	return y
}
