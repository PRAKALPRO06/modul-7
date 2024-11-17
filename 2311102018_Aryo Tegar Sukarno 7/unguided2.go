//2311102018 ARYO TEGAR SUKARNO
package main

import (
	"fmt"
	"math"
)
func calculateMean(arr []int) float64 {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return float64(sum) / float64(len(arr))
}
func calculateStdDev(arr []int) float64 {
	mean := calculateMean(arr)
	var varianceSum float64
	for _, v := range arr {
		varianceSum += math.Pow(float64(v)-mean, 2)
	}
	return math.Sqrt(varianceSum / float64(len(arr)))
}
func calculateFrequency(arr []int, target int) int {
	frequency := 0
	for _, v := range arr {
		if v == target {
			frequency++
		}
	}
	return frequency
}

func main() {
	var n, x, indexToRemove, targetNumber int

	fmt.Print("Masukkan jumlah elemen array: ")
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
		fmt.Println(arr[i])
	}
	
	fmt.Println("\nElemen dengan indeks genap:")
	for i := 0; i < len(arr); i += 2 {
		fmt.Println(arr[i])
	}
	fmt.Print("\nMasukkan bilangan x untuk indeks kelipatan: ")
	fmt.Scan(&x)
	fmt.Println("Elemen dengan indeks kelipatan", x, ":")
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Println(arr[i])
		}
	}
	fmt.Print("\nMasukkan indeks elemen yang ingin dihapus: ")
	fmt.Scan(&indexToRemove)
	if indexToRemove >= 0 && indexToRemove < len(arr) {
		arr = append(arr[:indexToRemove], arr[indexToRemove+1:]...)
		fmt.Println("Array setelah penghapusan elemen:")
		fmt.Println(arr)
	} else {
		fmt.Println("Indeks tidak valid!")
	}
	mean := calculateMean(arr)
	fmt.Printf("\nRata-rata elemen array: %.2f\n", mean)
	
	stdDev := calculateStdDev(arr)
	fmt.Printf("Standar deviasi elemen array: %.2f\n", stdDev)

	fmt.Print("\nMasukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&targetNumber)
	frequency := calculateFrequency(arr, targetNumber)
	fmt.Printf("Frekuensi bilangan %d dalam array: %d\n", targetNumber, frequency)
}
