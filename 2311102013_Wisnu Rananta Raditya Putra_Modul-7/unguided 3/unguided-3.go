//Wisnu Rananta Raditya Putra (2311102013) IF-11-06

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NMAX int = 127

type tabel [NMAX]rune 

func isiArray(t *tabel, n *int, line string) {
	*n = 0
	for _, char := range line {
		if *n >= NMAX {
			break
		}
		t[*n] = char
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]), " ")
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
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
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Masukkan teks (ketik '.' untuk berhenti):")

	for scanner.Scan() {
		line := scanner.Text()
		if strings.ToUpper(line) == "." {
			break
		}

		var tab tabel
		var m int
		isiArray(&tab, &m, line)

		fmt.Print("Teks: ")
		cetakArray(tab, m)

		balikanArray(&tab, m)
		fmt.Print("Reverse Teks: ")
		cetakArray(tab, m)
		
		isPalindrom := palindrom(tab, m)
		fmt.Println("Palindrom:", isPalindrom)
		fmt.Println()
	}
}