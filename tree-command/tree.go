package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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

func tree(root string) error {
	err := filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		// if err != nil {
		// return err
		// }

		println(fi.Name()[0], "dot:", '.')
		if fi.Name()[0] == '.' {
			return filepath.SkipDir
		}

		fmt.Println(fi.Name())
		return nil

	})
	return err
}
