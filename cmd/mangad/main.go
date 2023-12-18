package main

import (
	"log"

	flowermanga "github.com/breno5g/mangad/internal/flower-manga"
)

type chapter struct {
	number float64
	Url    string
}

type manga struct {
	title    string
	chapters []chapter
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flowermanga.GetMangaChapters("https://flowermanga.com/manga/jujutsu-kaisen/")
}
