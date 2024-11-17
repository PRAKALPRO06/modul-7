package main

import (
	"fmt"
)

const MAX int = 127

type tabel [MAX]rune

func isiArray(t *tabel, n *int) {

	var kar rune
	*n = 0
	fmt.Print("Masukkan teks (akhiri dengan titik): ")
	fmt.Scanf("%c", &kar)
	for kar != '.' && *n < MAX {
		t[*n] = kar
		*n++
		fmt.Scanf("%c", &kar)
	}
}

func cetakArray(t *tabel, n int) {

	fmt.Print("Teks: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikkanArray(t *tabel, n int) {

	var temp tabel
	for i := 0; i < n; i++ {
		temp[i] = t[n-1-i]
	}
	for i := 0; i < n; i++ {
		t[i] = temp[i]
	}
}

func palindrom(t *tabel, n int) bool {

	var temp tabel
	for i := 0; i < n; i++ {
		temp[i] = t[i]
	}

	balikkanArray(&temp, n)

	for i := 0; i < n; i++ {
		if t[i] != temp[i] {
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
	cetakArray(&tab, n)

	fmt.Print("Reverse teks: ")
	balikkanArray(&tab, n)
	cetakArray(&tab, n)

	balikkanArray(&tab, n)

	fmt.Print("Palindrom: ")
	if palindrom(&tab, n) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
