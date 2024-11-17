package main

import (
	"fmt"
	"math"
)

type Statistik struct {
	data []float64
}

func BuatStatistik(ukuran int) *Statistik {
	angka := make([]float64, ukuran)
	fmt.Println("Masukkan", ukuran, "bilangan:")
	for i := 0; i < ukuran; i++ {
		fmt.Scan(&angka[i])
	}
	return &Statistik{data: angka}
}

func (s *Statistik) ProsesData() {
	fmt.Println("\na. Keseluruhan isi array:")
	fmt.Println(s.data)

	fmt.Print("\nb. Elemen dengan indeks ganjil: ")
	for i := 1; i < len(s.data); i += 2 {
		fmt.Printf("%.2f ", s.data[i])
	}
	fmt.Println()

	fmt.Print("\nc. Elemen dengan indeks genap: ")
	for i := 0; i < len(s.data); i += 2 {
		fmt.Printf("%.2f ", s.data[i])
	}
	fmt.Println()

	var kelipatan int
	fmt.Print("\nMasukkan nilai kelipatan untuk menampilkan elemen: ")
	fmt.Scan(&kelipatan)
	fmt.Printf("\nd. Elemen dengan indeks kelipatan %d: ", kelipatan)
	for i := 0; i < len(s.data); i += kelipatan {
		fmt.Printf("%.2f ", s.data[i])
	}
	fmt.Println()

	var indeksHapus int
	fmt.Print("\nMasukkan indeks yang akan dihapus: ")
	fmt.Scan(&indeksHapus)
	if indeksHapus >= 0 && indeksHapus < len(s.data) {
		s.data = append(s.data[:indeksHapus], s.data[indeksHapus+1:]...)
		fmt.Println("\ne. Array setelah penghapusan indeks", indeksHapus, ":")
		fmt.Println(s.data)
	}

	total := 0.0
	for _, nilai := range s.data {
		total += nilai
	}
	average := total / float64(len(s.data))
	fmt.Printf("\nf. Rata-rata: %.2f\n", average)

	var totalKuadrat float64
	for _, nilai := range s.data {
		totalKuadrat += (nilai - average) * (nilai - average)
	}
	stdev := math.Sqrt(totalKuadrat / float64(len(s.data)))
	fmt.Printf("\ng. Simpangan baku: %.2f\n", stdev)

	var angkaFrekuensi float64
	fmt.Print("\nMasukkan bilangan yang ingin dihitung frekuensinya: ")
	fmt.Scan(&angkaFrekuensi)
	frekuensi := 0
	for _, nilai := range s.data {
		if nilai == angkaFrekuensi {
			frekuensi++
		}
	}
	fmt.Printf("\nh. Frekuensi bilangan %.2f: %d kali\n", angkaFrekuensi, frekuensi)
}

func main() {
	var jumlahData int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&jumlahData)

	statistik := BuatStatistik(jumlahData)
	statistik.ProsesData()
}
