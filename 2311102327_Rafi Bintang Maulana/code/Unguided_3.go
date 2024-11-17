package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NMAX int = 127
type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	fmt.Print("Teks\t\t: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ToUpper(input)
	*n = 0
	for _, char := range input {
		if char == '.' {
			break
		}
		if *n < NMAX {
			t[*n] = char
			*n++
		}
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)
	
	fmt.Print("Palindrom\t? ")
	if palindrom(tab, m) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}