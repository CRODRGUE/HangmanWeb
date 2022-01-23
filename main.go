package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"time"
)

// ==========================================================================================
//                         !!! Serveur Pour Le Jeu Hangman Web !!!
// ==========================================================================================
// By RODRIGUES Cyril

func main() {
	//Prends tout les fichier qui termine par ".html" pour lire leur defin (le nom du template !)
	templ, _ := template.ParseGlob("./templetes/*.html")

	//Page d'acc du jeux ! (index.html)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		templ.ExecuteTemplate(rw, "index", "")
	})

	//Page d'intialisation du jeux (Nom + difficulter) (game.html)
	http.HandleFunc("/game", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		player.player.id = ""
		nilRoad := r.URL.Path
		ruler1, _ := regexp.MatchString("/favicon.ico", nilRoad)
		if !ruler1 {
			queries := r.URL.Query()
			player.player.id = queries.Get("id_joueur")
			fmt.Println("==> ", player.player.id, " <===")
		}
		templ.ExecuteTemplate(rw, "game", "")
	})

	//Structure pour la page de jeux !
	type gameplay struct {
		Nom         string
		FindWord    []string
		Mode        string
		Vie         int
		ListeLetter []string
		Info        string
		Game        string
		End         string
		Photo       string
	}

	//Page de jeux ! (gameplay.html)
	http.HandleFunc("/gameClassique", func(rw http.ResponseWriter, r *http.Request) {
		nilRoad := r.URL.Path
		ruler1, _ := regexp.MatchString("/favicon.ico", nilRoad)
		queries := r.URL.Query()
		if !player.player.statu {
			player.file = "./words.txt"
			queries.Set("input_letter", "")
			hangmanMain()
		}
		if !ruler1 {
			player.letter = queries.Get("input_letter")
			fmt.Println("==> ", player.letter, " <===")
		}
		time.Sleep(150)
		DataGame1 := gameplay{player.player.id, player.word, "facile", player.counter, player.letterlist, player.player.info, player.player.game, "up", player.player.photo}
		fmt.Println("==>", player.word)
		fmt.Println("==>", player.counter)
		templ.ExecuteTemplate(rw, "gameplay", DataGame1)
	})

	http.HandleFunc("/gameExpert", func(rw http.ResponseWriter, r *http.Request) {
		nilRoad := r.URL.Path
		ruler1, _ := regexp.MatchString("/favicon.ico", nilRoad)
		queries := r.URL.Query()
		if !player.player.statu {
			player.file = "./words.txt"
			queries.Del("input_letter")
			hangmanMain()
			time.Sleep(200)
			player.counter -= 5
		}
		if !ruler1 {
			player.letter = queries.Get("input_letter")
			fmt.Println("==> ", player.letter, " <===")
		}
		time.Sleep(150)
		fmt.Println(player.counter)
		DataGame1 := gameplay{player.player.id, player.word, "expert", player.counter, player.letterlist, player.player.info, player.player.game, "up", player.player.photo}
		templ.ExecuteTemplate(rw, "gameplay", DataGame1)
	})

	http.HandleFunc("/gameMultiLang", func(rw http.ResponseWriter, r *http.Request) {
		nilRoad := r.URL.Path
		ruler1, _ := regexp.MatchString("/favicon.ico", nilRoad)
		queries := r.URL.Query()
		if !player.player.statu && !(0 <= player.counter && player.counter <= 10) {
			player.file = "./multlang.txt"
			queries.Del("input_letter")
			hangmanMain()
		}
		if !ruler1 {
			player.letter = queries.Get("input_letter")
			fmt.Println("==> ", player.letter, " <===")
		}
		time.Sleep(150)
		DataGame1 := gameplay{player.player.id, player.word, "multi-langues", player.counter, player.letterlist, player.player.info, player.player.game, "up", player.player.photo}
		templ.ExecuteTemplate(rw, "gameplay", DataGame1)
	})

	// Voici une autre route qui met à disposition des fichiers pour les autres fichiers du serveur,
	// !!! attention elle repond à la route localhost:555/static donc reste accessible !!!
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	// l'adresse du serveur en local
	http.ListenAndServe("localhost:666", nil)
}
