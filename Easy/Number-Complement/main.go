package main

func findComplement(num int) int {
	origin := num
	counter := 1
	for num != 0 {
		num = num >> 1
		counter = counter<<1 + 1
	}
	return ^origin & (counter >> 1)
}
