package main

import "fmt"

const a_size int = 8
var a = [a_size] int {1, 2, 3, 4, 5, 6, 7, 8}

func tampilkanIsiArray() {
	for i:=0; i<a_size; i++ {
		if(a[i] != 0) {
			fmt.Print(a[i],", ")
		}
	}
}
func tampilkanElIdxGanjil() {
	for i:=1; i<a_size; i+=2 {
		if(a[i] != 0) {
			fmt.Print(a[i],", ")
		}
	}
}
func tampilkanElIdxGenap() {
	for i:=0; i<a_size; i+=2 {
		if(a[i] != 0) {
			fmt.Print(a[i],", ")
		}
	}
}
func tampilkanElIdxKelipatan(x int) {
	for i:=0; i<a_size; i+=x {
		if(a[i] != 0) {
			fmt.Print(a[i],", ")
		}
	}
}
func hapusEl(idx int) {
	a[idx] = 0
}
func hitungMean() float64 {
	var sum float64 = 0
	count := 0
	for i:=0; i<a_size; i++ {
		if(a[i] != 0) {
			sum += float64(a[i])
			count++
		}
	}
	mean := sum/float64(count)
	return mean
}
func tampilkanMean() {
	fmt.Println("Rata-rata elemen dalam array adalah ", hitungMean())
}
func hitungSTD() float64 {
	mean := hitungMean()
	var sum float64 = 0
	count := 0
	for i:=0; i<a_size; i++ {
		if(a[i] != 0) {
			sum += (float64(a[i]) - mean)*(float64(a[i]) - mean)
			count++
		}
	}
	std := sum/float64(count)
	return std
}
func tampilkanSTD() {
	fmt.Println("Standar deviasi elemen dalam array adalah ", hitungSTD())
}
func tampilkanFrekuensi(x int) {
	count := 0
	for i:=0; i<a_size; i++ {
		if(a[i] == x) {
			count++
		}
	}
	fmt.Println("Frekuensi ", x, " adalah ", count)
}

func main() {


	fmt.Println("Pilih menu:")
	fmt.Println("1. Tampilkan isi array")
	fmt.Println("2. Tampilkan semua elemen array dengan indeks ganjil")
	fmt.Println("3. Tampilkan semua elemen array dengan indeks genap")
	fmt.Println("4. Tampilkan semua elemen array dengan indeks kelipatan x")
	fmt.Println("5. Hapus elemen array")
	fmt.Println("6. Tampilkan rata-rata dari semua elemen array")
	fmt.Println("7. Tampilkan standar deviasi semua elemen array")
	fmt.Println("8. Tampilkan frekuensi elemen array")
	fmt.Print("Pilihan: ")

	var pil int
	fmt.Scan(&pil)

	if(pil == 1) {
		tampilkanIsiArray()
	} else if(pil == 2) {
		tampilkanElIdxGanjil()
	} else if(pil == 3) {
		tampilkanElIdxGenap()
	} else if(pil == 4) {
		var x int
		fmt.Print("Masukkan angka kelipatan : ")
		fmt.Scan(&x)
		tampilkanElIdxKelipatan(x)
	} else if(pil == 5) {
		var x int
		fmt.Print("Masukkan index elemen yang ingin dihapus : ")
		fmt.Scan(&x)
		hapusEl(x)
	} else if(pil == 6) {
		tampilkanMean()
	} else if(pil == 7) {
		tampilkanSTD()
	} else if(pil == 8) {
		var x int
		fmt.Print("Masukkan elemen : ")
		fmt.Scan(&x)
		tampilkanFrekuensi(x)
	} else {
		fmt.Println("Input tidak valid")
	}
}