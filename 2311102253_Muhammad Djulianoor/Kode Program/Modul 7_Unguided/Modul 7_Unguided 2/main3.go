package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil []string

	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)

	fmt.Println("\nMasukkan skor pertandingan antara", klubA, "dan", klubB)
	fmt.Println("Masukkan skor negatif untuk menghentikan input.")
	pertandingan := 1
	for {
		fmt.Printf("Pertandingan %d - Skor %s: ", pertandingan, klubA)
		fmt.Scan(&skorA)
		fmt.Printf("Pertandingan %d - Skor %s: ", pertandingan, klubB)
		fmt.Scan(&skorB)
		if skorA < 0 || skorB < 0 {
			fmt.Println("Pertandingan selesai.")
			break
		}
		if skorA > skorB {
			hasil = append(hasil, klubA)
			fmt.Printf("Hasil %d: %s menang\n", pertandingan, klubA)
		} else if skorA < skorB {
			hasil = append(hasil, klubB)
			fmt.Printf("Hasil %d: %s menang\n", pertandingan, klubB)
		} else {
			hasil = append(hasil, "Draw")
			fmt.Printf("Hasil %d: Draw\n", pertandingan)
		}
		pertandingan++
	}

	fmt.Println("\nHasil Pertandingan:")
	for i, h := range hasil {
		fmt.Printf("Hasil %d: %s\n", i+1, h)
	}
}
