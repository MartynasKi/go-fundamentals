package main

import (
	"fmt"
)

type Score struct{ Points int }

func (s Score) WithBonus(n int) Score { s.Points += n; return s }

func (s *Score) Add(n int) { s.Points += n }
func main() {
	score := Score{Points: 7}
	fmt.Println(score.WithBonus(3).Points, score.Points)
	score.Add(3)
	fmt.Println(score.Points)
}
