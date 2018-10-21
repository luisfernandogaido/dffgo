package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("./results.txt")
	if err != nil {
		log.Fatal(err)
	}
	dados := string(bytes)
	sDados := strings.Split(dados, "\r\n")
	n := len(sDados)
	repeticoes := make([][]string, 0)
	var repeticao []string
	for i := 0; i < n; i++ {
		l := sDados[i]
		if l == "" {
			break
		}
		if l[0:1] == "-" {
			if len(repeticao) > 0 {
				repeticoes = append(repeticoes, repeticao)
			}
			repeticao = make([]string, 0)
			continue
		}
		r := strings.TrimSpace(l)
		r = r[1:][:len(r)-2]
		repeticao = append(repeticao, r)
	}
	if len(repeticao) > 0 {
		repeticoes = append(repeticoes, repeticao)
	}
	for _, repeticao := range repeticoes {
		n := len(repeticao)
		for i := 1; i < n; i++ {
			if err := os.Remove(repeticao[i]); err != nil {
				log.Println(err)
			}
		}
	}
}
