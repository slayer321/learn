package main

import (
	"fmt"
	"log"
	"os"
	"text/scanner"

	"github.com/slayer321/jsonparser/token"
)

var (
	dir    = "./tests/step"
	i      = 0
	result = false
)

func lexer(fileName string) []token.Token {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Unable to read file in lexer %v\n", err)
	}
	var s scanner.Scanner
	s.Init(f)

	var sliceToken []token.Token

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		// fmt.Printf("output %v\n", s.TokenText())
		//sliceToken = append(sliceToken, s.TokenText())

		switch s.TokenText() {
		case string(token.OPENCURLYBRACKET):
			sliceToken = append(sliceToken, token.Token{
				Type:  token.OPENCURLYBRACKET,
				Value: s.TokenText(),
			})
		case string(token.CLOSEDCURLYBRACKET):
			sliceToken = append(sliceToken, token.Token{
				Type:  token.CLOSEDCURLYBRACKET,
				Value: s.TokenText(),
			})
		case string(token.COLON):
			sliceToken = append(sliceToken, token.Token{
				Type:  token.COLON,
				Value: s.TokenText(),
			})

		case string(token.COMMA):
			sliceToken = append(sliceToken, token.Token{
				Type:  token.COMMA,
				Value: s.TokenText(),
			})
		default:
			sliceToken = append(sliceToken, token.Token{
				Type:  token.STRING,
				Value: s.TokenText(),
			})
		}

	}

	return sliceToken
}

// {"key": "value"}
func print(i int) {
	fmt.Printf("The Value of i is %v\n", i)
}

//	{
//		"key": "value",
//		"key2": "value"
//	 }

func isStringCorrect(value string) bool {

	stringLength := len(value)
	// first := fmt.Sprintf("%c", value[0])
	// last := fmt.Sprintf("%c", value[stringLength-1])

	first := value[0]
	last := value[stringLength-1]

	if first == '"' && last == '"' {
		return true
	}
	return false
}

func objectParser(tokens []token.Token, iterator token.TokenIterator) bool {

	// fmt.Println(tokens[i])
	// fmt.Println(tokens[i+1])
	iterator.Next()
	// print(i)

	if iterator.HasNext() {

		if tokens[i].Type == token.STRING {
			// fmt.Println(tokens[i].Value)
			value := tokens[i].Value

			correctString := isStringCorrect(value)
			if !correctString {
				return false
			}
			i++
			iterator.Next()

			if tokens[i].Type == token.COLON {
				i++
				iterator.Next()

				if tokens[i].Type == token.STRING {
					i++
					iterator.Next()

					if tokens[i].Type == token.CLOSEDCURLYBRACKET {
						return true
					} else if tokens[i].Type == token.COMMA {
						i++
						iterator.Next()

						return objectParser(tokens, iterator)
					}
				}
			}

		} else {
			return false
		}
	} else {
		return false
	}

	return false

}

func parser(tokenSlice []token.Token) (int, bool) {

	if len(tokenSlice) == 0 {
		fmt.Println("Invalid Json")
		return 1, false
	}

	// if tokenSlice[0] == "{" && tokenSlice[1] == "}" {
	//q 	fmt.Println("Valid Json")
	// 	return 0
	// }

	// {"key": "value"}

	// Iterator := token.NewToken(tokenSlice)
	for _, val := range tokenSlice {

		// var tokens strings.Builder
		// token := []string
		switch {
		case val.Type == token.OPENCURLYBRACKET:
			if tokenSlice[i+1].Type == token.CLOSEDCURLYBRACKET {
				return 0, true
			}

			iterator := token.NewTokenIterator(tokenSlice)
			i++
			result = objectParser(tokenSlice, iterator)
			return 0, result
		default:
			fmt.Println("Invalid")

		}

	}
	return -1, false
}

func jsonParser(fileName string) (int, string) {

	tokenSlice := lexer(fileName)
	// fmt.Printf("%+v\n", tokenSlice)

	exitCode, result := parser(tokenSlice)
	var out string
	out = "valid"
	if !result {
		out = "invalid"
	}
	return exitCode, out
}

func main() {
	// const src = `
	// // This is scanned code.
	// if a > 10 {
	// 	someParsable = text
	// }`

	// validJson := "./test/step1/valid.json"
	// inValidJson := "./test/step1/invalid.json"
	// filename := "./tests/step2/invalid2.json"
	// //filename := os.Args[1:]
	// var wg sync.WaitGroup

	// wg.Add(1)
	// go print(i)

	// _, out := jsonParser(filename)
	// // fmt.Println(out)
	// if out {
	// 	fmt.Println("Valid")
	// } else {
	// 	fmt.Println("Invalid")
	// }

	// wg.Done()
	// wg.Wait()

	// if exitCode != 0 {
	// 	os.Exit(1)
	// } else {
	// 	os.Exit(0)
	// }

	// ------------------------------------------------
	//folderNumber := os.Args[1:]

	files, err := os.ReadDir(dir + "2")
	if err != nil {
		log.Fatalf("Unable to read files %v\n", err)
	}
	for _, file := range files {
		// fmt.Printf("Printing the file name %s\n", file.Name())
		// jsonParser(fmt.Sprintf("%s%s/%s", dir, folderNumber[0], file.Name()))
		_, out := jsonParser(fmt.Sprintf("%s%s/%s", dir, "2", file.Name()))
		fmt.Printf("File name %v is %v\n", file.Name(), out)
		i = 0
	}

	// -----------------------------------------------------

	// 	var s scanner.Scanner
	// 	s.Init(strings.NewReader(src))
	// 	s.Filename = "example"

	//	for tok := s.Next(); tok != scanner.EOF; tok = s.Scan() {
	//		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	//	}
}
