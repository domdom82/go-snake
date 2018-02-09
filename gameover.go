package main

import (
	tl "github.com/JoelOtter/termloop"
	"io/ioutil"
	"os"
)

type GameOver struct {
	*tl.Entity
}


func (t *GameOver) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		if event.Ch == 'q' {
			os.Exit(0)
		}
	}
}

func (t *GameOver) Draw(s *tl.Screen) {
	screenWidth,screenHeight := s.Size()
	myWidth,myHeight := t.Size()
	scoreWidth,scoreHeight := score.Size()

	t.SetPosition(screenWidth / 2 - myWidth / 2, screenHeight / 2 - myHeight / 2)
	score.SetPosition(screenWidth / 2 - scoreWidth / 2, screenHeight / 2 - scoreHeight / 2 + 10)
	t.Entity.Draw(s)
}


func NewGameOver() *GameOver {
	t := new(GameOver)

	dat, err := ioutil.ReadFile("gameover.txt")

	if err != nil {
		panic(err)
	}

	t.Entity = tl.NewEntityFromCanvas(1, 1, tl.CanvasFromString(string(dat)))

	return t
}
