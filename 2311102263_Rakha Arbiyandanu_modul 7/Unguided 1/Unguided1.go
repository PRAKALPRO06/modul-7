package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scanln(&n)

	array := make([]int, n)

	// Mengisi array
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scanln(&array[i])
	}

	// a. Menampilkan keseluruhan isi array
	fmt.Println("\nKeseluruhan isi array:")
	fmt.Println(array)

	// b. Menampilkan elemen array dengan indeks ganjil
	fmt.Println("\nElemen dengan indeks ganjil:")
	for i := 1; i < len(array); i += 2 {
		fmt.Printf("Indeks %d: %d\n", i, array[i])
	}

	// c. Menampilkan elemen array dengan indeks genap
	fmt.Println("\nElemen dengan indeks genap:")
	for i := 0; i < len(array); i += 2 {
		fmt.Printf("Indeks %d: %d\n", i, array[i])
	}

	// d. Menampilkan elemen dengan indeks kelipatan bilangan x
	var x int
	fmt.Print("\nMasukkan bilangan x untuk kelipatan indeks: ")
	fmt.Scanln(&x)
	fmt.Println("Elemen dengan indeks kelipatan", x, ":")
	for i := 0; i < len(array); i++ {
		if i%x == 0 {
			fmt.Printf("Indeks %d: %d\n", i, array[i])
		}
	}

	// e. Menghapus elemen array pada indeks tertentu
	var index int
	fmt.Print("\nMasukkan indeks elemen yang ingin dihapus: ")
	fmt.Scanln(&index)
	if index >= 0 && index < len(array) {
		array = append(array[:index], array[index+1:]...)
		fmt.Println("Array setelah elemen dihapus:")
		fmt.Println(array)
	} else {
		fmt.Println("Indeks tidak valid.")
	}

	// f. Menampilkan rata-rata bilangan dalam array
	sum := 0
	for _, value := range array {
		sum += value
	}
	rataRata := float64(sum) / float64(len(array))
	fmt.Printf("\nRata-rata nilai array: %.2f\n", rataRata)

	// g. Menampilkan standar deviasi (simpangan baku)
	var deviasiSum float64
	for _, value := range array {
		deviasiSum += math.Pow(float64(value)-rataRata, 2)
	}
	standarDeviasi := math.Sqrt(deviasiSum / float64(len(array)))
	fmt.Printf("Simpangan baku array: %.2f\n", standarDeviasi)

	// h. Menampilkan frekuensi suatu bilangan
	var target int
	fmt.Print("\nMasukkan bilangan untuk menghitung frekuensinya: ")
	fmt.Scanln(&target)
	frekuensi := 0
	for _, value := range array {
		if value == target {
			frekuensi++
		}
	}
	fmt.Printf("Frekuensi bilangan %d: %d kali\n", target, frekuensi)
}
