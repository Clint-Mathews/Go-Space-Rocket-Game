package main

import (
	gameframe "Clint-Mathews/Go-Space-Rocket-Game/game"
)

const BOUNDARY = 11

func main() {
	index := (BOUNDARY - 2)
	newGame := gameframe.NewGame(BOUNDARY, index)
	go newGame.StartRender()
	select {}
}
