package main
import (
	"fmt"
)


const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0
	for {
		fmt.Scanf("%c", &ch)
		if ch == '\n' || ch == 'T' {
			break
		}
		t[*n] = ch
		*n++
		if *n >= NMAX {
			break
		}
	}
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
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
	var n int

	fmt.Print("Masukkan teks: ")
	isiArray(&tab, &n)

	fmt.Print("Teks: ")
	cetakArray(tab, n)

	balikanArray(&tab, n)
	fmt.Print("Reverse teks: ")
	cetakArray(tab, n)

	if palindrom(tab, n) {
		fmt.Println("Teks ini adalah palindrom.")
	} else {
		fmt.Println("Teks ini bukan palindrom.")
	}
}

//Reza Alvonzo 2311102026 IF 06