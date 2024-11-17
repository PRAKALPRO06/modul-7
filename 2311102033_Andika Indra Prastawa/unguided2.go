package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var clubA, clubB string
	var results []string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Klub A: ")
	scanner.Scan()
	clubA = scanner.Text()

	fmt.Print("Klub B: ")
	scanner.Scan()
	clubB = scanner.Text()

	matchCount := 1
	for {
		fmt.Printf("Pertandingan %d : ", matchCount)
		scanner.Scan()
		input := scanner.Text()
		scores := strings.Split(input, " ")

		if len(scores) != 2 {
			fmt.Println("Input tidak valid. Harus terdiri dari 2 skor.")
			continue
		}

		scoreA, err1 := strconv.Atoi(scores[0])
		scoreB, err2 := strconv.Atoi(scores[1])

		if err1 != nil || err2 != nil || scoreA < 0 || scoreB < 0 {
			fmt.Println("Skor tidak valid. Program berhenti.")
			break
		}

		if scoreA > scoreB {
			results = append(results, clubA)
		} else if scoreA < scoreB {
			results = append(results, clubB)
		} else {
			results = append(results, "Draw")
		}

		matchCount++
	}

	for i, result := range results {
		fmt.Printf("Hasil %d : %s\n", i+1, result)
	}
	fmt.Println("Pertandingan selesai")
}
