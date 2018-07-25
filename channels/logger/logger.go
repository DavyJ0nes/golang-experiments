package logger

import (
	"fmt"
	"io"
	"sync"
)

// Logger describes a concurrent logger object
type Logger struct {
	ch chan string
	wg sync.WaitGroup
}

// New instantiates a new logger and returns pointer to it
func New(w io.Writer, cap int) *Logger {
	l := Logger{
		ch: make(chan string, cap),
	}

	l.wg.Add(1)
	go func() {
		for v := range l.ch {
			fmt.Fprintf(w, v)
		}
		l.wg.Done()
	}()

	return &l
}

// Stop stops the logger by closing its channel
func (l *Logger) Stop() {
	close(l.ch)
	l.wg.Wait()
}

// Println pushes the string to channel, otherwise prints DROP
func (l *Logger) Println(s string) {
	select {
	case l.ch <- s:
	default:
		fmt.Println("DROP")
	}
}
