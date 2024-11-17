package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

// Fungsi untuk mengecek palindrom
func isPalindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func isiArray(t *tabel, n *int) {
	var char rune
	*n = 0
	fmt.Print("Masukkan karakter (akhiri dengan titik): ")

	for {
		_, err := fmt.Scanf("%c", &char)
		if err != nil || char == '.' || *n >= NMAX {
			break
		}

		// Mengabaikan karakter newline dan spasi
		if char == '\n' || char == '\r' || char == ' ' {
			continue
		}

		t[*n] = char
		*n++
	}
}

func cetakArray(t tabel, n int) {
	fmt.Print("Isi array: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func main() {
	var tab tabel
	var n int

	isiArray(&tab, &n)

	isPalindrome := isPalindrom(tab, n)

	fmt.Println("\nArray sebelum dibalik:")
	cetakArray(tab, n)

	if isPalindrome {
		fmt.Println("Kata yang dimasukkan adalah palindrom")
	} else {
		fmt.Println("Kata yang dimasukkan bukan palindrom")
	}

	balikanArray(&tab, n)

	fmt.Println("\nArray setelah dibalik:")
	cetakArray(tab, n)
}
