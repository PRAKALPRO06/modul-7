package main

import (
	"fmt"
	// Nama : Egi Umar Ferdhika
	// NIM 	: 2311102277
)

type Match struct {
	score1 int
	score2 int
}

func main() {
	var club1, club2 string
	fmt.Print("Masukkan nama klub 1: ")
	fmt.Scan(&club1)
	fmt.Print("Masukkan nama klub 2: ")
	fmt.Scan(&club2)

	matches := make([]Match, 0)
	winners := make([]string, 0)

	for {
		var match Match
		fmt.Printf("\nPertandingan %d:\n", len(matches)+1)
		fmt.Printf("Skor %s: ", club1)
		fmt.Scan(&match.score1)
		fmt.Printf("Skor %s: ", club2)
		fmt.Scan(&match.score2)

		if match.score1 < 0 || match.score2 < 0 {
			break
		}

		matches = append(matches, match)

		var winner string
		if match.score1 > match.score2 {
			winner = club1
		} else if match.score2 > match.score1 {
			winner = club2
		} else {
			winner = "Draw"
		}
		winners = append(winners, winner)
	}

	fmt.Println("\nHasil pertandingan:")
	for i, match := range matches {
		fmt.Printf("Pertandingan %d: %d - %d\n", i+1, match.score1, match.score2)
	}

	fmt.Println("\nDaftar hasil:")
	for i, winner := range winners {
		fmt.Printf("Hasil %d: %s\n", i+1, winner)
	}
}
