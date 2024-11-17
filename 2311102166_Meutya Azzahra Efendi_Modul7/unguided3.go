// Meutya Azzahra Efendi
// 2311102166

package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

// Fungsi untuk membaca input karakter hingga titik
func bacaInputKarakter(t *tabel) int {
	var input rune
	var n int
	fmt.Println("Masukkan karakter-karakter (akhir dengan titik '.'): ")
	for {
		fmt.Scanf("%c", &input)
		if input == '.' { // titik sebagai penghenti input
			break
		}
		if n < NMAX {
			(*t)[n] = input
			n++
		}
	}
	return n
}

// Prosedur untuk mencetak array
func cetakArray(t tabel, n int) {
	fmt.Print("Teks: ")
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]))
	}
	fmt.Println()
}

// Prosedur untuk membalikkan urutan array
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

// Fungsi untuk memeriksa apakah array membentuk palindrom
func adalahPalindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel

	// Isi array tab dengan memanggil fungsi bacaInputKarakter
	m := bacaInputKarakter(&tab)

	// Cetak array tab sebelum dibalik
	cetakArray(tab, m)

	// Balikan isi array tab
	balikanArray(&tab, m)

	// Cetak array tab setelah dibalik
	fmt.Print("Reverse teks: ")
	cetakArray(tab, m)

	// Cek apakah array membentuk palindrom
	isPalindrom := adalahPalindrom(tab, m)
	fmt.Printf("Palindrom: %v\n", isPalindrom)
}