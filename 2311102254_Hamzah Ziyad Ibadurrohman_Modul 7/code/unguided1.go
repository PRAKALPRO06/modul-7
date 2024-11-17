package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	// Deklarasi array dengan kapasitas sesuai input N
	arr := make([]int, n)

	// Mengisi array
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	// a. Menampilkan keseluruhan isi dari array
	fmt.Println("Isi array:", arr)

	// b. Menampilkan elemen-elemen array dengan indeks ganjil saja
	fmt.Println("Elemen dengan indeks ganjil:")
	for i := 1; i < n; i += 2 {
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}

	// c. Menampilkan elemen-elemen array dengan indeks genap
	fmt.Println("Elemen dengan indeks genap:")
	for i := 0; i < n; i += 2 {
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}

	// d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x
	var x int
	fmt.Print("Masukkan nilai x untuk menampilkan elemen dengan indeks kelipatan x: ")
	fmt.Scan(&x)
	fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Printf("arr[%d] = %d\n", i, arr[i])
		}
	}

	// e. Menghapus elemen array pada indeks tertentu
	var delIndx int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&delIndx)
	if delIndx >= 0 && delIndx < n {
		arr = append(arr[:delIndx], arr[delIndx+1:]...)
		fmt.Println("Isi array setelah penghapusan:", arr)
	} else {
		fmt.Println("Indeks tidak valid.")
	}

	// f. Menampilkan rata-rata dari bilangan yang ada di dalam array
	sum := 0
	for _, val := range arr {
		sum += val
	}
	ratarata := float64(sum) / float64(len(arr))
	fmt.Printf("Rata-rata dari elemen dalam array: %.2f\n", ratarata)

	// g. Menampilkan standar deviasi atau simpangan baku dari bilangan dalam array
	var varianceSum float64
	for _, val := range arr {
		varianceSum += math.Pow(float64(val)-ratarata, 2)
	}
	stdDev := math.Sqrt(varianceSum / float64(len(arr)))
	fmt.Printf("Simpangan baku dari elemen dalam array: %.2f\n", stdDev)

	// h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array
	var target int
	fmt.Print("Masukkan nilai untuk menghitung frekuensinya dalam array: ")
	fmt.Scan(&target)
	frekuensi := 0
	for _, val := range arr {
		if val == target {
			frekuensi++
		}
	}
	fmt.Printf("Frekuensi %d di dalam array: %d\n", target, frekuensi)
}
