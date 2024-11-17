package main 

import "fmt"

func main() {
	var klubA, klubB string
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	hasil := []string{}
	i := 1
	for {
		var skorKlubA, skorKlubB int
		fmt.Print("Pertandingan ", i, " : ")
		fmt.Scan(&skorKlubA, &skorKlubB)
		if (skorKlubA < 0 || skorKlubB < 0) {
			break
		} else if (skorKlubA > skorKlubB) {
			hasil = append(hasil, klubA)
		} else if (skorKlubA < skorKlubB) {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		}
		i++
	}
	for i:=0; i<len(hasil); i++ {
		fmt.Println("Hasil ", i+1, " : ", hasil[i])
	}
}