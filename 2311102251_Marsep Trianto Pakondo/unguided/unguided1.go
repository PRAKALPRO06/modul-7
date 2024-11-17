package main

import (
	"fmt"
	"math"
)

func isianArray(bilBul[]int, n int)  {
	fmt.Print("\nIsian dari array : \n")
	for i := 0; i < n; i++ {
		fmt.Printf("Nilai indeks ke-%v : %v\n", i, bilBul[i])
	}
}

func nilaiIndeksGanjil(bilBul[]int, n int)  {
	fmt.Printf("\nNilai indeks Ganjil : \n")

	for i := 0; i < n; i++ {
		if i % 2 == 1 {
			fmt.Printf("Indeks ke-%v. %v\n", i, bilBul[i])
		}
	}
}

func nilaiIndeksGenap(bilBul[]int, n int)  {
	fmt.Printf("\nNilai indeks Genap : \n")

	for i := 0; i < n; i++ {
		if i % 2 == 0 {
			fmt.Printf("Indeks ke-%v. %v\n", i, bilBul[i])
		}
	}
}

func nilaiKelipatanBilangan(bilBul[]int, n int)  {
	var kelipatan int

	fmt.Print("\nMasukkan kelipatan berapa : ")
	fmt.Scan(&kelipatan)
	
	fmt.Printf("\nNilai kelipatan %v :\n", kelipatan)

	for i := 0; i < n; i++ {
		if bilBul[i] % kelipatan == 0 {
			fmt.Printf("%v. %v\n", i+1, bilBul[i])
		}
	}
}

func haspusElemenTertentu(bilBul[]int, n int)  {
	var elemen int

	fmt.Print("\nHapus elemen indeks ke berapa : ")
	fmt.Scan(&elemen)
	
	fmt.Printf("\nNilai elemen %v :\n", elemen)

	for i := 0; i < n; i++ {
		if i == elemen {
			bilBul[i] = ' '
		}
	}
	for i := 0; i < n; i++ {
		if bilBul[i] == ' ' {
			continue
		}
		fmt.Printf("%v. %v\n", i+1, bilBul[i])
	}
}

func rataRata(bilBul[]int, n int) float64 {
	jumlah := 0
	bagi := 0
	for i := 0; i < n; i++ {
		if bilBul[i] == ' ' {
			continue
		}
		jumlah += bilBul[i]
		bagi += 1
	}

	return float64(jumlah) / float64(bagi)
}

func standarDeviasi(bilBul[]int, n int) float64  {
	rata := rataRata(bilBul[:], n)
	jumlah := 0.0 
	bagi := 0

	for i := 0; i < n; i++ {
		if bilBul[i] == ' ' {
			continue
		}
		jumlah += math.Pow(float64(bilBul[i])-rata, 2)
		bagi += 1
	}
	return math.Sqrt(jumlah / float64(bagi))
}

func frekuensi(bilBul[]int, n, cari int) int {
	var count int

	for i := 0; i < n; i++ {
		if bilBul[i] == cari {
			count++
		}
	}

	return count
}


func main()  {
	var n, cari int
	var bilBul[20]int // Asumsi panjang array 20
	
	fmt.Print("Masukkan panjang array : ")
	fmt.Scan(&n)
	
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan bilangan bulat indeks ke-%v : ", i)
		fmt.Scan(&bilBul[i])
	}
	
	// Isi Data Array
	isianArray(bilBul[:], n)

	// Isi Nilai Indeks Ganjil
	nilaiIndeksGanjil(bilBul[:], n)

	// Isi Nilai Indeks Genap
	nilaiIndeksGenap(bilBul[:], n)

	// Mencari Nilai Kelipatan X
	nilaiKelipatanBilangan(bilBul[:], n)

	// Hapus Element di Indeks Tertentu
	haspusElemenTertentu(bilBul[:], n)

	// Rata-rata
	fmt.Print("\nRata-ratanya adalah ", rataRata(bilBul[:], n))

	// Standar Deviasi
	fmt.Print("\n\nStandar deviasinya adalah ", standarDeviasi(bilBul[:], n))

	// Frekuensi
	fmt.Print("\n\nMasukkan bilangan yang ingin dicari frekuensinya : ")
	fmt.Scan(&cari)
	fmt.Printf("\nFrekuensi %v muncul %v kali", cari, frekuensi(bilBul[:], n, cari))
	
}