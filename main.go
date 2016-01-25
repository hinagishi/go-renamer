package main

import (
	"bufio"
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Usage: go-renamer filepath")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	f, err := os.Stat(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if f.IsDir() {
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print(f.Name() + "--> ")
		for scanner.Scan() {
			nf := scanner.Text()
			if nf == "" {
				fmt.Print(f.Name() + "--> ")
				continue
			}
			os.Rename(os.Args[1], nf)
			break
		}
	}
}
