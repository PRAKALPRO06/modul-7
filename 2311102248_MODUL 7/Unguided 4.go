package main

// Erwin Rivaldo Silaban
// 2311102248
import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var kar rune
	fmt.Print("Masukkan Teks (akhiri dengan '.'): ")
	for i := 0; i < NMAX; i++ {
		fmt.Scanf("%c", &kar)
		if kar == '.' {
			break
		}
		t[i] = kar
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Print("\n")
}

func balikkanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		temp := t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
	fmt.Print("Teks setelah dibalik: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Print("\n")
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	fmt.Print("Teks asli: ")
	cetakArray(tab, m)

	balikkanArray(&tab, m)

	if palindrom(tab, m) {
		fmt.Println("Apakah teks ini palindrom? true")
	} else {
		fmt.Println("Apakah teks ini palindrom? false")
	}
}
