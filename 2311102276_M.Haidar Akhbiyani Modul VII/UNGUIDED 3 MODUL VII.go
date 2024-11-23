package main

import (
	"fmt"
	"strings"
)

func balikArray(arr []rune) []rune {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
	return arr
}

func cekPalindrom(arr []rune) bool {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		if arr[i] != arr[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var input1, input2 string

	fmt.Print("Teks: ")
	fmt.Scanln(&input1)

	txtLower1 := strings.ToLower(input1)
	array1 := []rune(txtLower1)
	balik1 := balikArray(array1)

	fmt.Printf("Reverse teks: %s\n", strings.ToLower(string(balik1)))
	fmt.Println()

	fmt.Print("Teks: ")
	fmt.Scanln(&input2)

	txtLower2 := strings.ToLower(input2)
	array2 := []rune(txtLower2)
	balik2 := balikArray(array2)

	fmt.Printf("Reverse teks: %s\n", strings.ToLower(string(balik2)))
	fmt.Println()

	isPalindrom1 := cekPalindrom(array1)
	fmt.Printf("Teks :%s\n", input1)
	fmt.Printf("polindrom ?%t\n", isPalindrom1)
	fmt.Println()

	isPalindrom2 := cekPalindrom(array2)
	fmt.Printf("Teks :%s\n", input2)
	fmt.Printf("polindrom ?%t\n", isPalindrom2)
}