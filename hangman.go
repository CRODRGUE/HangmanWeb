package main

import (
	"fmt"
	"time"
)

type hangmanData struct {
	wordToFind string
	word       []string
	letterlist []string
	counter    int
	file       string
	letter     string
	player     playerData
}

type playerData struct {
	id    string
	statu bool
	game  string
	info  string
	photo string
}

var player hangmanData

func hangmanMain() {
	player.init()
	player.showRandLetter()

	fmt.Println("===================================================")
	fmt.Println("---> ", player.wordToFind)
	fmt.Println(player.word)
	fmt.Println(player.counter)
	fmt.Println("===================================================")

	player.displayJose()
	for player.checkWord() && player.counter > 0 {
		player.checkInput()
		player.displayJose()
		fmt.Println(player.word)
	}
	time.Sleep(200)
	player.player.statu = false
}
