package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Membaca nama klub
	clubA := readInput(reader, "Masukkan nama Klub A: ")
	clubB := readInput(reader, "Masukkan nama Klub B: ")

	var winners []string

	for {
		// Membaca skor
		scoreA, stop := readScore("Masukkan skor Klub A (atau negatif untuk berhenti): ")
		if stop {
			break
		}
		scoreB, stop := readScore("Masukkan skor Klub B (atau negatif untuk berhenti): ")
		if stop {
			break
		}

		// Menentukan pemenang
		switch {
		case scoreA > scoreB:
			winners = append(winners, clubA)
			fmt.Printf("Hasil: %s menang!\n", clubA)
		case scoreB > scoreA:
			winners = append(winners, clubB)
			fmt.Printf("Hasil: %s menang!\n", clubB)
		default:
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

// Fungsi untuk membaca input string
func readInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Fungsi untuk membaca skor dan memeriksa penghentian
func readScore(prompt string) (int, bool) {
	var score int
	fmt.Print(prompt)
	fmt.Scan(&score)
	if score < 0 {
		return 0, true
	}
	return score, false
}
