package main
import "fmt"

func main()  {
	var klubA, klubB string
	var skorA, skorB int
	var hasil []string

	fmt.Print("Masukkan Klub A : ")
	fmt.Scan(&klubA)

	fmt.Print("Masukkan Klub B : ")
	fmt.Scan(&klubB)

	i := 1
	for {
		fmt.Printf("Pertandingan %v :", i)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		} else if skorA > skorB {
			hasil = append(hasil, klubA)
		} else if skorB > skorA {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		} 
		i++
	}

	for i := 0; i < len(hasil); i++ {
		fmt.Println("Hasil ", i+1, " : ", hasil[i])
	}
}