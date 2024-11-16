package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var karakter rune
	i := 0
	fmt.Print("Masukkan karakter: ")
	for fmt.Scanf("%c", &karakter); karakter != '.' && i < NMAX; fmt.Scanf("%c", &karakter) {
		t[i] = karakter
		i++
	}
	*n = i
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikkanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	temp := t 
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
	fmt.Print("Teks\t\t: ")
	cetakArray(tab, n)

	balikkanArray(&tab, n)
	fmt.Print("Reverse teks\t: ")
	cetakArray(tab, n)

	fmt.Print("Palindrom\t: ")
	if palindrom(tab, n) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}