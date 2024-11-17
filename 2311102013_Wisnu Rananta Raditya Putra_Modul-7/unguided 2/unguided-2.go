//Wisnu Rananta Raditya Putra (2311102012) IF-11-06

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
	var winningClubs []string

	fmt.Println("Masukkan nama klub A dan klub B untuk memulai pertandingan!")

	// Membaca nama klub A dan klub B
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Klub A: ")
	clubA, _ = reader.ReadString('\n')
	clubA = strings.TrimSpace(clubA)

	fmt.Print("Klub B: ")
	clubB, _ = reader.ReadString('\n')
	clubB = strings.TrimSpace(clubB)

	match := 1

	fmt.Println("\nMasukkan skor pertandingan dengan format 'skorA skorB'!")

	for {
		fmt.Printf("Pertandingan %d: ", match)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		scores := strings.Split(input, " ") 

		if len(scores) != 2 {
			fmt.Println("Input tidak valid. Pertandingan selesai!")
			break
		}

		scoreA, errA := strconv.Atoi(scores[0]) 
		scoreB, errB := strconv.Atoi(scores[1]) 

		if errA != nil || errB != nil || scoreA < 0 || scoreB < 0 {
			fmt.Println("Skor tidak valid. Pertandingan selesai!")
			break
		}

		if scoreA > scoreB {
			winningClubs = append(winningClubs, fmt.Sprintf("Hasil %d : %s", match, clubA))
		} else if scoreA < scoreB {
			winningClubs = append(winningClubs, fmt.Sprintf("Hasil %d : %s", match, clubB))
		} else {
			winningClubs = append(winningClubs, fmt.Sprintf("Hasil %d : Draw", match))
		}

		match++
	}

	for _, result := range winningClubs {
		fmt.Println(result)
	}
}