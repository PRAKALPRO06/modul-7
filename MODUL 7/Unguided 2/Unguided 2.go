package main

import (
	"fmt"
	"math"
)


func cetakArray(arr []int) {
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}

func hitungRata2_StandarDeviasi(arr []int) (float64, float64) {
	jumlah := 0
	for _, val := range arr {
		jumlah += val
	}
	rata := float64(jumlah) / float64(len(arr))

	var jumlahVariasi float64
	for _, val := range arr {
		jumlahVariasi += math.Pow(float64(val)-rata, 2)
	}
	variasi:= jumlahVariasi / float64(len(arr))
	stdDev := math.Sqrt(variasi)

	return rata, stdDev
}


func hitungFrekuensi(arr []int, num int) int {
	count := 0
	for _, val := range arr {
		if val == num {
			count++
		}
	}
	return count
}

func main() {
	var N, x, index, num int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&N)

	arr := make([]int, N)
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Isi array:")
	cetakArray(arr)

	fmt.Println("Elemen dengan indeks ganjil:")
	for i := 1; i < len(arr); i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
	
	fmt.Println("Elemen dengan indeks genap:")
	    for _, val := range arr {
		if val%2 == 0 { 
			fmt.Printf("%d ", val)
		}
	}
	fmt.Println()


	fmt.Print("Masukkan nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Println("Elemen dengan indeks kelipatan", x, ":")
	for _, val := range arr {
		if val%x == 0 {
			fmt.Printf("%d ", val)
		}
	}
	fmt.Println()

	fmt.Print("Masukkan indeks untuk dihapus: ")
	fmt.Scan(&index)
	if index >= 0 && index < len(arr) {
		arr = append(arr[:index], arr[index+1:]...)
		fmt.Println("Isi array setelah penghapusan:")
		cetakArray(arr)
	} else {
		fmt.Println("Index tidak valid")
	}

	rata2, stdDev := hitungRata2_StandarDeviasi(arr)
	fmt.Printf("Rata-rata: %.2f\n", rata2)
	fmt.Printf("Standar deviasi: %.2f\n", stdDev)

	fmt.Print("Masukkan bilangan untuk menghitung frekuensinya: ")
	fmt.Scan(&num)
	frekuensi := hitungFrekuensi(arr, num)
	fmt.Printf("Frekuensi bilangan %d: %d\n", num, frekuensi)
}