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

// Fungsi untuk menampilkan elemen array berdasarkan filter
func displayFilteredElements(arr []int, filterFunc func(int) bool) {
	for i, val := range arr {
		if filterFunc(i) {
			fmt.Printf("%d ", val)
		}
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&n)

	// Validasi jumlah elemen array
	if n <= 0 {
		fmt.Println("Jumlah elemen harus lebih dari 0.")
		return
	}

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
	displayFilteredElements(arr, func(i int) bool { return i%2 != 0 })

	// c. Menampilkan elemen array dengan indeks genap
	fmt.Println("\nElemen array dengan indeks genap:")
	displayFilteredElements(arr, func(i int) bool { return i%2 == 0 })

	// d. Menampilkan elemen array dengan indeks kelipatan bilangan x
	var x int
	fmt.Print("\nMasukkan bilangan x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	if x > 0 {
		fmt.Printf("Elemen array dengan indeks kelipatan %d:\n", x)
		displayFilteredElements(arr, func(i int) bool { return i%x == 0 })
	} else {
		fmt.Println("Bilangan x harus lebih dari 0.")
	}

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
	if len(arr) > 0 {
		average := calculateAverage(arr)
		fmt.Printf("\nRata-rata bilangan dalam array: %.2f\n", average)

		// g. Menampilkan standar deviasi dari bilangan dalam array
		stdDev := calculateStdDev(arr, average)
		fmt.Printf("Standar deviasi bilangan dalam array: %.2f\n", stdDev)
	} else {
		fmt.Println("\nArray kosong setelah penghapusan, rata-rata dan standar deviasi tidak dapat dihitung.")
	}

	// h. Menampilkan frekuensi dari suatu bilangan tertentu
	var target int
	fmt.Print("\nMasukkan bilangan yang ingin dihitung frekuensinya: ")
	fmt.Scan(&target)
	frequency := calculateFrequency(arr, target)
	fmt.Printf("Frekuensi bilangan %d dalam array: %d\n", target, frequency)
}
