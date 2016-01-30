/*
go-renamer is a tool to rename file.
It renames a file to a specified file name if the target is a file.
It also renames all files in a specified directory.
	Author: hinagishi
*/
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

/*
FileName contains an old file name, a new file
and a modified flag.
*/
type FileName struct {
	Oldname string
	Newname string
	Modify  bool
}

/*
Options contains commandline-argments
*/
type Options struct {
	Trim   string
	Suffix string
	Append string
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
	var filename []FileName
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
		fmt.Printf("\x1b[31m%s\x1b[0m\n", "Already exists the same name file")
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
	fmt.Print("\n")
	for i := 0; i < len(filename); i++ {
		fmt.Println(strconv.Itoa(i) + ": " + filename[i].Oldname + " --> " + filename[i].Newname)
	}
	fmt.Print("\n")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	var opt Options
	var target string
	fmt.Println(os.Args)
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-t" {
			if i+1 >= len(os.Args) {
				usage()
				return
			} else {
				opt.Trim = os.Args[i+1]
				i++
			}
		} else {
			target = os.Args[i]
		}
	}
	if target == "" {
		usage()
		return
	}
	f, err := os.Stat(target)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if f.IsDir() {
		renameAll(target)
	} else {
		path := filepath.Dir(target)
		if opt.Trim != "" {
			base := strings.Trim(filepath.Base(target), opt.Trim)
			os.Rename(target, path+"/"+base)
		} else {
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print(f.Name() + "--> ")
			for scanner.Scan() {
				nf := scanner.Text()
				if nf == "" {
					fmt.Print(f.Name() + "--> ")
					continue
				}
				path := filepath.Dir(target)
				os.Rename(target, path+"/"+nf)
				break
			}
		}
	}
}
