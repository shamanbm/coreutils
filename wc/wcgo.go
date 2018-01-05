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
		tl, tw, tb, l, w, b          int
	)

	flag.BoolVar(&flagLine, "line", false, "imprime o total de linhas")
	flag.BoolVar(&flagWord, "word", false, "imprime o total de palavras")
	flag.BoolVar(&flagByte, "byte", false, "imprime o total de bytes")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Uso: %s [OPÇÕES] [ARQUIVO] ...\n", os.Args[0])
		fmt.Fprintln(os.Stderr, `
Imprime o total de linhas, palavras e bytes de um arquivo. Se mais de um arquivo
for especificado, imprime o sub-total de cada arquivo. Caso seja omitido lê os dados 
da entrada padrão.
`)
		flag.PrintDefaults()
	}

	flag.Parse()

	if !flagLine && !flagWord && !flagByte {
		flagLine, flagWord, flagByte = true, true, true
	}

	if flag.NArg() == 0 {
		l, w, b = count(os.Stdin, flagLine, flagWord, flagByte)
		fmt.Printf("%-6v %-6v %-6v\n", l, w, b)
	} else {
		for _, filename = range flag.Args() {
			if file, err = os.Open(filename); err != nil {
				fmt.Fprintf(os.Stderr, "erro: %s\n", err.Error())
				os.Exit(1)
			}

			defer file.Close()

			l, w, b = count(file, flagLine, flagWord, flagByte)
			tl += l
			tw += w
			tb += b
			fmt.Printf("%-6v %-6v %-6v %v\n", l, w, b, filename)
		}
		if flag.NArg() > 1 {
			fmt.Printf("%-6v %-6v %-6v total\n", tl, tw, tb)
		}
	}
	return
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
