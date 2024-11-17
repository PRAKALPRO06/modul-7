package main

// Erwin Rivaldo Silaban
// 2311102245

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang []string

	fmt.Print("Masukkan nama Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Masukkan nama Klub B: ")
	fmt.Scanln(&klubB)

	for i := 1; ; i++ {
		fmt.Printf("Pertandingan %d - Masukkan skor %v dan %v : ", i, klubA, klubB)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			fmt.Println("Skor tidak valid. Pertandingan selesai.")
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)

		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)

		} else {
			pemenang = append(pemenang, "Draw")

		}
	}

	fmt.Println("\nDaftar klub yang memenangkan pertandingan:")
	for i := 1; i <= len(pemenang); i++ {
		fmt.Printf("Hasil %v : %v\n", i, pemenang[i-1])
	}
}
