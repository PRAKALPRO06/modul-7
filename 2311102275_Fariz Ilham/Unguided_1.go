package main

import (
	"fmt"
	"math"
)

func mean(array []int) float64 {
	sum := 0
	for _, v := range array {
		sum += v
	}
	return float64(sum) / float64(len(array))
}

func stdDev(array []int, avg float64) float64 {
	var sum float64
	for _, v := range array {
		sum += math.Pow(float64(v)-avg, 2)
	}
	return math.Sqrt(sum / float64(len(array)))
}

func countFrequency(array []int, target int) int {
	count := 0
	for _, v := range array {
		if v == target {
			count++
		}
	}
	return count
}

func removeIndex(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}

func main() {
	var n, x, target, index int

	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&n)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i+1)
		fmt.Scan(&array[i])
	}

	fmt.Println("a. Menampilkan keseluruhan isi array")
	fmt.Println(array)

	fmt.Println("\nb. Menampilkan elemen dengan indeks ganjil")
	for i := 1; i < n; i += 2 {
		fmt.Printf("array[%d] = %d\n", i, array[i])
	}

	fmt.Println("\nc. Menampilkan elemen dengan indeks genap")
	for i := 0; i < n; i += 2 {
		fmt.Printf("array[%d] = %d\n", i, array[i])
	}

	fmt.Print("\nd. Masukkan bilangan x untuk menampilkan elemen dengan indeks kelipatan x: ")
	fmt.Scan(&x)
	fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Printf("array[%d] = %d\n", i, array[i])
		}
	}

	fmt.Print("\ne. Masukkan indeks yang akan dihapus: ")
	fmt.Scan(&index)
	array = removeIndex(array, index)
	fmt.Println("Isi array setelah penghapusan:")
	fmt.Println(array)

	avg := mean(array)
	fmt.Printf("\nf. Rata-rata elemen array: %.2f\n", avg)

	fmt.Printf("g. Standar deviasi elemen array: %.2f\n", stdDev(array, avg))

	fmt.Print("\nh. Masukkan bilangan yang ingin dihitung frekuensinya: ")
	fmt.Scan(&target)
	fmt.Printf("Frekuensi bilangan %d di dalam array: %d kali\n", target, countFrequency(array, target))
}
