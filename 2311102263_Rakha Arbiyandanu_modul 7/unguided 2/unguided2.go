package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	fmt.Print("Masukkan nama klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Masukkan nama klub B: ")
	fmt.Scanln(&klubB)

	pemenang := []string{}
	for i := 1; ; i++ {
		var skorA, skorB int
		fmt.Printf("Pertandingan %d (Skor %s vs %s): ", i, klubA, klubB)
		fmt.Scanln(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			fmt.Println("Skor tidak valid, proses input selesai.")
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)
			fmt.Printf("Hasil %d: %s\n", i, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
			fmt.Printf("Hasil %d: %s\n", i, klubB)
		} else {
			fmt.Printf("Hasil %d: Draw\n", i)
		}
	}


	fmt.Println("\nDaftar klub yang memenangkan pertandingan:")
	for _, klub := range pemenang {
		fmt.Println(klub)
	}
}
