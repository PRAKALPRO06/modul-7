package main

import (
	"fmt"
	"math"
)

const NMAX int = 127

type tabel [NMAX]rune

var tab tabel
var m int

func isiArray(t *tabel, n *int) {
	fmt.Print("Teks: ")
	for {
		var a rune
		_, err := fmt.Scanf("%c", &a) 
		if err != nil || a == '.' {   
			break
		}
		if *n < NMAX { 
			t[*n] = a
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]))
	}
	fmt.Print("\n")
}

func balikkanArray(t *tabel, n int) {
	for i := 0; i < int(math.Round(float64(n/2))); i++ {
		temp := t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}

func Palindrom(t tabel, n int) bool {
	for i := 0; i < int(math.Round(float64(n/2))); i++ {
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
	fmt.Print("Reverse Teks: ")
	balikkanArray(&tab, m)
	cetakArray(tab, m)
	fmt.Print("Palindrom: ")
	fmt.Print(Palindrom(tab, m))
}