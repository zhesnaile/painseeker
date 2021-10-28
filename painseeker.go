package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Struct and Methods
// {{{
type Writer interface {
	Write(p []byte) (n int, err error)
}

type ConsoleWriter struct{}
type FileWriter struct {
	file string
}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func (fw FileWriter) Write(data []byte) (int, error) {
	f, err := os.OpenFile(fw.file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	n, err := f.WriteString(string(data) + "\n")
	return n, err
}

// }}}

// assign a variable to a flag and it's shortened version.
func shortFlag(variable *string, name string, shortname string, def string, usage string) {
	flag.StringVar(variable, name, def, usage+".")
	flag.StringVar(variable, shortname, def, usage+" (short).")
}

//check for input/output file, and for help flag in case it shows up.
func handleFlags(inFile *string, outFile *string) (oFile bool) {
	// {{{
	oFile = true
	shortFlag(inFile, "input", "in", "", "Specifies the file to read from")
	shortFlag(outFile, "output", "out", "", "Specifies the file to write to")
	var h = flag.Bool("help", false, "Print this page.")

	flag.Parse()

	if *h == true {
		flag.PrintDefaults()
		os.Exit(0)
	}
	if *inFile == "" {
		flag.PrintDefaults()
		var err error = errors.New("You need a file to read from!")
		log.Println(err)
		os.Exit(1)
	}
	if *outFile == "" {
		oFile = false
	}
	return
	// }}}
}

func main() {
	var (
		inFile  string
		outFile string
		w       Writer
	)
	oF := handleFlags(&inFile, &outFile)
	if oF == true {
		w = FileWriter{
			file: outFile,
		}
	} else {
		w = ConsoleWriter{}
	}
	f, err := os.Open(inFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "//") {
			w.Write(scanner.Bytes())
		} else if strings.Contains(line, "/*") {
			w.Write(scanner.Bytes())
			b := !strings.Contains(line, "*/")
			for scanner.Scan() && b {
				line = scanner.Text()
				w.Write(scanner.Bytes())
				if strings.Contains(line, "*/") {
					b = false
				}
			}
		}
	}
}
