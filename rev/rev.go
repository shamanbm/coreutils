/*
	app: rev
	Descrição: reverte a sequência de caracteres de uma expressão, seja ela
	lida a partir de um ou mais arquivos ou por meio da entrada padrão.
	Um exemplo ridículo de aplicação desenvolvida utilizando a forma idiomática do Go,
	que alias ainda estou me acostumando. Foda-se 'Hello World'.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var fs *bufio.Scanner

	if len(os.Args) == 1 {
		fs = bufio.NewScanner(os.Stdin)
		for fs.Scan() {
			fmt.Println(strReverse(fs.Text()))
		}
		checkError(fs.Err())
	} else {
		for _, filename := range os.Args[1:] {
			_, err := os.Stat(filename)
			checkError(err)

			f, err := os.Open(filename)
			checkError(err)

			fs = bufio.NewScanner(f)

			for fs.Scan() {
				fmt.Println(strReverse(fs.Text()))
			}

			checkError(fs.Err())
			f.Close()
		}
	}
}

func strReverse(s string) string {
	tmp := make([]byte, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		tmp = append(tmp, s[i])
	}
	return string(tmp)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: erro: %s\n", os.Args[0], err.Error())
		os.Exit(1)
	}
}
