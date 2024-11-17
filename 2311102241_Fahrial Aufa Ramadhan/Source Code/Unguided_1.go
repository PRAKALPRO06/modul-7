package main

import (
	"fmt"
	"math"
)

func tampilarray_2311102241(array []int, n int) {
	fmt.Print("Isi array: ")
	for i := 0; i < n; i++ {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}

func tampilkanIndeksGanjil(array []int, n int) {
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}

func tampilkanIndeksGenap(array []int, n int) {
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}

func tampilkanKelipatanX(array []int, n int, x int) {
	fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(array[i], " ")
		}
	}
	fmt.Println()
}
func hapusElemen(array []int, n int, indeks int) ([]int, int) {
	if indeks >= 0 && indeks < n {
		array = append(array[:indeks], array[indeks+1:]...)
		fmt.Println("Elemen pada indeks", indeks, "telah dihapus.")
		tampilarray_2311102241(array, n-1)
		return array, n - 1
	} else {
		fmt.Println("Indeks tidak valid.")
		return array, n
	}
}
func hitungRataRata(array []int, n int) float64 {
	total := 0
	for i := 0; i < n; i++ {
		total += array[i]
	}
	return float64(total) / float64(n)
}
func hitungStandarDeviasi(array []int, n int) float64 {
	rataRata := hitungRataRata(array, n)
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += math.Pow(float64(array[i])-rataRata, 2)
	}
	return math.Sqrt(sum / float64(n))
}
func hitungFrekuensi(array []int, n int, bilangan int) int {
	frekuensi := 0
	for i := 0; i < n; i++ {
		if array[i] == bilangan {
			frekuensi++
		}
	}
	return frekuensi
}

func main() {
	var kapasitas, n, x, indeksHapus, bilangan int
	fmt.Print("Masukkan kapasitas array: ")
	fmt.Scanln(&kapasitas)
	array := make([]int, kapasitas)

	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scanln(&n)

	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scanln(&array[i])
	}
	for {
		fmt.Println("\nPilih operasi:")
		fmt.Println("[a.] Tampilkan seluruh isi array")
		fmt.Println("[b.] Tampilkan elemen dengan indeks ganjil")
		fmt.Println("[c.] Tampilkan elemen dengan indeks genap")
		fmt.Println("[d.] Tampilkan elemen dengan indeks kelipatan x")
		fmt.Println("[e.] Hapus elemen pada indeks tertentu")
		fmt.Println("[f.] Tampilkan rata-rata")
		fmt.Println("[g.] Tampilkan standar deviasi")
		fmt.Println("[h.] Tampilkan frekuensi bilangan tertentu")
		fmt.Println("[i.] Keluar")

		var pilihan string
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "a":
			tampilarray_2311102241(array, n)
		case "b":
			tampilkanIndeksGanjil(array, n)
		case "c":
			tampilkanIndeksGenap(array, n)
		case "d":
			fmt.Print("Masukkan nilai x: ")
			fmt.Scanln(&x)
			tampilkanKelipatanX(array, n, x)
		case "e":
			fmt.Print("Masukkan indeks yang akan dihapus: ")
			fmt.Scanln(&indeksHapus)
			array, n = hapusElemen(array, n, indeksHapus)
		case "f":
			fmt.Println("Rata-rata:", hitungRataRata(array, n))
		case "g":
			fmt.Println("Standar deviasi:", hitungStandarDeviasi(array, n))
		case "h":
			fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
			fmt.Scanln(&bilangan)
			fmt.Println("Frekuensi:", hitungFrekuensi(array, n, bilangan))
		case "i":
			fmt.Println("Program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
