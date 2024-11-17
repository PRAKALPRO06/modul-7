package main

import (
	"fmt"
	"math"
)

func hitungRataRata(arr []int) float64 {
	total := 0
	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}
	return float64(total) / float64(len(arr))
}

func hitungStandarDeviasi(arr []int, rataRata float64) float64 {
	var jumlahKuadrat float64
	for i := 0; i < len(arr); i++ {
		selisih := float64(arr[i]) - rataRata
		jumlahKuadrat += selisih * selisih
	}
	return math.Sqrt(jumlahKuadrat / float64(len(arr)))
}

func hitungFrekuensi(arr []int, target int) int {
	jumlah := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			jumlah++
		}
	}
	return jumlah
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	fmt.Println("\nIsi array:", arr)

	fmt.Println("\nElemen dengan indeks ganjil:")
	for i := 1; i < len(arr); i += 2 {
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}

	fmt.Println("\nElemen dengan indeks genap:")
	for i := 0; i < len(arr); i += 2 {
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}

	var x int
	fmt.Print("\nMasukkan bilangan kelipatan: ")
	fmt.Scan(&x)
	fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
	for i := x; i < len(arr); i += x {
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}

	var idxHapus int
	fmt.Print("\nMasukkan indeks yang ingin dihapus: ")
	fmt.Scan(&idxHapus)
	if idxHapus >= 0 && idxHapus < len(arr) {
		arr = append(arr[:idxHapus], arr[idxHapus+1:]...)
		fmt.Println("Isi array setelah dihapus:", arr)
	} else {
		fmt.Println("Indeks tidak valid!")
	}

	if len(arr) > 0 {
		rataRata := hitungRataRata(arr)
		fmt.Printf("\nRata-rata array: %.2f\n", rataRata)

		stdDeviasi := hitungStandarDeviasi(arr, rataRata)
		fmt.Printf("Standar deviasi array: %.2f\n", stdDeviasi)
	} else {
		fmt.Println("Array kosong, tidak bisa hitung rata-rata atau standar deviasi.")
	}

	var target int
	fmt.Print("\nMasukkan bilangan untuk hitung frekuensi: ")
	fmt.Scan(&target)
	frekuensi := hitungFrekuensi(arr, target)
	fmt.Printf("Frekuensi bilangan %d di array: %d\n", target, frekuensi)
}
