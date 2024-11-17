//Wisnu Rananta Raditya Putra (2311102013) IF-11-06

package main

import (
	"fmt"
	"math"
)

func main() {
	var n_2311102013 int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scanln(&n_2311102013)

	array := make([]int, n_2311102013)

	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n_2311102013; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scanln(&array[i])
	}

	fmt.Println("\nKeseluruhan isi array:")
	fmt.Println(array)

	fmt.Println("\nElemen dengan indeks ganjil:")
	for i := 1; i < len(array); i += 2 {
		fmt.Printf("Indeks %d: %d\n", i, array[i])
	}

	fmt.Println("\nElemen dengan indeks genap:")
	for i := 0; i < len(array); i += 2 {
		fmt.Printf("Indeks %d: %d\n", i, array[i])
	}

	var x int
	fmt.Print("\nMasukkan bilangan x untuk kelipatan indeks: ")
	fmt.Scanln(&x)
	fmt.Println("Elemen dengan indeks kelipatan", x, ":")
	for i := 0; i < len(array); i++ {
		if i%x == 0 {
			fmt.Printf("Indeks %d: %d\n", i, array[i])
		}
	}

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

	sum := 0
	for _, value := range array {
		sum += value
	}
	rataRata := float64(sum) / float64(len(array))
	fmt.Printf("\nRata-rata nilai array: %.2f\n", rataRata)

	var deviasiSum float64
	for _, value := range array {
		deviasiSum += math.Pow(float64(value)-rataRata, 2)
	}
	standarDeviasi := math.Sqrt(deviasiSum / float64(len(array)))
	fmt.Printf("Simpangan baku array: %.2f\n", standarDeviasi)

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