package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	c string
)

func fileByteReader(f *os.File) int {
	if f != os.Stdin {
		_, err := f.Seek(0, 0)
		if err != nil {
			log.Fatalf("Failed to seek the file %v\n", err)
		}
	}
	scanner := bufio.NewScanner(f)
	//scanner.Split(bufio.ScanWords)
	totalBytes := 0
	for scanner.Scan() {
		noOfBytes := scanner.Bytes()
		totalBytes += len(noOfBytes)

		// fmt.Printf("Text value is %s\n", scanner.Text())
		// log.Printf("Printing the number of bytes %v\n", len(noOfBytes))
	}

	return totalBytes
}

func numberOflines(f *os.File) int {
	if f != os.Stdin {
		_, err := f.Seek(0, 0)
		if err != nil {
			log.Fatalf("Failed to seek the file %v\n", err)
		}
	}
	scanner := bufio.NewScanner(f)

	linesCount := 0
	for scanner.Scan() {
		linesCount++
	}

	return linesCount
}

func numberOfwords(f *os.File) int {
	if f != os.Stdin {
		_, err := f.Seek(0, 0)
		if err != nil {
			log.Fatalf("Failed to seek the file %v\n", err)
		}
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	numberOfWords := 0
	for scanner.Scan() {
		numberOfWords++
	}

	return numberOfWords
}

func numberOfchar(f *os.File) int {
	if f != os.Stdin {
		_, err := f.Seek(0, 0)
		if err != nil {
			log.Fatalf("Failed to seek the file %v\n", err)
		}
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanBytes)

	numberOfchars := 0

	for scanner.Scan() {
		numberOfchars++
	}

	return numberOfchars

}

//var flagType *bool

func main() {
	bytesValue := flag.Bool("c", false, "number of bytes in the file")
	numberOfline := flag.Bool("l", false, "number of lines in the file")
	numberOfWords := flag.Bool("w", false, "number of words in the file")
	numberOfChar := flag.Bool("m", false, "number of bytes in the file")
	flag.Parse()

	var f *os.File
	var err error

	if filename := flag.Arg(0); filename != "" {
		f, err = os.Open(flag.Arg(0))
		if err != nil { // _, err := f.Seek(0, 0)
			// if err != nil {
			// 	log.Fatalf("Failed to seek the file %v\n", err)
			// }
			log.Fatalf("Unable to open the file %v\n", err)
		}
		defer f.Close()

	} else {

		f = os.Stdin
	}
	switch {
	case *bytesValue:
		totalBytes := fileByteReader(f)
		fmt.Printf("%v  %s\n", totalBytes, f.Name())
	case *numberOfline:
		linesCount := numberOflines(f)
		fmt.Printf("%d  %s\n", linesCount, f.Name())
	case *numberOfWords:
		numberOfWords := numberOfwords(f)
		fmt.Printf("%d  %s\n", numberOfWords, f.Name())
	case *numberOfChar:
		numberOfchars := numberOfchar(f)
		fmt.Printf("%d  %s\n", numberOfchars, f.Name())
	default:
		totalBytes := fileByteReader(f)
		linesCount := numberOflines(f)
		numberOfWords := numberOfwords(f)
		fmt.Printf("%d  %d  %d  %s\n", totalBytes, linesCount, numberOfWords, f.Name())

	}

}
