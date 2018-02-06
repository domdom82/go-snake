package main

import tl "github.com/JoelOtter/termloop"
import "log"
import "bufio"
import (
	"os"
	"math/rand"
	"time"
	"fmt"
)

const fps = 10

// Termloop stuff
var game *tl.Game
var score *Score

func gameOver(snake *Snake) {
	//TODO do something nicer here
	fmt.Println("Final Score:", score.score)
	os.Exit(0)
}


func main() {
	// Set up RNG
	rand.Seed(time.Now().UTC().UnixNano())

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

	// Create food
	food := NewFood()

	// Create score
	score = NewScore()

	level.AddEntity(snake)
	level.AddEntity(food)
	game.Screen().AddEntity(score)

	game.Start()
}
