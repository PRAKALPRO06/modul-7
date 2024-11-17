package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung rata-rata dari array
func hitungRataRata(arr []int) float64 {
	total := 0
	for _, v := range arr {
		total += v
	}
	return float64(total) / float64(len(arr))
}

// Fungsi untuk menghitung standar deviasi dari array
func hitungStandarDeviasi(arr []int, rataRata float64) float64 {
	total := 0.0
	for _, v := range arr {
		total += math.Pow(float64(v)-rataRata, 2)
	}
	return math.Sqrt(total / float64(len(arr)))
}

// Fungsi untuk menghapus elemen pada indeks tertentu
func hapusElemen(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

// Fungsi untuk menghitung frekuensi bilangan dalam array
func hitungFrekuensi(arr []int, bilangan int) int {
	count := 0
	for _, v := range arr {
		if v == bilangan {
			count++
		}
	}
	return count
}

func main() {
	// Input jumlah elemen array
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	// Membuat array dengan kapasitas tertentu
	arr := make([]int, n)

	// Mengisi elemen array
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	// a. Menampilkan keseluruhan isi dari array
	fmt.Println("\nKeseluruhan isi array:", arr)

	// b. Menampilkan elemen-elemen array dengan indeks ganjil
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// c. Menampilkan elemen-elemen array dengan indeks genap
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// d. Menampilkan elemen dengan indeks kelipatan bilangan x
	var x int
	fmt.Print("\nMasukkan bilangan x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
	for i := 0; i < len(arr); i += x {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// e. Menghapus elemen array pada indeks tertentu
	var index int
	fmt.Print("\nMasukkan indeks untuk menghapus elemen: ")
	fmt.Scan(&index)
	if index >= 0 && index < len(arr) {
		arr = hapusElemen(arr, index)
		fmt.Println("Isi array setelah elemen dihapus:", arr)
	} else {
		fmt.Println("Indeks tidak valid!")
	}

	// f. Menampilkan rata-rata dari bilangan yang ada di dalam array
	rataRata := hitungRataRata(arr)
	fmt.Printf("\nRata-rata dari array: %.2f\n", rataRata)

	// g. Menampilkan standar deviasi dari bilangan yang ada di dalam array
	stdDeviasi := hitungStandarDeviasi(arr, rataRata)
	fmt.Printf("Standar deviasi dari array: %.2f\n", stdDeviasi)

	// h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array
	var bilangan int
	fmt.Print("\nMasukkan bilangan untuk mencari frekuensinya: ")
	fmt.Scan(&bilangan)
	frekuensi := hitungFrekuensi(arr, bilangan)
	fmt.Printf("Frekuensi bilangan %d di dalam array: %d\n", bilangan, frekuensi)
}
