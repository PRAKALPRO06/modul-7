package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	klubMenang := []string{}
	pertandingan := 1

	fmt.Print("Masukkan nama Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Masukkan nama Klub B: ")
	fmt.Scan(&klubB)

	for {
		fmt.Printf("\nPertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			klubMenang = append(klubMenang, klubA)
		} else if skorB > skorA {
			klubMenang = append(klubMenang, klubB)
		} else {
			klubMenang = append(klubMenang, "Draw")
		}

		pertandingan++
	}

	fmt.Println("\nDaftar Hasil Pertandingan:")
	for i, pemenang := range klubMenang {
		fmt.Printf("Hasil %d : %s\n", i+1, pemenang)
	}

	fmt.Println("\nPertandingan selesai")
}
