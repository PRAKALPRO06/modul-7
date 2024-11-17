package main

import (
	"fmt"
)

func main() {
	var klub1, klub2 string
	fmt.Print("Masukkan nama Klub A: ")
	fmt.Scanln(&klub1)
	fmt.Print("Masukkan nama Klub B: ")
	fmt.Scanln(&klub2)

	var skor1, skor2 int
	var hasil []string

	for pertandingan := 1; ; pertandingan++ {
		fmt.Printf("\nPertandingan %d:\n", pertandingan)
		fmt.Printf("Masukkan skor untuk %s: ", klub1)
		fmt.Scan(&skor1)
		fmt.Printf("Masukkan skor untuk %s: ", klub2)
		fmt.Scan(&skor2)

		if skor1 < 0 || skor2 < 0 {
			fmt.Println("\nSkor negatif terdeteksi. Proses dihentikan.")
			break
		}

		if skor1 > skor2 {
			hasil = append(hasil, fmt.Sprintf("Hasil %d: %s", pertandingan, klub1))
		} else if skor2 > skor1 {
			hasil = append(hasil, fmt.Sprintf("Hasil %d: %s", pertandingan, klub2))
		} else {
			hasil = append(hasil, fmt.Sprintf("Hasil %d: Draw", pertandingan))
		}
	}

	fmt.Println("\nHasil Pertandingan:")
	for _, h := range hasil {
		fmt.Println(h)
	}
}
