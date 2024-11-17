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
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scanln(&array[i])
	}

	for {
	
		fmt.Println("\nPilih operasi:")
		fmt.Println("a. Menampilkan keseluruhan isi array")
		fmt.Println("b. Menampilkan elemen-elemen array dengan indeks ganjil")
		fmt.Println("c. Menampilkan elemen-elemen array dengan indeks genap")
		fmt.Println("d. Menampilkan elemen array dengan indeks kelipatan bilangan tertentu")
		fmt.Println("e. Menghapus elemen array pada indeks tertentu")
		fmt.Println("f. Menampilkan rata-rata dari bilangan di dalam array")
		fmt.Println("g. Menampilkan standar deviasi dari bilangan di dalam array")
		fmt.Println("h. Menampilkan frekuensi dari suatu bilangan tertentu")
		fmt.Println("q. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "a":
	
			fmt.Println("Isi array:", array)

		case "b":
			
			fmt.Println("Elemen dengan indeks ganjil:")
			for i := 1; i < len(array); i += 2 {
				fmt.Printf("array[%d] = %d\n", i, array[i])
			}

		case "c":
		
			fmt.Println("Elemen dengan indeks genap:")
			for i := 0; i < len(array); i += 2 {
				fmt.Printf("array[%d] = %d\n", i, array[i])
			}

		case "d":
		
			var x int
			fmt.Print("Masukkan bilangan x: ")
			fmt.Scanln(&x)
			fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
			for i := 0; i < len(array); i++ {
				if i%x == 0 {
					fmt.Printf("array[%d] = %d\n", i, array[i])
				}
			}

		case "e":
	
			var idx int
			fmt.Print("Masukkan indeks yang akan dihapus: ")
			fmt.Scanln(&idx)
			if idx >= 0 && idx < len(array) {
				array = append(array[:idx], array[idx+1:]...)
				fmt.Println("Array setelah penghapusan:", array)
			} else {
				fmt.Println("Indeks tidak valid!")
			}

		case "f":
			
			sum := 0
			for _, val := range array {
				sum += val
			}
			average := float64(sum) / float64(len(array))
			fmt.Printf("Rata-rata: %.2f\n", average)

		case "g":
			sum := 0
			for _, val := range array {
				sum += val
			}
			mean := float64(sum) / float64(len(array))
			var varianceSum float64
			for _, val := range array {
				varianceSum += math.Pow(float64(val)-mean, 2)
			}
			stdDev := math.Sqrt(varianceSum / float64(len(array)))
			fmt.Printf("Standar deviasi: %.2f\n", stdDev)

		case "h":
			
			var target int
			fmt.Print("Masukkan bilangan yang dicari frekuensinya: ")2
			fmt.Scanln(&target)
			count := 0
			for _, val := range array {
				if val == target {
					count++
				}
			}
			fmt.Printf("Frekuensi bilangan %d: %d kali\n", target, count)

		case "q":
		
			fmt.Println("Keluar dari program.")
			return

		default:
			fmt.Println("Pilihan tidak valid! Silakan coba lagi.")
		}
	}
}
