package main

import tl "github.com/JoelOtter/termloop"
import "log"
import "bufio"
import "os"

const fps = 30

// Termloop stuff
var game *tl.Game

// BodyPart is a part of the snake
type BodyPart struct {
	*tl.Entity
	speed         float32
	x             float32
	y             float32
	direction     Direction
	nextDirection Direction
	nextThink     int
	head          bool
	next          *BodyPart
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

// Tick for a snake body
func (part *BodyPart) Tick(event tl.Event) {
	// only the head can change direction
	if part.head {
		if event.Type == tl.EventKey { // Is it a keyboard event?
			switch event.Key { // If so, switch on the pressed key.
			case tl.KeyArrowRight:
				part.direction = RIGHT
			case tl.KeyArrowLeft:
				part.direction = LEFT
			case tl.KeyArrowUp:
				part.direction = UP
			case tl.KeyArrowDown:
				part.direction = DOWN
			case tl.KeySpace:
				//TESTING
				grow(part)
			}
			if part.next != nil {
				part.next.nextDirection = part.direction
			}
			part.nextDirection = part.direction
		}
	}
}

// Draw for a snake body part
func (part *BodyPart) Draw(screen *tl.Screen) {
	part.nextThink--
	if part.nextThink <= 0 {
		part.nextThink = 5
		// apply my direction to child
		if part.next != nil {
			part.next.nextDirection = part.direction
		}
		// change my direction as planned
		part.direction = part.nextDirection
		part.speed = 5.0
	}

	frameSpeed := part.speed / fps

	switch part.direction {
	case RIGHT:
		part.x += frameSpeed
	case LEFT:
		part.x -= frameSpeed
	case UP:
		part.y -= frameSpeed
	case DOWN:
		part.y += frameSpeed
	}

	part.SetPosition(int(part.x), int(part.y))
	part.Entity.Draw(screen)

}

// NewBodyPart - Make a new body part
func NewBodyPart(x float32, y float32, speed float32, dir Direction) *BodyPart {
	p := BodyPart{tl.NewEntity(1, 1, 1, 1), speed, x, y, dir, dir, int(speed), false, nil}
	p.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '\u2B1C'})
	return &p
}

func grow(head *BodyPart) *BodyPart {
	var body *BodyPart
	if head == nil {
		// new snake
		body = NewBodyPart(40.0, 10.0, 5.0, DOWN)
		body.head = true
	} else {
		// add behind head
		body = NewBodyPart(head.x, head.y, 0, head.direction)
		if head.next != nil {
			body.next = head.next
		}
		head.next = body
	}
	game.Screen().AddEntity(body)
	return body
}

func printSnake(head *BodyPart) {
	log.Println(head)
	n := head.next
	for n != nil {
		log.Println(n)
		n = n.next
	}
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

	// Create snake head
	grow(nil)

	game.Start()
}
