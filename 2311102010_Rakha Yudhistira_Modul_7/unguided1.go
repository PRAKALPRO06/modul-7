//2311102010_RAKHA YUDHISTIRA_IF-11-06

package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung rata-rata array
func hitungRataRata(arr []int) float64 {
	jumlah := 0
	for _, nilai := range arr {
		jumlah += nilai
	}
	return float64(jumlah) / float64(len(arr))
}

// Fungsi untuk menghitung standar deviasi
func hitungStandarDeviasi(arr []int, rataRata float64) float64 {
	jumlah := 0.0
	for _, nilai := range arr {
		jumlah += math.Pow(float64(nilai)-rataRata, 2)
	}
	return math.Sqrt(jumlah / float64(len(arr)))
}

// Fungsi untuk menghitung frekuensi bilangan tertentu dalam array
func hitungFrekuensi(arr []int, target int) int {
	jumlah := 0
	for _, nilai := range arr {
		if nilai == target {
			jumlah++
		}
	}
	return jumlah
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
	var indeks int
	fmt.Print("\nMasukkan indeks yang ingin dihapus: ")
	fmt.Scan(&indeks)
	if indeks >= 0 && indeks < len(arr) {
		arr = append(arr[:indeks], arr[indeks+1:]...)
		fmt.Println("Array setelah penghapusan:")
		fmt.Println(arr)
	} else {
		fmt.Println("Indeks tidak valid!")
	}

	// f. Menampilkan rata-rata dari bilangan dalam array
	rataRata := hitungRataRata(arr)
	fmt.Printf("\nRata-rata bilangan dalam array: %.2f\n", rataRata)

	// g. Menampilkan standar deviasi dari bilangan dalam array
	standarDeviasi := hitungStandarDeviasi(arr, rataRata)
	fmt.Printf("Standar deviasi bilangan dalam array: %.2f\n", standarDeviasi)

	// h. Menampilkan frekuensi dari suatu bilangan tertentu
	var target int
	fmt.Print("\nMasukkan bilangan yang ingin dihitung frekuensinya: ")
	fmt.Scan(&target)
	frekuensi := hitungFrekuensi(arr, target)
	fmt.Printf("Frekuensi bilangan %d dalam array: %d\n", target, frekuensi)
}
