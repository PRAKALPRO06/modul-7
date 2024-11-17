	package main

	import (
		"bufio"
		"fmt"
		"os"
		"strings"
	)

	const MAX = 127

	type Tabel [MAX]rune

	// Fungsi untuk mengisi array
	func isiArray(tab *Tabel, n *int) {
		var input string
		fmt.Print("Masukkan teks (akhiri dengan titik): ")
		reader := bufio.NewReader(os.Stdin)
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		// Memasukkan karakter ke array hingga menemukan '.'
		*n = 0
		for _, char := range input {
			if char == '.' {
				break
			}
			if *n < MAX {
				tab[*n] = char
				*n++
			}
		}
	}

	// Fungsi untuk mencetak array
	func cetakArray(tab Tabel, n int) {
		for i := 0; i < n; i++ {
			fmt.Print(string(tab[i]))
		}
		fmt.Println()
	}

	// Fungsi untuk membalik isi array
	func balikkanArray(tab *Tabel, n int) {
		for i := 0; i < n/2; i++ {
			tab[i], tab[n-1-i] = tab[n-1-i], tab[i]
		}
	}

	// Fungsi untuk memeriksa apakah array merupakan palindrom
	func palindrom(tab Tabel, n int) bool {
		for i := 0; i < n/2; i++ {
			if tab[i] != tab[n-1-i] {
				return false
			}
		}
		return true
	}

	// Fungsi utama
	func main() {
		var tab Tabel
		var n int

		// Mengisi array
		isiArray(&tab, &n)

		// Menampilkan isi array
		fmt.Print("Teks: ")
		cetakArray(tab, n)

		// Membalik array
		var reversedTab Tabel
		copy(reversedTab[:], tab[:])
		balikkanArray(&reversedTab, n)

		// Menampilkan array yang telah dibalik
		fmt.Print("Reverse: ")
		cetakArray(reversedTab, n)

		// Memeriksa apakah palindrom
		if palindrom(tab, n) {
			fmt.Println("Palindrom: YA")
		} else {
			fmt.Println("Palindrom: TIDAK")
		}
	}