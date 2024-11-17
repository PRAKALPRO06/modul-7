package main

import "fmt"

type Match struct {
	clubA  string
	clubB  string
	scoreA int
	scoreB int
	winner string
}

func main() {
	var clubA, clubB string
	fmt.Print("Klub A: ")
	fmt.Scan(&clubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&clubB)

	var matches []Match

	for i := 1; ; i++ {
		var scoreA, scoreB int
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scan(&scoreA, &scoreB)

		if scoreA < 0 || scoreB < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}

		var winner string
		if scoreA > scoreB {
			winner = clubA
		} else if scoreA < scoreB {
			winner = clubB
		} else {
			winner = "Draw"
		}

		match := Match{
			clubA:  clubA,
			clubB:  clubB,
			scoreA: scoreA,
			scoreB: scoreB,
			winner: winner,
		}
		matches = append(matches, match)
	}

	fmt.Println("\nHasil pertandingan:")
	for i, match := range matches {
		fmt.Printf("Hasil %d: %s\n", i+1, match.winner)
	}
}
