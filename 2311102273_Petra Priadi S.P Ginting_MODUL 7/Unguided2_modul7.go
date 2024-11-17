package main

import "fmt"

func main() {

	var klub1, klub2 string

	var totalhasilPertandingan []string

	fmt.Print("Nama Klub A : ")
	fmt.Scan(&klub1)
	fmt.Print("Nama Klub B : ")
	fmt.Scan(&klub2)

	var skor1, skor2 int
	pertandinganKe := 1

	for {
		fmt.Printf("Pertandingan %d : ", pertandinganKe)
		fmt.Scan(&skor1, &skor2)

		if skor1 < 0 || skor2 < 0 {
			break
		}

		if skor1 > skor2 {
			totalhasilPertandingan = append(totalhasilPertandingan, klub1)
		} else if skor2 > skor1 {
			totalhasilPertandingan = append(totalhasilPertandingan, klub2)
		} else {
			totalhasilPertandingan = append(totalhasilPertandingan, "Draw")
		}

		pertandinganKe++
	}

	fmt.Println("\nHasil pertandingan Akhir:")
	for i, hasil := range totalhasilPertandingan {
		fmt.Printf("Hasil %d : %s\n", i+1, hasil)
	}

	fmt.Println("Pertandingan selesai")
}
