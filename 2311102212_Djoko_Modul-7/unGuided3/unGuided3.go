package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NMAX int = 127

type tabel struct {
	tab [NMAX]rune
	m  int 
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t.tab[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t.tab[i], t.tab[n-i-1] = t.tab[n-i-1], t.tab[i]
	}
}

func isiArray(t *tabel) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan teks: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	t.m = len(input)
	for i, c := range input {
		t.tab[i] = c
	}
}

func main() {
	var tab tabel

	isiArray(&tab)

	fmt.Print("Teks: ")
	cetakArray(tab, tab.m)

	balikanArray(&tab, tab.m)

	fmt.Print("Reverse teks: ")
	cetakArray(tab, tab.m)
}