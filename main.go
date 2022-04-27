package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var voice = flag.String("v", "en_us_002", "Use a good voice like "+strings.Join(voices, " or "))
var output = flag.String("o", "", "Output to an mp3 file")
var quiet = flag.Bool("q", false, "Don't play sound audibly")

func main() {
	flag.Parse()
	text := "This is a test"
	if len(flag.Args()) > 0 {
		text = strings.Join(flag.Args(), " ")
	}
	mp3Data, err := tiktoktts(text, *voice)
	if err != nil {
		log.Printf("Ack %q %v: %v", text, *voice, err)
	}
	if !*quiet {
		playMp3(io.NopCloser(bytes.NewReader(mp3Data)))
		time.Sleep(time.Second)
	}
	if *output != "" {
		err = os.WriteFile(*output, mp3Data, 0700)
		if err != nil {
			log.Printf("Problem writing file %q: %v", *output, err)
		}
	}
}

func sayText(s string) {
	mp3Data, err := tiktoktts(s, *voice)
	if err != nil {
		log.Printf("Ack %q %v: %v", s, *voice, err)
		return
	}
	playMp3(io.NopCloser(bytes.NewReader(mp3Data)))
}
