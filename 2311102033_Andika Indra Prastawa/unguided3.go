package main

import (
	"bufio"
	"fmt"
	"os"
)

const NMAX int = 127

type tabel struct {
	tab [NMAX]rune
	m   int
}

func isiArray(t *tabel, n *int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan teks (akhiri dengan titik):")

	i := 0
	for i < NMAX {
		r, _, _ := reader.ReadRune()
		if r == '.' {
			break
		}
		t.tab[i] = r
		i++
	}
	t.m = i
	*n = i
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t.tab[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t.tab[i], t.tab[n-1-i] = t.tab[n-1-i], t.tab[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t.tab[i] != t.tab[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	fmt.Println("Teks:")
	cetakArray(tab, m)

	balikanArray(&tab, m)
	fmt.Println("Reverse teks:")
	cetakArray(tab, m)

	isPalindrome := palindrom(tab, m)
	fmt.Println("Palindrome?")
	fmt.Println(isPalindrome)
}
