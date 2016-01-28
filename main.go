package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type FileName struct {
	Oldname string
	Newname string
	Modify  bool
}

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
	filename := make([]FileName, 0)
	for i := 0; i < len(files); i++ {
		if !files[i].IsDir() {
			filename = append(filename, FileName{Oldname: files[i].Name(), Newname: "", Modify: false})
		}
	}

	setName(filename, 0)

	for {
		showChangeList(filename)
		fmt.Print("Really change files name or modify?(y/n/m): ")
		reader := bufio.NewReader(os.Stdin)
		c, err := reader.ReadByte()
		if err != nil {
			fmt.Println(err)
			return
		}
		if c == 'y' {
			for i := 0; i < len(filename); i++ {
				os.Rename(f+filename[i].Oldname, f+filename[i].Newname)
			}
			break
		} else if c == 'm' {
			modifyName(filename)
			setName(filename, 0)
		} else {
			break
		}
	}
}

func modifyName(filename []FileName) {
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
	if checkName(filename, string(c)) {
		fmt.Println("Already exists the same name file")
		return
	}
	filename[num].Newname = string(c)
	filename[num].Modify = true
}

func checkName(filename []FileName, f string) bool {
	for _, fn := range filename {
		if fn.Newname == f {
			return true
		}
	}
	return false
}

func setName(filename []FileName, start int) {
	index := 0
	for i := start; i < len(filename); i++ {
		if filename[i].Modify {
			continue
		}
		if filename[i].Oldname[0] == '.' {
			filename[i].Newname = filename[i].Oldname
		} else if strings.Index(filename[i].Oldname, ".") == -1 {
			tmp := fmt.Sprintf("%03d", index)
			filename[i].Newname = tmp
			index++
		} else {
			tmp := strings.Split(filename[i].Oldname, ".")
			suffix := tmp[len(tmp)-1]
			fname := fmt.Sprintf("%03d.%s", index, suffix)
			filename[i].Newname = fname
			index++
		}
	}
}

func showChangeList(filename []FileName) {
	for i := 0; i < len(filename); i++ {
		fmt.Println(strconv.Itoa(i) + ": " + filename[i].Oldname + " --> " + filename[i].Newname)
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
