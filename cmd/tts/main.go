package main

import (
	"flag"
	"github.com/nasermirzaei89/tts"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	out := flag.String("o", "tts.mp3", "Output mp3 file path")
	flag.Parse()

	text := strings.Join(flag.Args(), " ")
	log.Print(text)
	voice, err := tts.Say(text)
	if err != nil {
		log.Panic(err)
	}
	err = ioutil.WriteFile(*out, voice, 0644)

	if err != nil {
		log.Panic(err)
	}
}
