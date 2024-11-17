package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, hapusIndex, nomer int

	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Println("\nIsi array:")
	fmt.Println(arr)

	fmt.Println("\nElemen dengan indeks ganjil:")
	for i := 1; i < len(arr); i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	fmt.Println("\nElemen dengan indeks genap:")
	for i := 0; i < len(arr); i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	fmt.Print("\nMasukkan nilai x untuk indeks kelipatan: ")
	fmt.Scan(&x)
	fmt.Printf("\nElemen dengan indeks kelipatan %d:\n", x)
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Printf("%d ", arr[i])
		}
	}
	fmt.Println()

	fmt.Print("\nMasukkan indeks yang ingin dihapus: ")
	fmt.Scan(&hapusIndex)
	if hapusIndex >= 0 && hapusIndex < len(arr) {
		arr = append(arr[:hapusIndex], arr[hapusIndex+1:]...)
		fmt.Println("Isi array setelah elemen dihapus:")
		fmt.Println(arr)
	} else {
		fmt.Println("Indeks tidak valid.")
	}

	var sum float64
	for _, value := range arr {
		sum += float64(value)
	}
	average := sum / float64(len(arr))
	fmt.Printf("\nRata-rata elemen dalam array: %.2f\n", average)

	var varianceSum float64
	for _, value := range arr {
		varianceSum += math.Pow(float64(value)-average, 2)
	}
	stdDev := math.Sqrt(varianceSum / float64(len(arr)))
	fmt.Printf("Standar deviasi elemen dalam array: %.2f\n", stdDev)

	fmt.Print("\nMasukkan bilangan yang ingin dihitung frekuensinya: ")
	fmt.Scan(&nomer)
	count := 0
	for _, value := range arr {
		if value == nomer {
			count++
		}
	}
	fmt.Printf("Frekuensi bilangan %d: %d kali\n", nomer, count)
}
