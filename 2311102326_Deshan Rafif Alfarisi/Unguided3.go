package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

// Fungsi untuk mengisi array dengan karakter
func isiArray(t *tabel, n *int) {
	var ch rune
	fmt.Println("Masukkan teks (akhiri dengan '.'): ")
	for *n = 0; *n < NMAX; *n++ {
		fmt.Scanf("%c", &ch)
		if ch == '.' { // Berhenti saat menemukan titik
			break
		}
		t[*n] = ch
	}
}

// Fungsi untuk mencetak isi array
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

// Fungsi untuk membalikkan isi array
func balikkanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

// Fungsi untuk memeriksa apakah array membentuk palindrome
func palindrome(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	// Memanggil fungsi isiArray untuk mengisi data
	isiArray(&tab, &n)

	// Mencetak isi array asli
	fmt.Print("Teks: ")
	cetakArray(tab, n)

	// Membalikkan isi array
	balikkanArray(&tab, n)
	fmt.Print("Reverse teks: ")
	cetakArray(tab, n)

	// Memeriksa apakah teks merupakan palindrome
	if palindrome(tab, n) {
		fmt.Println("Palindrome: true")
	} else {
		fmt.Println("Palindrome: false")
	}
}
