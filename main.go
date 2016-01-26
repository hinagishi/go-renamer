package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func usage() {
	fmt.Println("Usage: go-renamer filepath")
}

func renameAll(f string) {
	if !strings.HasSuffix(f, "/") {
		f += "/"
	}
	files, err := ioutil.ReadDir(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	i := 0
	newfiles := make([]string, len(files))
	for p, n := range files {
		if n.Name()[0] == '.' {
			newfiles[p] = n.Name()
		} else if strings.Index(n.Name(), ".") == -1 {
			newfiles[p] = fmt.Sprintf("%03d", i)
			i++
		} else {
			tmp := strings.Split(n.Name(), ".")
			suffix := tmp[len(tmp)-1]
			newfiles[p] = fmt.Sprintf("%03d.%s", i, suffix)
			i++
		}
	}

	for i := 0; i < len(newfiles); i++ {
		fmt.Println(files[i].Name() + " --> " + newfiles[i])
	}
	fmt.Print("Really change files name?(y/n): ")
	reader := bufio.NewReader(os.Stdin)
	c, err := reader.ReadByte()
	if err != nil {
		fmt.Println(err)
		return
	}
	if c == 'y' {
		for i := 0; i < len(newfiles); i++ {
			os.Rename(f+files[i].Name(), f+newfiles[i])
		}
	}
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
		renameAll(os.Args[1])
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
