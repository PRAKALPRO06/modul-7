package main

import "fmt"

type Pertandingan struct {
	klubA    string
	klubB    string
	skorA    int
	skorB    int
	pemenang string
}

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil []string
	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	for n := 1; ; n++ {
		fmt.Printf("Pertandingan %d : ", n)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		var pemenang string
		if skorA > skorB {
			pemenang = klubA
		} else if skorA < skorB {
			pemenang = klubB
		} else {
			pemenang = "Draw"
		}

		hasil = append(hasil, pemenang)
	}

	fmt.Println("\nDaftar klub yang memenangkan pertandingan:")
	for n, h := range hasil {
		fmt.Printf("Hasil %d : %s\n", n+1, h)
	}

	fmt.Println("Pertandingan selesai")
}
