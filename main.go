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
var link string

func main() {
	//Prends tous les fichier qui ce termine par ".html" pour lire leur defin (le nom du template !)
	templ, _ := template.ParseGlob("./templetes/*.html")

	//Page d'acc du jeux ! (index.html)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		templ.ExecuteTemplate(rw, "index", "")
	})

	//Page d'intialisation du jeux (Nom du joueur) (game.html)
	http.HandleFunc("/game", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		player.counter = -1
		player.letter = "err"
		player.player.id = ""
		templ.ExecuteTemplate(rw, "game", "")
	})

	//Page d'intialisation du jeux (Mode de jeu) (gamebis.html)
	http.HandleFunc("/gamebis", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		nilRoad := r.URL.Path
		ruler1, _ := regexp.MatchString("/favicon.ico", nilRoad)
		if !ruler1 {
			queries := r.URL.Query()
			player.player.id = queries.Get("id_joueur")
			fmt.Println("==> ", player.player.id, " <===")
		}
		templ.ExecuteTemplate(rw, "gamebis", "")
	})

	//Structure pour la page de jeux !
	type gameplay struct {
		Game        string
		Nom         string
		FindWord    []string
		Mode        string
		Vie         int
		ListeLetter []string
		Info        string
		Photo       string
		Link        string
	}

	//Pages de jeux ! (gameplay.html)
	http.HandleFunc("/gameClassique", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		nilRoad := r.URL.Path
		ruler1, _ := regexp.MatchString("/favicon.ico", nilRoad)
		queries := r.URL.Query()
		if !player.player.statu {
			link = string(r.URL.Path)
			player.file = "./words.txt"
			player.counter = 10
			queries.Del("input_letter")
			hangmanMain()
		}
		if !ruler1 {
			player.letter = queries.Get("input_letter")
			fmt.Println("==> ", player.letter, " <===")
		}
		time.Sleep(150)
		DataGame1 := gameplay{player.player.game, player.player.id, player.word, "facile", player.counter, player.letterlist, player.player.info, player.player.photo, link}
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
			player.counter = 5
			queries.Del("input_letter")
			hangmanMain()
			fmt.Println("-->", player.counter)
		}
		if !ruler1 {
			player.letter = queries.Get("input_letter")
			fmt.Println("==> ", player.letter, " <===")
		}
		time.Sleep(150)
		fmt.Println(player.counter)
		DataGame1 := gameplay{"up", player.player.id, player.word, "expert", player.counter, player.letterlist, player.player.info, player.player.photo, "google.fr"}
		templ.ExecuteTemplate(rw, "gameplay", DataGame1)
	})

	http.HandleFunc("/gameMultiLang", func(rw http.ResponseWriter, r *http.Request) {
		nilRoad := r.URL.Path
		ruler1, _ := regexp.MatchString("/favicon.ico", nilRoad)
		queries := r.URL.Query()
		if !player.player.statu {
			player.file = "./multlang.txt"
			player.counter = 15
			queries.Del("input_letter")
			hangmanMain()
		}
		if !ruler1 {
			player.letter = queries.Get("input_letter")
			fmt.Println("==> ", player.letter, " <===")
		}
		time.Sleep(150)
		DataGame1 := gameplay{"up", player.player.id, player.word, "multi-langues", player.counter, player.letterlist, player.player.info, player.player.photo, "google.fr"}
		templ.ExecuteTemplate(rw, "gameplay", DataGame1)
	})

	// Voici une autre route qui met à disposition des fichiers pour les autres fichiers du serveur (html; css...),
	// !!! attention elle repond à la route localhost:666/static donc reste accessible par tous !!!
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	// l'adresse du serveur en local
	http.ListenAndServe("localhost:666", nil)
}
