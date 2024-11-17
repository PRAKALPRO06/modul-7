package main

import "fmt"

func main() {
	var Klub_2311102241, klubB string
	var skorA, skorB int

	fmt.Print("Klub A: ")
	fmt.Scanln(&Klub_2311102241)
	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)

	var hasil []string

	i := 1a
	for {
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scanln(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil = append(hasil, Klub_2311102241)
		} else if skorB > skorA {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		}

		i++
	}

	fmt.Println("\nHasil:")
	for i, v := range hasil {
		fmt.Printf("Hasil %d: %s\n", i+1, v)
	}

	fmt.Println("Pertandingan selesai")
}
