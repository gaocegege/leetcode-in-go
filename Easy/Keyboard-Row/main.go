package main

import (
	"fmt"
	"strings"
)

func findWords(words []string) []string {
	keyMap := map[string]int{
		"Q": 0,
		"W": 0,
		"E": 0,
		"R": 0,
		"T": 0,
		"Y": 0,
		"U": 0,
		"I": 0,
		"O": 0,
		"P": 0,
		"A": 1,
		"S": 1,
		"D": 1,
		"F": 1,
		"G": 1,
		"H": 1,
		"J": 1,
		"K": 1,
		"L": 1,
		"Z": 2,
		"X": 2,
		"C": 2,
		"V": 2,
		"B": 2,
		"N": 2,
		"M": 2,
	}

	res := make([]string, 0, len(words))

	for _, s := range words {
		buf := strings.ToUpper(s)
		row := keyMap[string(buf[0])]
		flag := true
		for _, r := range buf {
			c := string(r)
			if keyMap[c] != row {
				flag = false
			}
		}
		if flag == true {
			res = append(res, s)
		}
	}
	return res
}

func main() {
	fmt.Println(findWords([]string{"adsdf", "sfd"}))
}
