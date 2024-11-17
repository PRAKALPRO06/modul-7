package main
//2311102018ARYO TEGAR
import (
	"fmt"
)

const MAX = 127

type Tabel struct {
	Tab [MAX]rune 
	N   int      
}


func isiArray(t *Tabel) {
	fmt.Println("Masukkan teks (akhiri dengan TITIK):")
	t.N = 0
	for {
		var input rune
		_, err := fmt.Scanf("%c", &input) 
				if err != nil {
			break
		}
		if input == '.' { 
			break
		}
		t.Tab[t.N] = input
		t.N++
		if t.N >= MAX {
			break
		}
	}
}


func cetakArray(t Tabel) {
	for i := 0; i < t.N; i++ {
		fmt.Print(string(t.Tab[i]))
	}
	fmt.Println()
}


func balikArray(t *Tabel) {
	for i, j := 0, t.N-1; i < j; i, j = i+1, j-1 {
		t.Tab[i], t.Tab[j] = t.Tab[j], t.Tab[i]
	}
}


func palindrome(t Tabel) bool {
	for i := 0; i < t.N/2; i++ {
		if t.Tab[i] != t.Tab[t.N-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab Tabel

	isiArray(&tab)
	fmt.Println("\nTeks:")
	cetakArray(tab)
	balikArray(&tab)
	fmt.Println("\nReverse Teks:")
	cetakArray(tab)
	fmt.Println("\nPalindrome?")
	if palindrome(tab) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
