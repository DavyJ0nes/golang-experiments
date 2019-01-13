package main

import (
	"strings"

	"github.com/chrislusf/glow/flow"
)

func main() {
	f := newFileMapReducer("/etc/passwd").
		Filter(filter).
		Map(split).
		Map(count).
		Reduce(reduce).
		Map(output)

	f.Run()
}

func newFileMapReducer(fileName string) *flow.Dataset {
	return flow.New().TextFile(fileName, 3)
}

func filter(line string) bool {
	return !strings.HasPrefix(line, "#")
}

func split(line string, ch chan string) {
	for _, token := range strings.Split(line, ":") {
		ch <- token
	}
}

func count(key string) int {
	return 1
}

func reduce(x int, y int) int {
	return x + y
}

func output(x int) {
	println("count:", x)
}
