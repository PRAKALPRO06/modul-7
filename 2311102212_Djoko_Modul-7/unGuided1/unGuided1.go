package main

import (
    "fmt"
	"math"
)

func main() {
        var bilangan [10]int

        for i := 0; i < 10; i++ {
            bilangan[i] = i
        }

		count := 10
		// Jawaban A
		fmt.Print("A.Array : ")
        for i := 0; i < count; i++ {
            fmt.Printf("%d ", bilangan[i])
        }

		// Jawab B
		fmt.Print("\nB.Ganjil : ")
		for i := 0; i < count; i++ {
			if bilangan[i] % 2 == 1 {
				fmt.Printf("%d ", bilangan[i])
			}
		}

		// Jawab C
		fmt.Print("\nC.Genap : ")
		for i := 0; i < count; i++ {
			if bilangan[i] == 0 {
				fmt.Printf("%d ", bilangan[i])
			}else if bilangan[i] % 2 == 0 {
				fmt.Printf("%d ", bilangan[i])
			}
		}

		// Jawab D
		fmt.Print("\nD.Kelipatan : ")
		var kelipatan int
		fmt.Scan(&kelipatan)
		for i := 0; i < count; i++ {
			if bilangan[i] % kelipatan == 0 {
				fmt.Printf("%d ", bilangan[i])
			}
		}

		// Jawab E
		var slice []int = make([]int, len(bilangan))
		copy(slice, bilangan[:])

		fmt.Print("\nE. Indeks : ")
		var bilHapus int
		fmt.Scan(&bilHapus)

		// Pengecekan indeks valid
		if bilHapus >= 0 && bilHapus < len(slice) {
			// Hapus elemen pada indeks
			slice = append(slice[:bilHapus], slice[bilHapus+1:]...)
			fmt.Println("Array setelah penghapusan:", slice)
		} else {
			fmt.Println("Indeks tidak valid")
		}

		// Jawab F
		fmt.Print("\nF.")
		var resultAvg int
		for i := 0; i < count; i++ {
			resultAvg += bilangan[i]
		}
		resultAvg = resultAvg/count
		fmt.Printf("Rata rata : %d", resultAvg)

		// Jawab G
		fmt.Print("\nG.")
		var variance float64
        for _, val := range bilangan {
                variance += math.Pow(float64(val-resultAvg), 2)
        }
        variance /= float64(count - 1)

        stdDev := math.Sqrt(variance)

        fmt.Printf("%.2f\n", stdDev)

		// Jawab H
		fmt.Print("\nH.")
		var sliceForH []int = make([]int, len(bilangan))
		copy(sliceForH, bilangan[:])
		frekuensi := hitungFrekuensiDenganArray(sliceForH)

        // Cetak hasil
        for i := 0; i <= 100; i++ {
                if frekuensi[i] > 0 {
                        fmt.Printf("Angka %d muncul sebanyak %d kali\n", i, frekuensi[i])
                }
        }
}

func hitungFrekuensiDenganArray(data []int) []int {
	// Asumsikan nilai maksimum dalam data adalah 100 (sesuaikan jika perlu)
	maxVal := 100
	freq := make([]int, maxVal+1) // Index 0 sampai maxVal

	for _, val := range data {
			freq[val]++
	}

	return freq
}