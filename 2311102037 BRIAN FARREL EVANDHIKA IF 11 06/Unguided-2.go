//2311102037_BRIAN FARREL EVANDHIKA_IF 11 06
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Menyimpan nama klub yang menang
	var winners []string

	fmt.Print("Masukkan nama Klub A: ")
	clubA, _ := reader.ReadString('\n')
	clubA = strings.TrimSpace(clubA)

	fmt.Print("Masukkan nama Klub B: ")
	clubB, _ := reader.ReadString('\n')
	clubB = strings.TrimSpace(clubB)

	// Loop untuk menerima skor pertandingan
	for {
		var scoreA, scoreB int

		fmt.Print("\nMasukkan skor Klub A (atau negatif untuk berhenti): ")
		fmt.Scan(&scoreA)
		if scoreA < 0 {
			break
		}

		fmt.Print("Masukkan skor Klub B (atau negatif untuk berhenti): ")
		fmt.Scan(&scoreB)
		if scoreB < 0 {
			break
		}

		// Menentukan pemenang
		if scoreA > scoreB {
			winners = append(winners, clubA)
			fmt.Printf("Hasil: %s menang!\n", clubA)
		} else if scoreB > scoreA {
			winners = append(winners, clubB)
			fmt.Printf("Hasil: %s menang!\n", clubB)
		} else {
			fmt.Println("Hasil: Seri (Draw). Tidak ada pemenang!")
		}
	}

	// Menampilkan daftar pemenang
	fmt.Println("\nDaftar klub yang memenangkan pertandingan:")
	for i, winner := range winners {
		fmt.Printf("Pertandingan %d: %s\n", i+1, winner)
	}

	fmt.Println("Program selesai.")
}
