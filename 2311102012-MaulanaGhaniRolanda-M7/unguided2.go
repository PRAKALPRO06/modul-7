// 2311102012
package main

import (
	"fmt"
)

func main() {
	var clubA, clubB string
	var scoreA, scoreB int
	var winners []string

	// Meminta nama klub
	fmt.Print("Klub A: ")
	fmt.Scanln(&clubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&clubB)

	// Memasukkan skor pertandingan
	for i := 1; ; i++ {
		fmt.Printf("Pertandingan %d (skor %s): ", i, clubA)
		fmt.Scanln(&scoreA)
		fmt.Printf("Pertandingan %d (skor %s): ", i, clubB)
		fmt.Scanln(&scoreB)

		// Memeriksa apakah skor valid
		if scoreA < 0 || scoreB < 0 {
			fmt.Println("Input skor tidak valid. Program selesai.")
			break
		}

		// Menentukan pemenang
		if scoreA > scoreB {
			winners = append(winners, clubA)
		} else if scoreA < scoreB {
			winners = append(winners, clubB)
		} else {
			winners = append(winners, "Draw")
		}
	}

	// Menampilkan hasil
	fmt.Println("\nHasil pertandingan:")
	for i, winner := range winners {
		fmt.Printf("Pertandingan %d: %s\n", i+1, winner)
	}
}
