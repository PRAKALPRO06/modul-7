package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, deleteIndex, targetNumber int
	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&n)

	array := make([]int, n)
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tampilkan keseluruhan isi array")
		fmt.Println("2. Tampilkan elemen dengan indeks ganjil")
		fmt.Println("3. Tampilkan elemen dengan indeks genap")
		fmt.Println("4. Tampilkan elemen dengan indeks kelipatan X")
		fmt.Println("5. Hapus elemen pada indeks tertentu")
		fmt.Println("6. Tampilkan rata-rata elemen array")
		fmt.Println("7. Tampilkan standar deviasi elemen array")
		fmt.Println("8. Tampilkan frekuensi dari bilangan tertentu")
		fmt.Println("9. Keluar")

		var choice int
		fmt.Print("Pilih menu: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Isi array:", array)
		case 2:
			fmt.Print("Elemen dengan indeks ganjil: ")
			for i := 1; i < len(array); i += 2 {
				fmt.Print(array[i], " ")
			}
			fmt.Println()
		case 3:
			fmt.Print("Elemen dengan indeks genap: ")
			for i := 0; i < len(array); i += 2 {
				fmt.Print(array[i], " ")
			}
			fmt.Println()
		case 4:
			fmt.Print("Masukkan nilai X: ")
			fmt.Scan(&x)
			fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
			for i := 0; i < len(array); i++ {
				if i%x == 0 {
					fmt.Print(array[i], " ")
				}
			}
			fmt.Println()
		case 5:
			fmt.Print("Masukkan indeks elemen yang akan dihapus: ")
			fmt.Scan(&deleteIndex)
			if deleteIndex >= 0 && deleteIndex < len(array) {
				array = append(array[:deleteIndex], array[deleteIndex+1:]...)
				fmt.Println("Array setelah penghapusan:", array)
			} else {
				fmt.Println("Indeks tidak valid!")
			}
		case 6:
			sum := 0
			for _, value := range array {
				sum += value
			}
			average := float64(sum) / float64(len(array))
			fmt.Printf("Rata-rata elemen array: %.2f\n", average)
		case 7:
			sum := 0
			for _, value := range array {
				sum += value
			}
			average := float64(sum) / float64(len(array))
			variance := 0.0
			for _, value := range array {
				variance += math.Pow(float64(value)-average, 2)
			}
			stdDev := math.Sqrt(variance / float64(len(array)))
			fmt.Printf("Standar deviasi elemen array: %.2f\n", stdDev)
		case 8:
			fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
			fmt.Scan(&targetNumber)
			frequency := 0
			for _, value := range array {
				if value == targetNumber {
					frequency++
				}
			}
			fmt.Printf("Frekuensi bilangan %d: %d kali\n", targetNumber, frequency)
		case 9:
			fmt.Println("Keluar dari program...")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi!")
		}
	}
}
