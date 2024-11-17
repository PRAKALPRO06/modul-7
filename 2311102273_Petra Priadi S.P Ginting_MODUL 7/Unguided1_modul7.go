package main

import (
	"fmt"
	"math"
)

func main() {
	var p, pilihan int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&p)

	arr := make([]int, p)

	for i := 0; i < p; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	for {
		fmt.Println("\nMenu (PetraGinting_2311102273):")
		fmt.Println("1. Tampilkan seluruh array")
		fmt.Println("2. Tampilkan elemen indeks ganjil")
		fmt.Println("3. Tampilkan elemen indeks genap")
		fmt.Println("4. Tampilkan elemen kelipatan X")
		fmt.Println("5. Hapus elemen pada indeks tertentu")
		fmt.Println("6. Tampilkan rata-rata")
		fmt.Println("7. Tampilkan standar deviasi")
		fmt.Println("8. Tampilkan frekuensi bilangan")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilkanArray(arr)
		case 2:
			tampilkanIndeksGanjil(arr)
		case 3:
			tampilkanIndeksGenap(arr)
		case 4:
			var x int
			fmt.Print("Masukkan nilai X: ")
			fmt.Scan(&x)
			tampilkanKelipatanX(arr, x)
		case 5:
			var idx int
			fmt.Print("Masukkan indeks yang akan dihapus: ")
			fmt.Scan(&idx)
			arr = hapusElemen(arr, idx)
		case 6:
			fmt.Printf("Rata-rata: %.2f\n", hitungRataRata(arr))
		case 7:
			fmt.Printf("Standar deviasi: %.2f\n", hitungStandarDeviasi(arr))
		case 8:
			var num int
			fmt.Print("Masukkan bilangan yang dicari: ")
			fmt.Scan(&num)
			fmt.Printf("Frekuensi %d: %d\n", num, hitungFrekuensi(arr, num))
		case 9:
			fmt.Println("Program selesai")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func tampilkanArray(arr []int) {
	fmt.Print("Isi array: ")
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func tampilkanIndeksGanjil(arr []int) {
	fmt.Print("Elemen indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func tampilkanIndeksGenap(arr []int) {
	fmt.Print("Elemen indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func tampilkanKelipatanX(arr []int, x int) {
	fmt.Printf("Elemen kelipatan %d: ", x)
	for _, v := range arr {
		if v%x == 0 {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()
}

func hapusElemen(arr []int, idx int) []int {
	return append(arr[:idx], arr[idx+1:]...)
}

func hitungRataRata(arr []int) float64 {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return float64(sum) / float64(len(arr))
}

func hitungStandarDeviasi(arr []int) float64 {
	mean := hitungRataRata(arr)
	var sumSquares float64
	for _, v := range arr {
		diff := float64(v) - mean
		sumSquares += diff * diff
	}
	variance := sumSquares / float64(len(arr))
	return math.Sqrt(variance)
}

func hitungFrekuensi(arr []int, num int) int {
	count := 0
	for _, v := range arr {
		if v == num {
			count++
		}
	}
	return count
}
