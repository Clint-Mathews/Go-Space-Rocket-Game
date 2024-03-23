package gameframe

import (
	"Clint-Mathews/Go-Space-Rocket-Game/queue"
	"bufio"
	"bytes"
	"fmt"
	"math/rand/v2"
	"os"
	"time"
)

type Game struct {
	user           string
	score          int
	rocketPosition int
	bound          int
	drawBuf        *bytes.Buffer
	astroids       *queue.Queue
}

type GameOperations interface {
	StartRender()
	Render(closeChannel <-chan struct{})
	Update()
	Input(closeChannel chan struct{})
}

func NewGame(bound, rocketPosition int) GameOperations {
	fmt.Print("Enter username: ")
	reader := bufio.NewReader(os.Stdin)
	username, _, _ := reader.ReadLine()
	return &Game{
		user:           string(username),
		score:          0,
		rocketPosition: rocketPosition,
		bound:          bound,
		drawBuf:        new(bytes.Buffer),
		astroids:       queue.New(bound - 2),
	}
}

func (game *Game) Update() {
	game.drawBuf.WriteString("Space Rocket Game!\n")
	game.drawBuf.WriteString(fmt.Sprintf("User: %s - Score: %d \n", game.user, game.score))
	game.drawBuf.WriteString("Press a key (q to quit): \n")
	// getAstroids := game.astroids.GetData()
	// fmt.Println("%v", getAstroids[0])
	// fmt.Println("%v", getAstroids[1])
	// fmt.Println("%v", getAstroids)
	for i := 1; i <= game.bound; i++ {
		if i == 1 || i == game.bound {
			for j := 0; j < game.bound; j++ {
				game.drawBuf.WriteString("= ")
			}
			game.drawBuf.WriteString("\n")
		} else {
			game.drawBuf.WriteString("=")
			for j := 0; j < 2*game.bound-3; j++ {
				count := 0
				// if val := getAstroids[i-2]; val != 0 {
				// 	for val > 0 {
				// 		game.drawBuf.WriteString(" ")
				// 		count++
				// 		val--
				// 	}
				// 	game.drawBuf.WriteString("*")
				// }
				if i == game.bound-1 && j == game.rocketPosition {
					game.drawBuf.WriteString("^")
				} else {
					if count > 0 {
						count--
						continue
					}
					game.drawBuf.WriteString(" ")
				}
			}
			game.drawBuf.WriteString("=")
			game.drawBuf.WriteString("\n")
		}
	}
}

func (game *Game) Render(closeChannel <-chan struct{}) {
	for {
		select {
		default:
			game.drawBuf.Reset()
			fmt.Print("\033[H\033[2J")
			game.Update()
			fmt.Fprint(os.Stdout, game.drawBuf.String())
			game.updateAstroids()
			time.Sleep(500 * time.Millisecond)
		case <-closeChannel:
			os.Exit(0)
		}
	}
}

func (game *Game) StartRender() {
	exit := make(chan struct{})
	go game.Render(exit)
	go game.Input(exit)
}

func (game *Game) Input(closeChannel chan struct{}) {
	for {
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()

		if err != nil || char == 'q' {
			fmt.Println("Exiting game!")
			fmt.Println("Score:", game.score)
			close(closeChannel)
			return
		}

		if (char == 'a' || char == 'A') && game.rocketPosition > 1 {
			game.rocketPosition = game.rocketPosition - 2
		}

		if (char == 'D' || char == 'd') && game.rocketPosition < 2*game.bound-5 {
			game.rocketPosition = game.rocketPosition + 2
		}
	}
}

func (g *Game) updateAstroids() {
	if g.astroids.IsQueueFull() {
		g.astroids.Dequeue()
	}
	randomInt := rand.IntN(g.bound + 1)
	randomMultipleOfTwo := randomInt + (1 - randomInt%2)
	g.astroids.Enqueue(randomMultipleOfTwo)
}
