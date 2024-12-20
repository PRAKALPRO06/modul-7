package main

import "fmt"

type Klub struct {
	skor1 int
	skor2 int
	hasil string
}

func main() {

	var klub1, klub2 string
	var jumlahMatch int
	var match [100]Klub

	fmt.Print("Masukkan Klub 1: ")
	fmt.Scanln(&klub1)
	fmt.Print("Masukkan Klub 2: ")
	fmt.Scanln(&klub2)

	for i := 0; ; i++ {
		fmt.Printf("Skor Pertandingan %d :  ", i+1)
		fmt.Printf("\n %s : ", klub1)
		fmt.Scanln(&match[i].skor1)
		fmt.Printf("%s : ", klub2)
		fmt.Scanln(&match[i].skor2)
		

		if match[i].skor1 < 0 || match[i].skor2 < 0 {
			fmt.Print("Pertandingan Selesai")
			break
		}

		if match[i].skor1 > match[i].skor2 {
			match[i].hasil = klub1
		} else if match[i].skor1 == match[i].skor2 {
			match[i].hasil = "Draw"
		} else {
			match[i].hasil = klub2
		}

		jumlahMatch++
	}

	fmt.Println("\nHasil pertandingan:")
	for i := 0; i < jumlahMatch; i++ {
		fmt.Printf("Hasil pertandingan %d: %s\n", i+1, match[i].hasil)
	}
}
