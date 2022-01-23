package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func (p *hangmanData) input() string {
	var letter string
	for player.letter == "" {
		time.Sleep(5)
	}
	letter = player.letter
	player.letter = ""
	return letter
}

func (p *hangmanData) checkInput() {
	letter := p.input()
	letter = strings.ToLower(letter)
	for _, letterToRune := range letter {
		if !(('a' <= letterToRune && letterToRune <= 'z') || ('A' <= letterToRune && letterToRune <= 'Z')) {
			p.player.info = "Tu peux seulement mettre des carateres alphabetiques !! Recommence encore une fois l'ami..."
			p.input()
		}
	}
	if len(letter) != 1 {
		if letter == p.wordToFind {
			for index, char := range p.wordToFind {
				p.word[index] = string(char)
			}
		} else {
			p.counter -= 2
			p.player.info = "Tu t'es tromper l'ami -2 !! Tu penser avoir le bon mot..."
		}
	} else {
		check := p.duplicateEntry(letter)
		if check {
			p.player.info = "Tu as deja entree cette lettre l'ami ! Ps: Change de lettre =)"
		} else {
			p.letterlist = append(p.letterlist, letter)
			p.CheckLetter(strings.ToLower(letter))
		}
	}
}

func (p hangmanData) duplicateEntry(letter string) bool {
	var checkL bool
	for _, letterList := range p.letterlist {
		if letter == string(letterList) {
			checkL = true
		}
	}
	return checkL
}

func (p *hangmanData) displayJose() {
	p.player.photo = `./static/image/` + strconv.Itoa(p.counter) + `.png`
	fmt.Println("lien de la photo :", p.player.photo)
}
