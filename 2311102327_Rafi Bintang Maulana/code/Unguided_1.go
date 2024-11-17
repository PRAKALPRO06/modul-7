package main
import (
	"fmt"
	"math"
)

func main() {
	var daftarAngka []int
	var jumlah int
	fmt.Print("Masukkan jumlah angka: ")
	fmt.Scan(&jumlah)

	if jumlah > 0 {
		daftarAngka = make([]int, jumlah)
		fmt.Println("Masukkan angka-angka:")
		for i := 0; i < jumlah; i++ {
			fmt.Printf("Angka ke-%d: ", i)
			fmt.Scan(&daftarAngka[i])
		}
	} else {
		fmt.Println("Jumlah angka harus lebih besar dari 0!")
		return
	}
	fmt.Println("\nSemua angka:", daftarAngka)

	fmt.Print("Angka dengan indeks ganjil: ")
	for i := 1; i < len(daftarAngka); i += 2 {
		fmt.Print(daftarAngka[i], " ")
	}
	fmt.Println()

	fmt.Print("Angka dengan indeks genap: ")
	for i := 0; i < len(daftarAngka); i += 2 {
		fmt.Print(daftarAngka[i], " ")
	}
	fmt.Println()

	var kelipatan int
	fmt.Print("Masukkan kelipatan x: ")
	fmt.Scan(&kelipatan)
	if kelipatan > 0 {
		fmt.Printf("Angka dengan indeks kelipatan %d: ", kelipatan)
		for i := 0; i < len(daftarAngka); i++ {
			if i%kelipatan == 0 {
				fmt.Print(daftarAngka[i], " ")
			}
		}
		fmt.Println()
	} else {
		fmt.Println("x harus lebih besar dari 0!")
	}

	for {
		fmt.Println("\n--- Ambil Tindakan ---")
		fmt.Println("1. Hapus angka pada indeks tertentu")
		fmt.Println("2. Hitung rata-rata angka")
		fmt.Println("3. Hitung standar deviasi angka")
		fmt.Println("4. Hitung frekuensi angka tertentu")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var indeks int
			fmt.Print("Masukkan indeks yang mau dihapus: ")
			fmt.Scan(&indeks)
			if indeks >= 0 && indeks < len(daftarAngka) {
				daftarAngka = append(daftarAngka[:indeks], daftarAngka[indeks+1:]...)
				fmt.Println("Daftar angka setelah dihapus:", daftarAngka)
			} else {
				fmt.Println("Indeks nggak valid!")
			}

		case 2:
			if len(daftarAngka) > 0 {
				total := 0
				for _, angka := range daftarAngka {
					total += angka
				}
				rata := float64(total) / float64(len(daftarAngka))
				fmt.Println("Rata-rata angka:", rata)
			} else {
				fmt.Println("Daftar angka kosong!")
			}

		case 3:
			if len(daftarAngka) > 0 {
				total := 0
				for _, angka := range daftarAngka {
					total += angka
				}
				rata := float64(total) / float64(len(daftarAngka))

				var variansi float64
				for _, angka := range daftarAngka {
					variansi += math.Pow(float64(angka)-rata, 2)
				}
				variansi /= float64(len(daftarAngka))
				stdDev := math.Sqrt(variansi)
				fmt.Println("Standar deviasi angka:", stdDev)
			} else {
				fmt.Println("Daftar angka kosong!")
			}

		case 4:
			if len(daftarAngka) > 0 {
				var angkaDicari int
				fmt.Print("Masukkan angka yang mau dihitung frekuensinya: ")
				fmt.Scan(&angkaDicari)
				hitung := 0
				for _, angka := range daftarAngka {
					if angka == angkaDicari {
						hitung++
					}
				}
				fmt.Println("Frekuensi angka", angkaDicari, ":", hitung)
			} else {
				fmt.Println("Daftar angka kosong!")
			}

		case 0:
			fmt.Println("Keluar dari program. Makasih!")
			return

		default:
			fmt.Println("Pilihan nggak valid! Pilih yang ada di menu.")
		}
	}
}