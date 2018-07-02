package main

import (
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage")
		fmt.Println("-----")
		fmt.Println("$ timer 5m")
		os.Exit(1)
	}

	durationArg := os.Args[len(os.Args)-1]

	duration, err := time.ParseDuration(durationArg)
	if err != nil {
		panic(err)
	}
	timer(duration)
}

func timer(d time.Duration) {
	fmt.Printf("Starting Timer for %s\n", d.String())
	now := time.Now()
	end := now.Add(d)

	// Having to do this for formatting Issues
	hours := big.NewInt(int64(time.Until(end).Hours()))
	mins := big.NewInt(int64(time.Until(end).Minutes()))
	secs := big.NewInt(int64(time.Until(end).Seconds()))
	mins = mins.Mod(mins, big.NewInt(60))
	secs = secs.Mod(secs, big.NewInt(60))
	fmt.Printf("\rTimer:\t %02d:%02d:%02d", hours, mins, secs)

	ticker := time.NewTicker(1 * time.Second)
	var i int64
	for range ticker.C {
		i++

		hours = big.NewInt(int64(time.Until(end).Hours()))
		mins = big.NewInt(int64(time.Until(end).Minutes()))
		secs = big.NewInt(int64(time.Until(end).Seconds()))
		mins = mins.Mod(mins, big.NewInt(60))
		secs = secs.Mod(secs, big.NewInt(60))

		fmt.Printf("\rTimer:\t %02d:%02d:%02d", hours, mins, secs)

		if i > int64(d/time.Second) {
			ticker.Stop()
			alert()
			break
		}

	}
}

func alert() {
	// Bit of a hack to clear the screen
	fmt.Printf("\r                    \n")
	done := make(chan struct{})

	home, _ := os.LookupEnv("HOME")
	f, err := os.Open(fmt.Sprintf("%s/.timer/alert.mp3", home))
	if err != nil {
		panic(err)
	}

	s, format, _ := mp3.Decode(f)

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		close(done)
	})))

	<-done
	fmt.Printf("\nDONE!\n")
}
