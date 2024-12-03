package main

import (
	"github.com/terfo1/postcreator/internal/app"
	"log"
)

func main() {
	application := app.New()
	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
