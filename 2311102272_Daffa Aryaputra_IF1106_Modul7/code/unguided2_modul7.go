package main

import "fmt"

func main() {

	var klubA, klubB string
	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)

	var pemenang []string
	var pertandingan int = 1

	for {
		var skorA, skorB int
		fmt.Printf("Pertandingan %d (Skor %s vs %s): ", pertandingan, klubA, klubB)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)
			fmt.Printf("Hasil %d: %s\n", pertandingan, klubA)
		} else if skorA < skorB {
			pemenang = append(pemenang, klubB)
			fmt.Printf("Hasil %d: %s\n", pertandingan, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
			fmt.Printf("Hasil %d: Draw\n", pertandingan)
		}

		pertandingan++
	}

	fmt.Println("\nDaftar klub yang memenangkan pertandingan:")
	for i, hasil := range pemenang {
		fmt.Printf("Pertandingan %d: %s\n", i+1, hasil)
	}

	fmt.Println("Pertandingan selesai.")
}