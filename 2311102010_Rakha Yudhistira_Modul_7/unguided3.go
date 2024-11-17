//2311102010_RAKHA YUDHISTIRA_IF-11-06

package main

import (
	"fmt"
)

const NMAX int = 127

type tabel struct {
	tab [NMAX]rune
	m   int
}

// Fungsi untuk mengisi array
func isiArray(t *tabel, n *int) {
	var input rune
	*n = 0
	fmt.Println("Masukkan karakter (akhiri dengan '.'): ")
	for {
		fmt.Scanf("%c\n", &input)
		if input == '.' || *n >= NMAX {
			break
		}
		t.tab[*n] = input
		*n++
	}
	t.m = *n
}

// Fungsi untuk mencetak array
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t.tab[i])
	}
	fmt.Println()
}

// Fungsi untuk membalik isi array
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t.tab[i], t.tab[n-1-i] = t.tab[n-1-i], t.tab[i]
	}
}

// Fungsi untuk memeriksa apakah array adalah palindrom
func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t.tab[i] != t.tab[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	// Isi array
	isiArray(&tab, &n)

	// Cetak array
	fmt.Print("Teks: ")
	cetakArray(tab, n)

	// Periksa apakah array adalah palindrom
	if palindrom(tab, n) {
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}
}