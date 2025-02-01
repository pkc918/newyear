package repl

import (
	"bufio"
	"fmt"
	"github.com/newyear/lexer"
	"github.com/newyear/token"
	"io"
)

var PROMPT = ">> "

// Start 交互式解释器，从输入的源代码中读取，直到读取完一行代码，将其传递给词法分析器实例，接着输出词法分析器生成的词法单元，直到EOF结束
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		_, _ = fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			_, _ = fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
