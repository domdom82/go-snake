package main

import (
	tl "github.com/JoelOtter/termloop"
)

// Direction of movement
type Direction int

// Direction values
const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

type Snake struct {
	*tl.Entity
	speed	  int
	body      []Point
	direction Direction
	grow	  int
}

// Head of snake
func (snake *Snake) head() *Point {
	return &snake.body[len(snake.body)-1]
}

// Tick for a snake
func (snake *Snake) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			if snake.direction != LEFT {
				snake.direction = RIGHT
			}
		case tl.KeyArrowLeft:
			if snake.direction != RIGHT {
				snake.direction = LEFT
			}
		case tl.KeyArrowUp:
			if snake.direction != DOWN {
				snake.direction = UP
			}
		case tl.KeyArrowDown:
			if snake.direction != UP {
				snake.direction = DOWN
			}
		case tl.KeySpace:
			//TESTING
			//food.Reset()
		}
	}
}

// Draw for a snake body
func (snake *Snake) Draw(screen *tl.Screen) {
	var head = *snake.head()
	w,h := snake.Size()

	switch snake.direction {
	case RIGHT:
		head.x += w * snake.speed
	case LEFT:
		head.x -= w * snake.speed
	case UP:
		head.y -= h * snake.speed
	case DOWN:
		head.y += h * snake.speed
	}

	// if snake runs off screen, re-enter at the other side
	screenWidth,screenHeight := game.Screen().Size()
	if head.x > screenWidth {
		head.x = 0
	}
	if head.x < 0 {
		head.x = screenWidth
	}
	if head.y > screenHeight {
		head.y = 0
	}
	if head.y < 0 {
		head.y = screenHeight
	}

	// handle snake-snake collision
	for b := 0; b < len(snake.body); b++ {
		if snake.body[b].x == head.x && snake.body[b].y == head.y {
			gameOver(snake)
		}
	}

	// handle snake growth and "movement"
	if snake.grow > 0 {
		snake.body = append(snake.body, head)
		snake.grow--
	} else {
		snake.body = append(snake.body[1:], head)
	}

	snake.SetPosition(head.x, head.y)

	for _, b := range snake.body {
		screen.RenderCell(b.x, b.y, &tl.Cell{
			Fg: tl.ColorRed,
			Ch: '\u2B1C',
		})
	}
}

func (snake *Snake) Position() (int, int) {
	return snake.Entity.Position()
}

func (snake *Snake) Size() (int, int) {
	return snake.Entity.Size()
}

func (snake *Snake) Collide(collision tl.Physical) {
	switch collision.(type) {
	case *Food:
		f := collision.(*Food)
		score.updateScore(f.score)
		snake.grow = 5
		f.Reset()
	}
}

func NewSnake() *Snake {
	s := new(Snake)
	s.Entity = tl.NewEntity(25, 10, 2, 1)
	s.body = []Point{
		{23, 10},
		{24, 10},
		{25, 10},
	}

	s.direction = RIGHT
	s.speed = 1
	return s
}
