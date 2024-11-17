package main

import "fmt"

// struct untuk menyimpan informasi klub
type Klub struct {
	nama string
	skor int
}

func main() {
	var klubA, klubB Klub
	var pemenang []string

	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA.nama)
	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB.nama)

	for {
		fmt.Printf("\nMasukkan skor %s: ", klubA.nama)
		fmt.Scan(&klubA.skor)
		fmt.Printf("Masukkan skor %s: ", klubB.nama)
		fmt.Scan(&klubB.skor)

		// Hentikan input jika salah satu atau kedua skor tidak valid (negatif)
		if klubA.skor < 0 || klubB.skor < 0 {
			fmt.Println("\nSkor tidak valid. Proses input dihentikan.")
			break
		}

		// Tentukan pemenang dan simpan nama klub yang menang
		if klubA.skor > klubB.skor {
			fmt.Printf("%s memenangkan pertandingan\n", klubA.nama)
			pemenang = append(pemenang, klubA.nama)
		} else if klubB.skor > klubA.skor {
			fmt.Printf("%s memenangkan pertandingan\n", klubB.nama)
			pemenang = append(pemenang, klubB.nama)
		} else {
			fmt.Println("Pertandingan berakhir seri")
		}
	}

	// Tampilkan daftar klub pemenang
	fmt.Println("\nDaftar klub yang memenangkan pertandingan:")
	for i, klub := range pemenang {
		fmt.Printf("Pemenang %d: %s\n", i+1, klub)
	}
}
