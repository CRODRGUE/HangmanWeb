package main

func (p *hangmanData) CheckLetter(letter string) {
	var checkF bool
	for indexFind, letterFind := range p.wordToFind {
		if letter == string(letterFind) {
			checkF = true
			p.word[indexFind] = letter
		}
	}
	if !checkF {
		p.player.info = "La lettre n'est pas presente dans le mot -1 essais"
		p.counter--
	} else {
		p.player.info = "La lettre est presente dans le mot !"
	}
}

func (p *hangmanData) checkWord() bool {
	var check bool
	for _, letter := range p.word {
		if letter == "_" {
			check = true
		}
	}
	if check && p.counter <= 0 {
		p.player.info = "Dommage tu as depasser le nombre maximum d'essais... =( \n Le mot à trouver était : " + p.wordToFind
		check = false
		p.player.game = ""
	} else if !check && p.counter != 0 {
		p.player.info = "Bravo l'aim ! tu as trouver le mot sercet... \n Qui était : " + p.wordToFind
		p.player.game = ""
	}
	return check
}
