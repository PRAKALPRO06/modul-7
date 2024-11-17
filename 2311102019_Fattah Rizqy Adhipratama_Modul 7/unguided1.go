package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scanln(&n)

	array := make([]int, n)

	// Input elemen array
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scanln(&array[i])
	}

	for {
		fmt.Println("\nPilih operasi yang ingin dilakukan:")
		fmt.Println("a. Tampilkan keseluruhan isi array.")
		fmt.Println("b. Tampilkan elemen array dengan indeks ganjil.")
		fmt.Println("c. Tampilkan elemen array dengan indeks genap.")
		fmt.Println("d. Tampilkan elemen array dengan indeks kelipatan x.")
		fmt.Println("e. Hapus elemen array pada indeks tertentu.")
		fmt.Println("f. Hitung rata-rata elemen array.")
		fmt.Println("g. Hitung standar deviasi elemen array.")
		fmt.Println("h. Hitung frekuensi suatu bilangan dalam array.")
		fmt.Println("i. Keluar.")
		fmt.Print("Pilihan Anda: ")

		var pilihan string
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "a":
			fmt.Println("Isi array:", array)

		case "b":
			fmt.Println("Elemen array dengan indeks ganjil:")
			for i := 1; i < len(array); i += 2 {
				fmt.Printf("%d ", array[i])
			}
			fmt.Println()

		case "c":
			fmt.Println("Elemen array dengan indeks genap:")
			for i := 0; i < len(array); i += 2 {
				fmt.Printf("%d ", array[i])
			}
			fmt.Println()

		case "d":
			var x int
			fmt.Print("Masukkan nilai x: ")
			fmt.Scanln(&x)
			fmt.Printf("Elemen array dengan indeks kelipatan %d:\n", x)
			for i := 0; i < len(array); i++ {
				if i%x == 0 {
					fmt.Printf("%d ", array[i])
				}
			}
			fmt.Println()

		case "e":
			var index int
			fmt.Print("Masukkan indeks elemen yang ingin dihapus: ")
			fmt.Scanln(&index)
			if index >= 0 && index < len(array) {
				array = append(array[:index], array[index+1:]...)
				fmt.Println("Array setelah penghapusan:", array)
			} else {
				fmt.Println("Indeks tidak valid!")
			}

		case "f":
			total := 0
			for _, value := range array {
				total += value
			}
			rata := float64(total) / float64(len(array))
			fmt.Printf("Rata-rata elemen array: %.2f\n", rata)

		case "g":
			// Hitung rata-rata terlebih dahulu
			total := 0
			for _, value := range array {
				total += value
			}
			rata := float64(total) / float64(len(array))

			// Hitung variansi
			var variansi float64
			for _, value := range array {
				variansi += math.Pow(float64(value)-rata, 2)
			}
			variansi /= float64(len(array))

			// Standar deviasi
			stdDeviasi := math.Sqrt(variansi)
			fmt.Printf("Standar deviasi elemen array: %.2f\n", stdDeviasi)

		case "h":
			var target int
			fmt.Print("Masukkan bilangan yang ingin dihitung frekuensinya: ")
			fmt.Scanln(&target)
			frekuensi := 0
			for _, value := range array {
				if value == target {
					frekuensi++
				}
			}
			fmt.Printf("Frekuensi bilangan %d dalam array: %d\n", target, frekuensi)

		case "i":
			fmt.Println("Keluar dari program.")
			return

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}