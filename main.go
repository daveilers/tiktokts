package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"strings"
	"time"
)

var voice = flag.String("v", "en_us_002", "Use a good voice like "+strings.Join(voices, " or "))

func main() {
	flag.Parse()
	text := "This is a test"
	if len(flag.Args()) > 0 {
		text = strings.Join(flag.Args(), " ")
	}
	sayText(text)
	time.Sleep(time.Second)
}

func sayText(s string) {
	data, err := tiktoktts(s, *voice)
	if err != nil {
		log.Printf("Ack %q %v: %v", s, *voice, err)
		return
	}
	playMp3(io.NopCloser(bytes.NewReader(data)))
}
