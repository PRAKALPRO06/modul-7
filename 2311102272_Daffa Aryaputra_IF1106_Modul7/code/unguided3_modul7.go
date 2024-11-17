package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input string
	fmt.Println("Masukkan teks (akhiri dengan titik '.'): ")
	fmt.Scanln(&input)

	*n = 0
	for _, char := range input {
		if char == '.' {
			break
		}
		t[*n] = char
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]))
	}
	fmt.Println()
}

func balikkanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	isiArray(&tab, &n)

	fmt.Print("Teks: ")
	cetakArray(tab, n)

	balikkanArray(&tab, n)

	fmt.Print("Reverse teks: ")
	cetakArray(tab, n)

	isPalindrom := palindrom(tab, n)
	fmt.Print("Palindrom: ")
	fmt.Println(isPalindrom)
}