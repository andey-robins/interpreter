package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/andey-robins/interpreter/lexer"
	"github.com/andey-robins/interpreter/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lex := lexer.New(line)
		par := parser.New(lex)

		program := par.ParseProgram()
		if len(par.Errors()) != 0 {
			printParserErrors(out, par.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
