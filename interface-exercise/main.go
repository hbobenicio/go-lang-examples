package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("couldn't read file %s: %v", path, err)
	}

	_, err = io.Copy(os.Stdin, file)
	if err != nil {
		return fmt.Errorf("couldn't copy content from %s to stdin: %v", path, err)
	}

	return nil
}

func main() {
	if len(os.Args) == 1 {
		io.Copy(os.Stdin, os.Stdin)
	} else {
		failed := false
		filePaths := os.Args[1:]

		for _, filePath := range filePaths {
			if err := readFile(filePath); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				failed = true
			}
		}

		if failed {
			os.Exit(1)
		}
	}
}
