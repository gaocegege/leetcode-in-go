package main

func distributeCandies(candies []int) int {
	candyMap := make(map[int]bool)
	counter := 0
	for _, candy := range candies {
		if !candyMap[candy] {
			counter++
			candyMap[candy] = true
		}
	}
	if counter > len(candies)/2 {
		counter = len(candies) / 2
	}
	return counter
}
