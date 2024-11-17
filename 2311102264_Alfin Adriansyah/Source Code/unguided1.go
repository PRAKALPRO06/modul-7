package main

import (
	"fmt"
	"math"
)

func main() {
	var jumlahElemen int
	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&jumlahElemen)

	// Membuat dan mengisi array
	bilangan := make([]int, jumlahElemen)
	for i := 0; i < jumlahElemen; i++ {
		fmt.Printf("Masukkan bilangan ke-%d: ", i)
		fmt.Scan(&bilangan[i])
	}

	// Menu utama program
	for {
		fmt.Println("\nPilihan Menu:")
		fmt.Println("1. Tampilkan semua isi array")
		fmt.Println("2. Tampilkan elemen dengan indeks ganjil")
		fmt.Println("3. Tampilkan elemen dengan indeks genap")
		fmt.Println("4. Tampilkan elemen dengan indeks kelipatan tertentu")
		fmt.Println("5. Hapus elemen pada indeks tertentu")
		fmt.Println("6. Hitung rata-rata")
		fmt.Println("7. Hitung standar deviasi")
		fmt.Println("8. Hitung frekuensi kemunculan suatu bilangan")
		fmt.Println("9. Keluar dari program")

		var pilihan int
		fmt.Print("Masukkan pilihan menu (1-9): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilkanArray(bilangan)
		case 2:
			tampilkanIndeksGanjil(bilangan)
		case 3:
			tampilkanIndeksGenap(bilangan)
		case 4:
			var kelipatan int
			fmt.Print("Masukkan nilai kelipatan: ")
			fmt.Scan(&kelipatan)
			tampilkanIndeksKelipatan(bilangan, kelipatan)
		case 5:
			var indeks int
			fmt.Print("Masukkan indeks yang akan dihapus: ")
			fmt.Scan(&indeks)
			bilangan = hapusElemen(bilangan, indeks)
		case 6:
			rerata := hitungRataRata(bilangan)
			fmt.Printf("Rata-rata: %.2f\n", rerata)
		case 7:
			simpanganBaku := hitungStandarDeviasi(bilangan)
			fmt.Printf("Standar deviasi: %.2f\n", simpanganBaku)
		case 8:
			var nilai int
			fmt.Print("Masukkan bilangan yang ingin dihitung frekuensinya: ")
			fmt.Scan(&nilai)
			frekuensi := hitungFrekuensi(bilangan, nilai)
			fmt.Printf("Frekuensi kemunculan bilangan %d: %d kali\n", nilai, frekuensi)
		case 9:
			fmt.Println("Terima kasih, program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid! Silakan pilih menu 1-9.")
		}
	}
}

func tampilkanArray(arr []int) {
	fmt.Print("Isi array: ")
	for i, nilai := range arr {
		fmt.Printf("[%d]=%d ", i, nilai)
	}
	fmt.Println()
}

func tampilkanIndeksGanjil(arr []int) {
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Printf("[%d]=%d ", i, arr[i])
	}
	fmt.Println()
}

func tampilkanIndeksGenap(arr []int) {
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Printf("[%d]=%d ", i, arr[i])
	}
	fmt.Println()
}

func tampilkanIndeksKelipatan(arr []int, kelipatan int) {
	fmt.Printf("Elemen dengan indeks kelipatan %d: ", kelipatan)
	for i := 0; i < len(arr); i++ {
		if i%kelipatan == 0 {
			fmt.Printf("[%d]=%d ", i, arr[i])
		}
	}
	fmt.Println()
}

func hapusElemen(arr []int, indeks int) []int {
	if indeks < 0 || indeks >= len(arr) {
		fmt.Println("Indeks tidak valid!")
		return arr
	}

	hasil := append(arr[:indeks], arr[indeks+1:]...)
	fmt.Println("Elemen berhasil dihapus!")
	return hasil
}

func hitungRataRata(arr []int) float64 {
	if len(arr) == 0 {
		return 0
	}
	jumlah := 0
	for _, nilai := range arr {
		jumlah += nilai
	}
	return float64(jumlah) / float64(len(arr))
}

func hitungStandarDeviasi(arr []int) float64 {
	if len(arr) == 0 {
		return 0
	}
	rerata := hitungRataRata(arr)

	var jumlahKuadratSelisih float64
	for _, nilai := range arr {
		selisih := float64(nilai) - rerata
		jumlahKuadratSelisih += selisih * selisih
	}

	variansi := jumlahKuadratSelisih / float64(len(arr))
	return math.Sqrt(variansi)
}

func hitungFrekuensi(arr []int, nilai int) int {
	frekuensi := 0
	for _, num := range arr {
		if num == nilai {
			frekuensi++
		}
	}
	return frekuensi
}
