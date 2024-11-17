// 2311102266_Hanif Reyhan Zhafran Arytona

package main

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

	for i := 1; i <= 3; i++ {
		fmt.Printf("Pertandingan %d - Masukkan skor %s: ", i, klubA)
		fmt.Scan(&skorA)
		fmt.Printf("Pertandingan %d - Masukkan skor %s: ", i, klubB)
		fmt.Scan(&skorB)

		if skorA < 0 || skorB < 0 {
			fmt.Println("Skor tidak valid. Pertandingan dibatalkan.")
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)
			fmt.Printf("Hasil pertandingan %d: %s menang\n", i, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
			fmt.Printf("Hasil pertandingan %d: %s menang\n", i, klubB)
		} else {
			fmt.Printf("Hasil pertandingan %d: Seri (Draw)\n", i)
		}
	}

	if len(pemenang) > 0 {
		fmt.Println("\nDaftar klub yang memenangkan pertandingan:")
		for _, klub := range pemenang {
			fmt.Println(klub)
		}
	} else {
		fmt.Println("\nTidak ada pemenang dalam pertandingan.")
	}
}
