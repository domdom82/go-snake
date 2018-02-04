package main

import tl "github.com/JoelOtter/termloop"
import "log"
import "bufio"
import (
	"os"
)

const fps = 10

// Termloop stuff
var game *tl.Game

type Point struct {
	x	int
	y 	int
}

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
	grow	  bool
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
			snake.grow = true
		}
	}
}

// Draw for a snake body
func (snake *Snake) Draw(screen *tl.Screen) {
	var head = *snake.head()

	switch snake.direction {
	case RIGHT:
		head.x += snake.speed
	case LEFT:
		head.x -= snake.speed
	case UP:
		head.y -= snake.speed
	case DOWN:
		head.y += snake.speed
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
	if snake.grow {
		snake.body = append(snake.body, head)
		snake.grow = false
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

func NewSnake() *Snake {
	s := new(Snake)
	s.Entity = tl.NewEntity(25, 10, 1, 1)
	s.body = []Point{
		{23, 10},
		{24, 10},
		{25, 10},
	}

	s.direction = RIGHT
	s.speed = 1
	return s
}


func gameOver(snake *Snake) {
	os.Exit(0)
}

func main() {
	// Set up logging
	logfile := "go-snake.log"
	file, err := os.Create(logfile)
	if err != nil {
		log.Fatal("Could not open log file ", logfile)
	}
	writer := bufio.NewWriter(file)
	log.SetOutput(writer)
	defer file.Close()
	defer writer.Flush()

	// Game setup
	game = tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorBlack,
	})
	game.Screen().SetLevel(level)
	game.Screen().SetFps(fps)

	// Create snake
	snake := NewSnake()

	game.Screen().AddEntity(snake)

	game.Start()
}
