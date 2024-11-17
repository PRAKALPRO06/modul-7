package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Println("Masukkan jumlah elemen array:")
	fmt.Scan(&n)

	// Input array
	array := make([]int, n)
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	for {
		fmt.Println("\nPilih operasi:")
		fmt.Println("a. Tampilkan keseluruhan isi array")
		fmt.Println("b. Tampilkan elemen dengan indeks ganjil")
		fmt.Println("c. Tampilkan elemen dengan indeks genap")
		fmt.Println("d. Tampilkan elemen dengan indeks kelipatan x")
		fmt.Println("e. Hapus elemen array pada indeks tertentu")
		fmt.Println("f. Tampilkan rata-rata array")
		fmt.Println("g. Tampilkan standar deviasi array")
		fmt.Println("h. Tampilkan frekuensi bilangan tertentu")
		fmt.Println("i. Keluar")

		var choice string
		fmt.Scan(&choice)

		switch choice {
		case "a":
			fmt.Println("Isi array:", array)

		case "b":
			fmt.Print("Elemen dengan indeks ganjil: ")
			for i := 1; i < len(array); i += 2 {
				fmt.Print(array[i], " ")
			}
			fmt.Println()

		case "c":
			fmt.Print("Elemen dengan indeks genap: ")
			for i := 0; i < len(array); i += 2 {
				fmt.Print(array[i], " ")
			}
			fmt.Println()

		case "d":
			var x int
			fmt.Println("Masukkan nilai x:")
			fmt.Scan(&x)
			fmt.Print("Elemen dengan indeks kelipatan ", x, ": ")
			for i := 0; i < len(array); i++ {
				if i%x == 0 {
					fmt.Print(array[i], " ")
				}
			}
			fmt.Println()

		case "e":
			var index int
			fmt.Println("Masukkan indeks elemen yang ingin dihapus:")
			fmt.Scan(&index)
			if index >= 0 && index < len(array) {
				array = append(array[:index], array[index+1:]...)
				fmt.Println("Array setelah penghapusan:", array)
			} else {
				fmt.Println("Indeks tidak valid.")
			}

		case "f":
			sum := 0
			for _, val := range array {
				sum += val
			}
			rataRata := float64(sum) / float64(len(array))
			fmt.Println("Rata-rata:", rataRata)

		case "g":
			sum := 0
			for _, val := range array {
				sum += val
			}
			rataRata := float64(sum) / float64(len(array))
			var variance float64
			for _, val := range array {
				variance += math.Pow(float64(val)-rataRata, 2)
			}
			stdDeviasi := math.Sqrt(variance / float64(len(array)))
			fmt.Println("Standar deviasi:", stdDeviasi)

		case "h":
			var target int
			fmt.Println("Masukkan bilangan yang ingin dicari frekuensinya:")
			fmt.Scan(&target)
			frequency := 0
			for _, val := range array {
				if val == target {
					frequency++
				}
			}
			fmt.Println("Frekuensi bilangan", target, ":", frequency)

		case "i":
			fmt.Println("Program selesai.")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}