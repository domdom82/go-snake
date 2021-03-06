package main

import (
	tl "github.com/JoelOtter/termloop"
	"math/rand"
)

type Food struct {
	*tl.Entity
	body Point
	score int
}

func (food *Food) Draw(screen *tl.Screen) {
	screen.RenderCell(food.body.x, food.body.y, &tl.Cell{
		Fg: tl.ColorMagenta,
		Ch: '*',
	})
}

func (food *Food) Position() (int, int) {
	return food.Entity.Position()
}

func (food *Food) Size() (int, int) {
	return food.Entity.Size()
}

func (food *Food) Reset(snake *Snake) {
	screenWidth, screenHeight := game.Screen().Size()

	x := rand.Intn(screenWidth - 1) + 1
	y := rand.Intn(screenHeight - 1) + 1
	food.SetPosition(x,y)
	food.body.x = x
	food.body.y = y

	//re-roll if we are inside snake
	for _,b := range snake.body {
		if b.x == x && b.y == y {
			food.Reset(snake)
		}
	}
}

func NewFood() *Food {
	f := new(Food)
	f.Entity = tl.NewEntity(60, 20, 1, 1)
	f.body = Point{60, 20	}

	f.score = 10
	return f
}


