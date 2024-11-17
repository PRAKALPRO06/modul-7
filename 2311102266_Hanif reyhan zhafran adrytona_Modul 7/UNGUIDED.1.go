// 2311102266_Hanif Reyhan Zhafran Arytona

package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, delIndex, target int

	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Jumlah elemen array harus lebih besar dari 0.")
		return
	}

	array := make([]int, n)
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}

	fmt.Println("\nIsi array lengkap:", array)

	fmt.Println("\nElemen dengan indeks ganjil:")
	for i := 1; i < n; i += 2 {
		fmt.Printf("Index %d: %d\n", i, array[i])
	}

	fmt.Println("\nElemen dengan indeks genap:")
	for i := 0; i < n; i += 2 {
		fmt.Printf("Index %d: %d\n", i, array[i])
	}

	fmt.Print("\nMasukkan nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	if x <= 0 {
		fmt.Println("Nilai x harus lebih besar dari 0.")
		return
	}
	fmt.Printf("\nElemen dengan indeks kelipatan %d:\n", x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Printf("Index %d: %d\n", i, array[i])
		}
	}

	fmt.Print("\nMasukkan indeks yang akan dihapus: ")
	fmt.Scan(&delIndex)
	if delIndex >= 0 && delIndex < len(array) {
		array = append(array[:delIndex], array[delIndex+1:]...)
		fmt.Println("Isi array setelah penghapusan:", array)
	} else {
		fmt.Println("Indeks tidak valid.")
		return
	}

	sum := 0
	for _, v := range array {
		sum += v
	}
	rataRata := float64(sum) / float64(len(array))
	fmt.Printf("\nRata-rata elemen array: %.2f\n", rataRata)

	var varianceSum float64
	for _, v := range array {
		varianceSum += math.Pow(float64(v)-rataRata, 2)
	}
	stdDeviasi := math.Sqrt(varianceSum / float64(len(array)))
	fmt.Printf("Standar deviasi elemen array: %.2f\n", stdDeviasi)

	fmt.Print("\nMasukkan bilangan untuk mengetahui frekuensinya: ")
	fmt.Scan(&target)
	count := 0
	for _, v := range array {
		if v == target {
			count++
		}
	}
	fmt.Printf("Frekuensi bilangan %d dalam array: %d\n", target, count)
}
