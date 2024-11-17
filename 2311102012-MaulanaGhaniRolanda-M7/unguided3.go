//2311102012
package main

import (
	"fmt"
)

const NMAX = 127

type tabel [NMAX]rune

// Prosedur untuk mengisi array dengan karakter dari input pengguna
func isiArray(t *tabel, n *int) {
	fmt.Print("Teks: ")
	var input string
	fmt.Scanln(&input)

	*n = len(input)
	for i := 0; i < *n; i++ {
		(*t)[i] = rune(input[i])
	}
}

// Prosedur untuk mencetak isi array
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]))
	}
	fmt.Println()
}

// Prosedur untuk membalik isi array
func balikArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		(*t)[i], (*t)[n-i-1] = (*t)[n-i-1], (*t)[i]
	}
}

// Fungsi untuk memeriksa apakah array membentuk palindrome
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

	// Memanggil prosedur untuk mengisi array
	isiArray(&tab, &n)

	// Menampilkan isi array asli
	fmt.Print("Teks: ")
	cetakArray(tab, n)

	// Membalik isi array
	balikArray(&tab, n)

	// Menampilkan isi array yang sudah dibalik
	fmt.Print("Reverse: ")
	cetakArray(tab, n)

	// Memeriksa apakah array membentuk palindrome
	if palindrom(tab, n) {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Bukan Palindrom")
	}
}
