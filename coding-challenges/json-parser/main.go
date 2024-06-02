package main

import (
	"fmt"
	"log"
	"os"
	"text/scanner"
)

var (
	dir = "./tests/step"
)

func lexer(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Unable to read file in lexer %v\n", err)
	}
	var s scanner.Scanner
	s.Init(f)

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("output %v\n", s.TokenText())
	}
}

func Parser(fileName string) {

	lexer(fileName)

}

func main() {
	// const src = `
	// // This is scanned code.
	// if a > 10 {
	// 	someParsable = text
	// }`

	// validJson := "./test/step1/valid.json"
	// inValidJson := "./test/step1/invalid.json"

	folderNumber := os.Args[1:]

	files, err := os.ReadDir(dir + folderNumber[0])
	if err != nil {
		log.Fatalf("Unable to read files %v\n", err)
	}

	for _, file := range files {
		fmt.Printf("Printing the file name %s\n", file.Name())
		Parser(fmt.Sprintf("%s%s/%s", dir, folderNumber[0], file.Name()))
	}

	// 	var s scanner.Scanner
	// 	s.Init(strings.NewReader(src))
	// 	s.Filename = "example"

	//	for tok := s.Next(); tok != scanner.EOF; tok = s.Scan() {
	//		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	//	}
}
