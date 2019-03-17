package lexis

import (
	"bufio"
	"fmt"
	"log"

	"github.com/enabokov/language/bnf"
)

var bnfConfig bnf.BNF

func init() {
	bnfConfig = bnf.Read()
}

func Analyze(filename string) []Token {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	var lines []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text()+"\n")
	}

	stream := readInputStream(lines)
	tokens := readTokenStream(stream)

	for i := 0; i < 500; i++ {
		token := tokens.next()
		if token != nil {
			fmt.Println(*token)
		}
	}

	return nil
}
