package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := []string{"."}

	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	for _, arg := range args {
		err := tree(arg)
		if err != nil {
			log.Printf("tree %s: %v\n", arg, err)
		}
	}
}

func tree(path string) error {
	fmt.Println(path)
	return nil
}
