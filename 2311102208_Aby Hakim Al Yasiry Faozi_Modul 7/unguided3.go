package main

import "fmt"

const NMAX = 127

type Tabel struct {
	tab [NMAX]rune
	n   int
}

func isiArray(t *Tabel) {
	fmt.Print("Teks: ")
	var input string
	fmt.Scanln(&input)
	for i, char := range input {
		t.tab[i] = char
	}
	t.n = len(input)
}

func cetakArray(t Tabel) {
	for i := 0; i < t.n; i++ {
		fmt.Print(string(t.tab[i]))
	}
	fmt.Println()
}

func balikkanArray(t *Tabel) {
	for i, j := 0, t.n-1; i < j; i, j = i+1, j-1 {
		t.tab[i], t.tab[j] = t.tab[j], t.tab[i]
	}
}

func palindrome(t Tabel) bool {
	for i, j := 0, t.n-1; i < j; i, j = i+1, j-1 {
		if t.tab[i] != t.tab[j] {
			return false
		}
	}
	return true
}

func main() {
	var tab Tabel

	isiArray(&tab)

	fmt.Print("Reverse teks: ")
	cetakArray(tab)

	balikkanArray(&tab)

	fmt.Print("Reverse teks: ")
	cetakArray(tab)

	if palindrome(tab) {
		fmt.Println("Palindrome: true")
	} else {
		fmt.Println("Palindrome: false")
	}
}
