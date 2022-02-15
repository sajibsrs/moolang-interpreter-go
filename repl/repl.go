package repl

import (
	"bufio"
	"fmt"
	"io"
	"moolang/lexer"
	"moolang/token"
)

const PROMT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		lex := lexer.New(line)

		for tk := lex.NextToken(); tk.Type != token.EOF; tk = lex.NextToken() {
			fmt.Printf("%v\n", tk)
		}
	}
}
