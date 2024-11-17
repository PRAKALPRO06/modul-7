package main
import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var yangMenang []string

	fmt.Print("Klub A : ")
	fmt.Scanln(&klubA)

	fmt.Print("Klub B : ")
	fmt.Scanln(&klubB)

	fmt.Println()

	pertandingan := 1
	for {
		fmt.Printf("Pertandingan %d : ", pertandingan)
		_, err := fmt.Scan(&skorA, &skorB)

		if err != nil {
			fmt.Println("Cara inputnya salah, input 2 nomer sekaligus!")
			continue
		}

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			yangMenang = append(yangMenang, klubA)
		} else if skorA < skorB {
			yangMenang = append(yangMenang, klubB)
		} else {
			yangMenang = append(yangMenang, "Draw")
		}

		pertandingan++
	}

	fmt.Println()
	for i, hasil := range yangMenang {
		fmt.Printf("Hasil %d : %s\n", i+1, hasil)
	}
	fmt.Println("Pertandingan selesai")
}