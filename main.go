package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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
	oldfiles := make([]string, len(files))
	newfiles := make([]string, len(files))
	for i := 0; i < len(files); i++ {
		oldfiles[i] = files[i].Name()
	}

	setName(oldfiles, newfiles, 0)

	for {
		showChangeList(oldfiles, newfiles)
		fmt.Print("Really change files name or modify?(y/n/m): ")
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
			break
		} else if c == 'm' {
			modifyName(oldfiles, newfiles)
		} else {
			break
		}
	}
}

func modifyName(oldfiles, newfiles []string) {
	fmt.Print("Input modify number -->")
	reader := bufio.NewReader(os.Stdin)
	c, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}
	num, err := strconv.Atoi(string(c))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Input new file name -->")
	c, _, err = reader.ReadLine()
	newfiles[num] = string(c)
}

func setName(oldfiles, newfiles []string, start int) {
	index := 0
	for i := start; i < len(oldfiles); i++ {
		if oldfiles[i][0] == '.' {
			newfiles[i] = oldfiles[i]
		} else if strings.Index(oldfiles[i], ".") == -1 {
			newfiles[i] = fmt.Sprintf("%03d", index)
			index++
		} else {
			tmp := strings.Split(oldfiles[i], ".")
			suffix := tmp[len(tmp)-1]
			newfiles[i] = fmt.Sprintf("%03d.%s", index, suffix)
			index++
		}
	}
}

func showChangeList(oldfiles, newfiles []string) {
	for i := 0; i < len(newfiles); i++ {
		fmt.Println(strconv.Itoa(i) + ": " + oldfiles[i] + " --> " + newfiles[i])
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
