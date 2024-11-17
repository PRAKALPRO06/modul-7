package main

import (
	"fmt"
	"math"
)

func main() {
	var erwin int
	fmt.Println("Masukkan Jumlah Elemen Dalam Array: ")
	fmt.Scanln(&erwin)

	numbers := make([]float64, erwin)

	for i := 0; i < erwin; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&numbers[i])
	}

	for {
		fmt.Println("\nMenu")
		fmt.Println("1. Tampilkan Seluruh Isi Elemen Array")
		fmt.Println("2. Tampilkan Elemen Indeks Ganjil")
		fmt.Println("3. Tampilkan Elemen Indeks Genap")
		fmt.Println("4. Menampilkan Elemen Array Berdasarkan Kelipatan Indeks")
		fmt.Println("5. Menghapus Elemen Array Pada Indeks Tertentu")
		fmt.Println("6. Menampilkan Rata-rata")
		fmt.Println("7. Tampilkan Standar Deviasi")
		fmt.Println("8. Menampilkan Frekuensi")
		fmt.Println("9. Keluar")

		var pilihan int
		fmt.Println("Masukkan Pilihan 1-9: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			for i, num := range numbers {
				fmt.Printf("Index %d: %.2f\n", i, num)
			}

		case 2:
			fmt.Println("Elemen pada indeks ganjil:")
			for i, num := range numbers {
				if i%2 != 0 {
					fmt.Printf("Index %d: %.2f\n", i, num)
				}
			}

		case 3:
			fmt.Println("Elemen pada indeks genap:")
			for i, num := range numbers {
				if i%2 == 0 {
					fmt.Printf("Index %d: %.2f\n", i, num)
				}
			}

		case 4:
			var x int
			fmt.Print("Masukkan Nilai X: ")
			fmt.Scan(&x)
			if x > 0 {
				fmt.Printf("\nElemen dengan indeks kelipatan %d:\n", x)
				for i, num := range numbers {
					if i%x == 0 {
						fmt.Printf("Index %d: %.2f\n", i, num)
					}
				}
			} else {
				fmt.Println("Nilai X harus lebih besar dari 0.")
			}

		case 5:
			var idx int
			fmt.Print("Masukkan indeks yang akan dihapus: ")
			fmt.Scan(&idx)
			if idx >= 0 && idx < len(numbers) {
				numbers = append(numbers[:idx], numbers[idx+1:]...)
				fmt.Printf("Elemen pada indeks %d telah dihapus\n", idx)
			} else {
				fmt.Println("Indeks tidak valid")
			}

		case 6:
			sum := 0.0
			count := len(numbers)
			for _, num := range numbers {
				sum += num
			}
			if count > 0 {
				fmt.Printf("Rata-rata: %.2f\n", sum/float64(count))
			} else {
				fmt.Println("Array kosong")
			}

		case 7:
			var mean, sumSquares float64
			count := len(numbers)
			for _, num := range numbers {
				mean += num
			}
			if count > 0 {
				mean /= float64(count)
				for _, num := range numbers {
					sumSquares += math.Pow(num-mean, 2)
				}
				stdDev := math.Sqrt(sumSquares / float64(count))
				fmt.Printf("Standar Deviasi: %.2f\n", stdDev)
			} else {
				fmt.Println("Array kosong")
			}

		case 8:
			var target float64
			fmt.Print("Masukkan bilangan yang ingin dihitung frekuensinya: ")
			fmt.Scan(&target)

			freq := 0
			for _, num := range numbers {
				if num == target {
					freq++
				}
			}
			fmt.Printf("Frekuensi %.2f dalam array: %d\n", target, freq)

		case 9:
			fmt.Println("Program selesai")
			return

		default:
			fmt.Println("Pilihan Tidak Tersedia")
		}
	}
}
