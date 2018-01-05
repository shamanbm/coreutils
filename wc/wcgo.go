package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	var (
		flagLine, flagWord, flagByte bool
		filename                     string
		file                         *os.File
		err                          error
	)

	flag.BoolVar(&flagLine, "line", false, "imprime o total de linhas")
	flag.BoolVar(&flagWord, "word", false, "imprime o total de palavras")
	flag.BoolVar(&flagByte, "byte", false, "imprime o total de bytes")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Uso: %s [OPÇÕES] [ARQUIVO] ...\n", os.Args[0])
		fmt.Fprintln(os.Stderr, `
Imprime o total de linhas, palavras e bytes de um arquivo. Se mais de um arquivo
for especificado, imprime o sub-total e caso seja omitido lê os dados da entrada padrão
`)
		flag.PrintDefaults()
	}

	flag.Parse()

	if !(flagLine && flagWord && flagByte) {
		flagLine, flagWord, flagByte = true, true, true
	}

	if flag.NArg() == 0 {
		fmt.Println(count(os.Stdin, flagLine, flagWord, flagByte))
	} else {
		for _, filename = range flag.Args() {
			if file, err = os.Open(filename); err != nil {
				fmt.Fprintf(os.Stderr, "erro: %s\n", err.Error())
				os.Exit(1)
			}
			fmt.Println(count(file, flagLine, flagWord, flagByte))
		}
	}
}

func count(r io.Reader, line, word, byte bool) (int, int, int) {
	var (
		scanner *bufio.Scanner
		l, w, b int
		err     error
	)

	scanner = bufio.NewScanner(r)

	for scanner.Scan() {
		if line {
			l++
		}
		if word {
			w += len(strings.Split(scanner.Text(), " "))
		}
		if byte {
			b += len(scanner.Text()) + 1
		}
	}

	if err = scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "erro: %s\n", err.Error())
		os.Exit(1)
	}

	return l, w, b
}
