package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var char rune
	*n = 0
	fmt.Println("Masukkan karakter ke dalam array (akhiri dengan karakter '.'): ")

	for {
		fmt.Scanf("%c\n", &char)
		if char == '.' || *n >= NMAX {
			break
		}
		t[*n] = char
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	fmt.Print("Teks: ")
	cetakArray(tab, m)

	balikanArray(&tab, m)

	fmt.Print("Reverse teks: ")
	cetakArray(tab, m)

	isPalindrom := palindrom(tab, m)
	if isPalindrom {
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}
}
