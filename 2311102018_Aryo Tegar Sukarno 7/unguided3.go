//2311102018
package main

import (
	"fmt"
)

func main() {
	var clubA, clubB string
	var scoreA, scoreB int
	var winners []string


	fmt.Print("Klub A: ")
	fmt.Scanln(&clubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&clubB)

	fmt.Println("\nMasukkan skor pertandingan (format: skorA skorB). Masukkan skor negatif untuk mengakhiri:")
	
	matchCount := 1
	for {
		fmt.Printf("Pertandingan %d: ", matchCount)
		fmt.Scan(&scoreA, &scoreB)

		
		if scoreA < 0 || scoreB < 0 {
			fmt.Println("Input skor negatif, program selesai.")
			break
		}
		if scoreA > scoreB {
			winners = append(winners, clubA)
			fmt.Printf("Hasil %d: %s\n", matchCount, clubA)
		} else if scoreB > scoreA {
			winners = append(winners, clubB)
			fmt.Printf("Hasil %d: %s\n", matchCount, clubB)
		} else {
			winners = append(winners, "Draw")
			fmt.Printf("Hasil %d: Draw\n", matchCount)
		}
		matchCount++
	}
	fmt.Println("\nHasil akhir pertandingan:")
	for i, winner := range winners {
		fmt.Printf("Hasil %d: %s\n", i+1, winner)
	}
}
