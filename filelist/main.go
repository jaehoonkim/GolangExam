package main

import (
	//"fmt"
	"os"
	"path/filepath"
	"flag"
	"bufio"
)

var listWriter *bufio.Writer

func WalkFunc(path string, info os.FileInfo, err error) error {
	var writer_err error

	if info.IsDir() == false {
		//fmt.Println(path)
		
		listWriter.WriteString(path)
		listWriter.WriteString("\r\n")
		writer_err = listWriter.Flush()
	}

	return writer_err
}

func main() {
	var path = flag.String("path", "./", "directory path")
	flag.Parse()
	resultFile := "./result.txt"
	f, err := os.Create(resultFile)
	if err != nil {
		panic(err)
	}
	listWriter = bufio.NewWriter(f)

	filepath.Walk(*path, WalkFunc)
}

