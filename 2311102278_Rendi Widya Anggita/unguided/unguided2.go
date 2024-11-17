package main

import (
	"fmt"
)

// Struct untuk menyimpan informasi klub dan skor
type Klub struct {
	nama string
	skor int
}

func main() {
	// Input nama klub
	var klub1, klub2 Klub
	fmt.Print("Masukkan nama Klub A: ")
	fmt.Scanln(&klub1.nama)
	fmt.Print("Masukkan nama Klub B: ")
	fmt.Scanln(&klub2.nama)

	// Array untuk menyimpan daftar pemenang
	var hasilPertandingan []string

	// Perulangan untuk memasukkan skor tiap pertandingan
	pertandingan := 1
	for {
		fmt.Printf("\nPertandingan %d\n", pertandingan)
		fmt.Printf("Masukkan skor untuk %s: ", klub1.nama)
		fmt.Scanln(&klub1.skor)
		fmt.Printf("Masukkan skor untuk %s: ", klub2.nama)
		fmt.Scanln(&klub2.skor)

		// Periksa validitas skor
		if klub1.skor < 0 || klub2.skor < 0 {
			fmt.Println("\nSkor tidak valid! Proses dihentikan.")
			break
		}

		// Tentukan hasil pertandingan
		if klub1.skor > klub2.skor {
			hasilPertandingan = append(hasilPertandingan, klub1.nama)
			fmt.Printf("Hasil pertandingan %d: %s\n", pertandingan, klub1.nama)
		} else if klub2.skor > klub1.skor {
			hasilPertandingan = append(hasilPertandingan, klub2.nama)
			fmt.Printf("Hasil pertandingan %d: %s\n", pertandingan, klub2.nama)
		} else {
			hasilPertandingan = append(hasilPertandingan, "Draw")
			fmt.Printf("Hasil pertandingan %d: Draw\n", pertandingan)
		}

		pertandingan++
	}

	// Tampilkan hasil semua pertandingan
	fmt.Println("\nHasil seluruh pertandingan:")
	for i, hasil := range hasilPertandingan {
		if hasil == "Draw" {
			fmt.Printf("Pertandingan %d: Draw\n", i+1)
		} else {
			fmt.Printf("Pertandingan %d dimenangkan oleh: %s\n", i+1, hasil)
		}
	}
}
