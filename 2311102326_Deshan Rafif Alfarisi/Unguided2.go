package main

import (
	"fmt"
)

func main() {
	// Meminta pengguna memasukkan nama kedua klub
	var klubA, klubB string
	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)

	// Array untuk menyimpan hasil pemenang
	var pemenang []string

	// Memulai loop untuk memasukkan skor pertandingan
	pertandingan := 1
	for {
		var skorA, skorB int
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		// Jika salah satu skor negatif, hentikan input
		if skorA < 0 || skorB < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}

		// Menentukan pemenang atau hasil draw
		if skorA > skorB {
			fmt.Printf("// %s menang\n", klubA)
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			fmt.Printf("// %s menang\n", klubB)
			pemenang = append(pemenang, klubB)
		} else {
			fmt.Println("// Draw")
			pemenang = append(pemenang, "Draw")
		}

		// Naikkan nomor pertandingan
		pertandingan++
	}

	// Menampilkan hasil akhir
	fmt.Println("\nHasil pertandingan:")
	for i, hasil := range pemenang {
		fmt.Printf("Hasil %d: %s\n", i+1, hasil)
	}
}
