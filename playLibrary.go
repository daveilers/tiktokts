package main

import (
	"io"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var sampleRate = beep.SampleRate(48000)

func init() {
	err := speaker.Init(sampleRate, sampleRate.N(time.Second/2))
	if err != nil {
		log.Printf("Speaker problem: %v", err)
	}
}

func playMp3(rc io.ReadCloser) {
	streamer, format, err := mp3.Decode(rc)
	if err != nil {
		log.Printf("Error decoding: %v", err)
		return
	}
	defer streamer.Close()

	done := make(chan bool)
	speaker.Play(beep.Seq(beep.Resample(4, format.SampleRate, sampleRate, streamer), beep.Callback(func() {
		done <- true
	})))

	<-done
}
