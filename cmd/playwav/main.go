package main

import (
	"log"
	"os"

	"github.com/xyproto/playsample"
)

func main() {
	player := playsample.NewPlayer()
	defer player.Close()

	filename := "example.wav"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	if err := player.PlayWav(filename); err != nil {
		log.Fatalln(err)
	}
}
