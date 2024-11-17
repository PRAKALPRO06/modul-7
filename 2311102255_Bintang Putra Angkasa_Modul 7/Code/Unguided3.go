package main

import (
	"fmt"
)

func balikkanArray(array []rune) []rune {
	n := len(array)
	hasil := make([]rune, n)
	for i, val := range array {
		hasil[n-i-1] = val
	}
	return hasil
}

func cekPalindrom(array []rune) bool {
	n := len(array)
	for i := 0; i < n/2; i++ {
		if array[i] != array[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var input string
	fmt.Print("Masukkan sekumpulan karakter: ")
	fmt.Scanln(&input)

	array := []rune(input)
	arrayTerbalik := balikkanArray(array)
	isPalindrom := cekPalindrom(array)

	fmt.Println("\nArray asli: ", string(array))
	fmt.Println("Array terbalik: ", string(arrayTerbalik))
	if isPalindrom {
		fmt.Println("Array tersebut merupakan palindrom.")
	} else {
		fmt.Println("Array tersebut bukan palindrom.")
	}
}
