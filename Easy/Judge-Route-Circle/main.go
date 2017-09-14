package main

import "fmt"

type Robot struct {
	X int
	Y int
}

func (r Robot) isHere() bool {
	return r.X == 0 && r.Y == 0
}

func judgeCircle(moves string) bool {
	r := &Robot{
		X: 0,
		Y: 0,
	}

	for _, i := range moves {
		s := string(i)
		if s == "U" {
			r.X++
		} else if s == "D" {
			r.X--
		} else if s == "L" {
			r.Y--
		} else if s == "R" {
			r.Y++
		}
	}

	return r.isHere()
}

func main() {
	fmt.Print(judgeCircle("UDD"))
}
