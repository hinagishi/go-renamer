# go-renamer
go-renamer is a tool for renaming files

## Usage
go-renamer [optins] [filename or dirname]

## Options
    -t str
        Strip the str from the beginning and end of a filename
    -p str
        Append the str to the beginning of a filename
    -s str
        Append the str to the end of a filename

## DEMO

```
cd _demo
./gen.sh
```

Then, a directory is created with 10 files.

```
go-renamer dirname
```

The 10 files are renamed.
