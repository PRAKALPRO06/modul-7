package main

import (
	"fmt"
)

// Konstanta maksimum jumlah elemen array
const NMAX int = 127

// Struct untuk merepresentasikan array dan jumlah elemen
type Tabel struct {
	tab [NMAX]rune // Array untuk menyimpan karakter
	n   int        // Jumlah elemen yang valid dalam array
}

// Fungsi untuk mengisi array dengan karakter dari input pengguna
func isiArray(t *Tabel) {
	fmt.Println("Masukkan karakter untuk kata (diakhiri dengan TITIK '.'):")

	for {
		var input rune
		fmt.Scanf("%c", &input)

		if input == '.' || t.n >= NMAX { // Berhenti jika ditemukan titik atau jumlah elemen maksimal tercapai
			break
		}
		t.tab[t.n] = input
		t.n++
	}
}

// Fungsi untuk mencetak isi array sebagai string
func cetakArray(t Tabel) {
	for i := 0; i < t.n; i++ {
		fmt.Printf("%c", t.tab[i])
	}
	fmt.Println()
}

// Fungsi untuk membalikkan urutan elemen dalam array
func balikkanArray(t Tabel) Tabel {
	var reversed Tabel
	reversed.n = t.n
	for i := 0; i < t.n; i++ {
		reversed.tab[i] = t.tab[t.n-1-i]
	}
	return reversed
}

// Fungsi untuk memeriksa apakah array membentuk palindrom
func isPalindrom(t Tabel) bool {
	for i, j := 0, t.n-1; i < j; i, j = i+1, j-1 {
		if t.tab[i] != t.tab[j] {
			return false
		}
	}
	return true
}

func main() {
	var tab Tabel

	// Meminta input dari pengguna
	isiArray(&tab)

	// Menampilkan teks asli
	fmt.Print("Teks: ")
	cetakArray(tab)

	// Membalikkan array dan menampilkan hasilnya
	reversedTab := balikkanArray(tab)
	fmt.Print("Reverse teks: ")
	cetakArray(reversedTab)

	// Menentukan apakah teks merupakan palindrom
	if isPalindrom(tab) {
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}
}
