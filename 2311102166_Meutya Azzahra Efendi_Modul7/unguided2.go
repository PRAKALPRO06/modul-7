// Meutya Azzahra Efendi
// 2311102166

package main

import "fmt"

func inputNamaKlub() (string, string) {
	var klubA, klubB string
	fmt.Print("Masukkan nama klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Masukkan nama klub B: ")
	fmt.Scanln(&klubB)
	return klubA, klubB
}

func inputSkor(pertandingan int, klubA, klubB string) (int, int) {
	var skorA, skorB int
	fmt.Printf("Pertandingan %d antara %s dan %s: ", pertandingan, klubA, klubB)
	_, err := fmt.Scanf("%d %d\n", &skorA, &skorB)

	// cek jika input tidak valid
	if err != nil {
		fmt.Println("Input tidak valid. Pastikan untuk memasukkan dua angka.")
		// mengosongkan input buffer
		var dummy string
		fmt.Scanln(&dummy)
		return -1, -1 // mengembalikan nilai tidak valid
	}
	return skorA, skorB
}

func main() {
	klubA, klubB := inputNamaKlub()
	var hasil []string

	// loop untuk input skor
	for pertandingan := 1; pertandingan <= 9; pertandingan++ {
		skorA, skorB := inputSkor(pertandingan, klubA, klubB)

		// break loop jika salah satu skor negatif
		if skorA < 0 || skorB < 0 {
			break
		}

		// tentukan hasil pertandingan
		if skorA > skorB {
			hasil = append(hasil, klubA) // klubA menang
		} else if skorA < skorB {
			hasil = append(hasil, klubB) // klubB menang
		} else {
			hasil = append(hasil, "Draw") // seri
		}
	}

	// tampilkan hasil
	fmt.Println("Hasil pertandingan:")
	for i := 0; i < 9; i++ {
		if i < len(hasil) {
			fmt.Printf("Hasil Pertandingan %d: %s\n", i+1, hasil[i])
		} else {
			// jika belum ada hasil untuk pertandingan tersebut
			fmt.Printf("Hasil Pertandingan %d: Belum selesai\n", i+1)
		}
	}
	fmt.Println("Pertandingan selesai")
}