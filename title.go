package main

import (
	tl "github.com/JoelOtter/termloop"
	"io/ioutil"
)

type Title struct {
	*tl.Entity
}


func (t *Title) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		gameScreen()
	}
}

func (t *Title) Draw(s *tl.Screen) {
	screenWidth,screenHeight := s.Size()
	myWidth,myHeight := t.Size()

	t.SetPosition(screenWidth / 2 - myWidth / 2, screenHeight / 2 - myHeight / 2)
	t.Entity.Draw(s)
}


func NewTitle() *Title {
	t := new(Title)

	dat, err := ioutil.ReadFile("title.txt")

	if err != nil {
		panic(err)
	}

	t.Entity = tl.NewEntityFromCanvas(1, 1, tl.CanvasFromString(string(dat)))

	return t
}
