package main

import (
	"fmt"
	// Nama : Egi Umar Ferdhika
	// NIM 	: 2311102277
)

const NMAX int = 127

// Struct untuk menampung array karakter
type tabel [NMAX]rune

// Fungsi untuk mengisi array
func isiArray(t *tabel, n *int) {
	fmt.Print("Masukkan teks: ")
	var input string
	fmt.Scan(&input)

	*n = 0
	// Mengisi array dengan karakter dari input
	for _, char := range input {
		if char != '.' && *n < NMAX {
			t[*n] = char
			*n++
		}
	}
}

// Fungsi untuk mencetak array
func cetakArray(t tabel, n int) {
	fmt.Print("Teks\t\t:")
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

// Fungsi untuk membalik array
func balikanArray(t *tabel, n int) {
	reversed := tabel{}
	for i := 0; i < n; i++ {
		reversed[i] = t[n-1-i]
	}

	fmt.Print("Reverse teks\t:")
	for i := 0; i < n; i++ {
		fmt.Printf("%c", reversed[i])
		t[i] = reversed[i] // Update array asli dengan hasil reversed
	}
	fmt.Println()
}

// Fungsi untuk mengecek palindrom
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

	// Input dan tampilkan array original
	isiArray(&tab, &m)
	cetakArray(tab, m)

	// Balik array dan tampilkan
	balikanArray(&tab, m)

	// Cek palindrom dan tampilkan hasil
	fmt.Print("Palindrom\t:")
	if palindrom(tab, m) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
