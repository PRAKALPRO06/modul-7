package main

import (
	"fmt"
)

func main() {
	// Meminta input nama klub
	var klubA, klubB string
	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)

	// Slice untuk menyimpan hasil pertandingan
	var hasil []string

	// Memasukkan skor pertandingan
	fmt.Println("Masukkan skor pertandingan (skor negatif untuk berhenti):")
	pertandingan := 1
	for {
		var skorA, skorB int
		fmt.Printf("Pertandingan %d (%s vs %s): ", pertandingan, klubA, klubB)
		fmt.Scanf("%d %d\n", &skorA, &skorB)

		// Berhenti jika salah satu skor negatif
		if skorA < 0 || skorB < 0 {
			break
		}

		// Menentukan pemenang
		if skorA > skorB {
			hasil = append(hasil, klubA)
			fmt.Printf("Hasil %d: %s menang\n", pertandingan, klubA)
		} else if skorB > skorA {
			hasil = append(hasil, klubB)
			fmt.Printf("Hasil %d: %s menang\n", pertandingan, klubB)
		} else {
			hasil = append(hasil, "Draw")
			fmt.Printf("Hasil %d: Seri\n", pertandingan)
		}

		pertandingan++
	}

	// Menampilkan daftar klub yang menang
	fmt.Println("\nDaftar klub yang menang:")
	for i, pemenang := range hasil {
		if pemenang != "Draw" {
			fmt.Printf("Hasil %d: %s\n", i+1, pemenang)
		}
	}

	fmt.Println("Pertandingan selesai.")
}