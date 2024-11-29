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

func lexer(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Unable to read file in lexer %v\n", err)
	}
	var s scanner.Scanner
	s.Init(f)

	var sliceToken []string
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		// fmt.Printf("output %v\n", s.TokenText())
		sliceToken = append(sliceToken, s.TokenText())
	}

	return sliceToken
}

func parser(tokeSlice []string) int {

	if len(tokeSlice) == 0 {
		fmt.Println("Invalid Json")
		return 1
	}

	if tokeSlice[0] == "{" && tokeSlice[1] == "}" {
		fmt.Println("Valid Json")
		return 0
	}
	return -1
}

func jsonParser(fileName string) int {

	tokenSlice := lexer(fileName)

	exitCode := parser(tokenSlice)
	return exitCode
}

func main() {
	// const src = `
	// // This is scanned code.
	// if a > 10 {
	// 	someParsable = text
	// }`

	// validJson := "./test/step1/valid.json"
	// inValidJson := "./test/step1/invalid.json"

	filename := os.Args[1:]
	exitCode := jsonParser(filename[0])
	if exitCode != 0 {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
	// folderNumber := os.Args[1:]

	// files, err := os.ReadDir(dir + folderNumber[0])
	// if err != nil {
	// 	log.Fatalf("Unable to read files %v\n", err)
	// }
	// for _, file := range files {
	// 	// fmt.Printf("Printing the file name %s\n", file.Name())
	// 	// jsonParser(fmt.Sprintf("%s%s/%s", dir, folderNumber[0], file.Name()))
	// 	jsonParser(fmt.Sprintf("%s%s/%s", dir, folderNumber[0], file.Name()))
	// }

	// 	var s scanner.Scanner
	// 	s.Init(strings.NewReader(src))
	// 	s.Filename = "example"

	//	for tok := s.Next(); tok != scanner.EOF; tok = s.Scan() {
	//		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	//	}
}
