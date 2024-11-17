package main

import (
	"fmt"
)

const NMAX = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input string
	*n = 0
	for {
		fmt.Scan(&input)
		if input == "." {
			break
		}
		for _, c := range input {
			if *n < NMAX {
				(*t)[*n] = c
				*n++
			}
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]), " ")
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	i := 0
	j := n - 1
	for i < j {
		t[i], t[j] = t[j], t[i]
		i++
		j--
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
	var t tabel
	var n int

	fmt.Print("Teks : ")
	isiArray(&t, &n)

	balikanArray(&t, n)

	fmt.Print("Reverse teks : ")
	cetakArray(t, n)

	fmt.Println("Palindrom ?", palindrom(t, n))
}
