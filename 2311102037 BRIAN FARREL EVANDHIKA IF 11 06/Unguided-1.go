//2311102037_BRIAN FARREL EVANDHIKA_IF 11 06
package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung rata-rata array
func calculateAverage(arr []int) float64 {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return float64(sum) / float64(len(arr))
}

// Fungsi untuk menghitung standar deviasi
func calculateStdDev(arr []int, mean float64) float64 {
	sum := 0.0
	for _, val := range arr {
		sum += math.Pow(float64(val)-mean, 2)
	}
	return math.Sqrt(sum / float64(len(arr)))
}

// Fungsi untuk menghitung frekuensi bilangan tertentu dalam array
func calculateFrequency(arr []int, target int) int {
	count := 0
	for _, val := range arr {
		if val == target {
			count++
		}
	}
	return count
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&n)

	// Inisialisasi array
	arr := make([]int, n)

	// Input elemen array
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	// a. Menampilkan keseluruhan isi array
	fmt.Println("\nKeseluruhan isi array:")
	fmt.Println(arr)

	// b. Menampilkan elemen array dengan indeks ganjil
	fmt.Println("\nElemen array dengan indeks ganjil:")
	for i := 1; i < len(arr); i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	// c. Menampilkan elemen array dengan indeks genap
	fmt.Println("\nElemen array dengan indeks genap:")
	for i := 0; i < len(arr); i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	// d. Menampilkan elemen array dengan indeks kelipatan bilangan x
	var x int
	fmt.Print("\nMasukkan bilangan x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Println("Elemen array dengan indeks kelipatan", x, ":")
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Printf("%d ", arr[i])
		}
	}
	fmt.Println()

	// e. Menghapus elemen array pada indeks tertentu
	var index int
	fmt.Print("\nMasukkan indeks yang ingin dihapus: ")
	fmt.Scan(&index)
	if index >= 0 && index < len(arr) {
		arr = append(arr[:index], arr[index+1:]...)
		fmt.Println("Array setelah penghapusan:")
		fmt.Println(arr)
	} else {
		fmt.Println("Indeks tidak valid!")
	}

	// f. Menampilkan rata-rata dari bilangan dalam array
	average := calculateAverage(arr)
	fmt.Printf("\nRata-rata bilangan dalam array: %.2f\n", average)

	// g. Menampilkan standar deviasi dari bilangan dalam array
	stdDev := calculateStdDev(arr, average)
	fmt.Printf("Standar deviasi bilangan dalam array: %.2f\n", stdDev)

	// h. Menampilkan frekuensi dari suatu bilangan tertentu
	var target int
	fmt.Print("\nMasukkan bilangan yang ingin dihitung frekuensinya: ")
	fmt.Scan(&target)
	frequency := calculateFrequency(arr, target)
	fmt.Printf("Frekuensi bilangan %d dalam array: %d\n", target, frequency)
}
