package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NMAX int = 127

type tabel struct {
	tab [NMAX]rune
	m   int
}

// Fungsi untuk mengisi array
func isiArray(t *tabel) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan karakter (akhiri dengan '.'): ")

	input, _ := reader.ReadString('.')
	input = strings.TrimSuffix(input, ".")      // Menghapus tanda akhir
	input = strings.ReplaceAll(input, "\n", "") // Menghapus newline
	input = strings.ReplaceAll(input, "\r", "") // Menghapus carriage return

	if len(input) > NMAX {
		input = input[:NMAX] // Membatasi panjang maksimal
	}

	for i, char := range input {
		t.tab[i] = char
	}
	t.m = len(input)
}

// Fungsi untuk mencetak array
func cetakArray(t tabel) {
	for i := 0; i < t.m; i++ {
		fmt.Printf("%c", t.tab[i])
	}
	fmt.Println()
}

// Fungsi untuk membalik isi array
func balikanArray(t *tabel) {
	for i := 0; i < t.m/2; i++ {
		t.tab[i], t.tab[t.m-1-i] = t.tab[t.m-1-i], t.tab[i]
	}
}

// Fungsi untuk memeriksa apakah array adalah palindrom
func palindrom(t tabel) bool {
	for i := 0; i < t.m/2; i++ {
		if t.tab[i] != t.tab[t.m-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel

	// Isi array
	isiArray(&tab)

	// Cetak array
	fmt.Print("Teks: ")
	cetakArray(tab)

	// Periksa apakah array adalah palindrom
	if palindrom(tab) {
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}
}
