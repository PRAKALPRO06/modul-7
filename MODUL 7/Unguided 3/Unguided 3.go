package main

import "fmt"

func main() {
	var klub1, klub2 string
	var skor1, skor2 int
	var hasil []string

	fmt.Print("Masukkan nama klub pertama: ")
	fmt.Scan(&klub1)
	fmt.Print("Masukkan nama klub kedua: ")
	fmt.Scan(&klub2)

	for i := 1; ; i++ {
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scan(&skor1, &skor2)

		if skor1 < 0 || skor2 < 0 {
			break
		}

		var pemenang string
		if skor1 > skor2 {
			pemenang = klub1
		} else if skor2 > skor1 {
			pemenang = klub2
		} else {
			pemenang = "Draw"
		}

		hasil = append(hasil, fmt.Sprintf("Hasil %d: %s", i, pemenang))
	}

	for _, h := range hasil {
		fmt.Println(h)
	}

	fmt.Println("Pertandingan selesai.")
}