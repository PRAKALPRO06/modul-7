package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

var tab tabel
var m int

func isiArray(t *tabel, n *int) {
	var input string
	fmt.Scan(&input)
	*n = len(input)

	for i := 0; i < *n; i++ {
		t[i] = rune(input[i])
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
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	var tbalikan tabel
	copy(tbalikan[:], t[:])

	balikanArray(&tbalikan, n)

	for i := 0; i < n; i++ {
		if t[i] != tbalikan[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print("Teks: ")
	isiArray(&tab, &m)

	fmt.Print("Reverse teks: ")
	balikanArray(&tab, m)
	cetakArray(tab, m)

	fmt.Print("Palindrom? ")
	fmt.Println(palindrom(tab, m))
}
