package main

import (
	tl "github.com/JoelOtter/termloop"
	"fmt"
)

type Score struct {
	*tl.Text
	score	int
}

func (s *Score) updateScore(x int) {
	s.score += x
	s.Text.SetText(fmt.Sprint("Score: ", s.score))
}

func NewScore() *Score {
	s := new(Score)
	s.Text = tl.NewText(0, 0, "Score:  0", tl.ColorWhite, tl.ColorBlue)
	s.score = 0

	return s
}