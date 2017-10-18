package main

import (
	"fmt"
)

func removeInvalidParentheses(s string) []string {
	res := &[]string{}
	remove(s, res, 0, 0, "(", ")")
	fmt.Println(res)
	return *res
}

func remove(s string, result *[]string, pos int, start int, first, second string) {
	stack := 0
	for i := start; i < len(s); i++ {
		if string(s[i]) == first {
			stack++
		} else if string(s[i]) == second {
			stack--
			if stack < 0 {
				fmt.Println(i)
				bugPos := i
				for removeIndex := pos; removeIndex <= bugPos; removeIndex++ {
					if string(s[removeIndex]) == second && (removeIndex == pos || string(s[removeIndex-1]) != second) {
						fmt.Println(s[:removeIndex] + s[removeIndex+1:])
						fmt.Println(i)
						remove(s[:removeIndex]+s[removeIndex+1:], result, removeIndex, i, first, second)
					}
				}
				return
			}
		}
	}
	if first == "(" {
		remove(reverse(s), result, 0, 0, second, first)
	} else {
		*result = append(*result, reverse(s))
	}
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(removeInvalidParentheses("()())()"))
}
