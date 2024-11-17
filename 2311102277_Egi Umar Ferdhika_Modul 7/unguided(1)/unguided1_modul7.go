package main

import (
	"fmt"
	"math"
	// Nama : Egi Umar Ferdhika
	// NIM 	: 2311102277
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	arr := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tampilkan semua isi array")
		fmt.Println("2. Tampilkan elemen dengan indeks ganjil")
		fmt.Println("3. Tampilkan elemen dengan indeks genap")
		fmt.Println("4. Tampilkan elemen dengan indeks kelipatan x")
		fmt.Println("5. Hapus elemen array pada indeks tertentu")
		fmt.Println("6. Tampilkan rata-rata")
		fmt.Println("7. Tampilkan standar deviasi")
		fmt.Println("8. Tampilkan frekuensi bilangan")
		fmt.Println("9. Keluar")

		var choice int
		fmt.Print("Pilih menu: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Isi array:", arr)
		case 2:
			for i := 1; i < len(arr); i += 2 {
				fmt.Printf("%v ", arr[i])
			}
			fmt.Println()
		case 3:
			for i := 0; i < len(arr); i += 2 {
				fmt.Printf("%v ", arr[i])
			}
			fmt.Println()
		case 4:
			var x int
			fmt.Print("Masukkan nilai x: ")
			fmt.Scan(&x)
			for i := 0; i < len(arr); i++ {
				if i%x == 0 {
					fmt.Printf("%v ", arr[i])
				}
			}
			fmt.Println()
		case 5:
			var idx int
			fmt.Print("Masukkan indeks yang akan dihapus: ")
			fmt.Scan(&idx)
			if idx >= 0 && idx < len(arr) {
				arr = append(arr[:idx], arr[idx+1:]...)
			}
		case 6:
			sum := 0.0
			for _, v := range arr {
				sum += v
			}
			fmt.Printf("Rata-rata: %.2f\n", sum/float64(len(arr)))
		case 7:
			mean := 0.0
			for _, v := range arr {
				mean += v
			}
			mean /= float64(len(arr))

			variance := 0.0
			for _, v := range arr {
				variance += math.Pow(v-mean, 2)
			}
			variance /= float64(len(arr))
			stdDev := math.Sqrt(variance)
			fmt.Printf("Standar deviasi: %.2f\n", stdDev)
		case 8:
			freq := make(map[float64]int)
			for _, v := range arr {
				freq[v]++
			}
			for num, count := range freq {
				fmt.Printf("Bilangan %.2f muncul %d kali\n", num, count)
			}
		case 9:
			return
		}
	}
}
