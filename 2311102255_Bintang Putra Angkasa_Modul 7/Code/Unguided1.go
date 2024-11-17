package main

import (
	"fmt"
	"math"
)

func isiArray(array []int) {
	for i := 0; i < len(array); i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i+1)
		fmt.Scan(&array[i])
	}
}

func tampilkanArray(array []int) {
	fmt.Println("\nIsi dari array:")
	for i, val := range array {
		fmt.Printf("Indeks %d: %d\n", i, val)
	}
}

func tampilkanIndeksGanjil(array []int) {
	fmt.Println("\nElemen dengan indeks ganjil:")
	for i := 1; i < len(array); i += 2 {
		fmt.Printf("Indeks %d: %d\n", i, array[i])
	}
}

func tampilkanIndeksGenap(array []int) {
	fmt.Println("\nElemen dengan indeks genap:")
	for i := 0; i < len(array); i += 2 {
		fmt.Printf("Indeks %d: %d\n", i, array[i])
	}
}

func tampilkanIndeksKelipatanX(array []int, x int) {
	fmt.Printf("\nElemen dengan indeks kelipatan %d:\n", x)
	for i := 0; i < len(array); i++ {
		if i%x == 0 {
			fmt.Printf("Indeks %d: %d\n", i, array[i])
		}
	}
}

func hitungRataRata(array []int) float64 {
	total := 0
	for _, val := range array {
		total += val
	}
	return float64(total) / float64(len(array))
}

func hitungStandarDeviasi(array []int) float64 {
	rataRata := hitungRataRata(array)
	total := 0.0
	for _, val := range array {
		total += math.Pow(float64(val)-rataRata, 2)
	}
	return math.Sqrt(total / float64(len(array)))
}

func hitungFrekuensi(array []int, bilangan int) int {
	frekuensi := 0
	for _, val := range array {
		if val == bilangan {
			frekuensi++
		}
	}
	return frekuensi
}

func hapusElemen(array *[]int, indeks int) {
	*array = append((*array)[:indeks], (*array)[indeks+1:]...)
}

func main() {
	var kapasitas int
	fmt.Print("Masukkan kapasitas array: ")
	fmt.Scan(&kapasitas)

	array := make([]int, kapasitas)

	isiArray(array)

	for {
		var pilihan, x, indeks, bilangan int
		fmt.Println("\nMenu:")
		fmt.Println("1. Tampilkan seluruh isi array")
		fmt.Println("2. Tampilkan elemen dengan indeks ganjil")
		fmt.Println("3. Tampilkan elemen dengan indeks genap")
		fmt.Println("4. Tampilkan elemen dengan indeks kelipatan x")
		fmt.Println("5. Hapus elemen pada indeks tertentu")
		fmt.Println("6. Hitung rata-rata elemen")
		fmt.Println("7. Hitung standar deviasi elemen")
		fmt.Println("8. Hitung frekuensi suatu bilangan")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih opsi (1-9): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilkanArray(array)
		case 2:
			tampilkanIndeksGanjil(array)
		case 3:
			tampilkanIndeksGenap(array)
		case 4:
			fmt.Print("Masukkan bilangan x: ")
			fmt.Scan(&x)
			tampilkanIndeksKelipatanX(array, x)
		case 5:
			fmt.Print("Masukkan indeks yang ingin dihapus: ")
			fmt.Scan(&indeks)
			if indeks >= 0 && indeks < len(array) {
				hapusElemen(&array, indeks)
				fmt.Println("Elemen berhasil dihapus.")
				tampilkanArray(array)
			} else {
				fmt.Println("Indeks tidak valid.")
			}
		case 6:
			fmt.Printf("Rata-rata elemen: %.2f\n", hitungRataRata(array))
		case 7:
			fmt.Printf("Standar deviasi elemen: %.2f\n", hitungStandarDeviasi(array))
		case 8:
			fmt.Print("Masukkan bilangan yang ingin dihitung frekuensinya: ")
			fmt.Scan(&bilangan)
			frekuensi := hitungFrekuensi(array, bilangan)
			fmt.Printf("Frekuensi bilangan %d: %d\n", bilangan, frekuensi)
		case 9:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
