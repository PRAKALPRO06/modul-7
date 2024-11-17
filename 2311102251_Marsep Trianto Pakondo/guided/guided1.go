package main
import (
	"fmt"
	"math"
)

type titik struct {
	x int
	y int
}

type lingkaran struct {
	pusat titik
	radius int
}

func hitungJarak(a, b titik) float64  {
	return math.Sqrt(float64((a.x-b.x)*(a.x-b.x)+(a.y-b.y)*(a.y-b.y)))
}

func titikDidalamLingkaran(t titik, l lingkaran) bool  {
	jarak := hitungJarak(t, l.pusat)
	return jarak <= float64(l.radius)
}

func main()  {
	var cx1, cy1, r1 int

	fmt.Print("Masukkan koordinat pusat dan raduis lingkaran 1 (cx1, cy1, r1): ")
	fmt.Scanln(&cx1, &cy1, &r1)

	lingkaran1 := lingkaran{pusat: titik{x: cx1, y: cy1}, radius: r1}
	
	var cx2, cy2, r2 int

	fmt.Print("Masukkan koordinat pusat dan raduis lingkaran 1 (cx2, cy2, r2): ")
	fmt.Scanln(&cx2, &cy2, &r2)

	lingkaran2 := lingkaran{pusat: titik{x: cx2, y: cy2}, radius: r2}
	
	var x, y int

	fmt.Print("Masukkan koordinat titik sembarang (x, y): ")
	fmt.Scanln(&x, &y)

	Titik := titik{x: x, y: y}	

	diDalam1 := titikDidalamLingkaran(Titik, lingkaran1)
	diDalam2 := titikDidalamLingkaran(Titik, lingkaran2)

	if diDalam1 && diDalam2 {
		fmt.Println("Titik didalam lingkaran 1 dan 2")
	} else if diDalam1 {
		fmt.Println("Titik didalam lingkaran 1")
	} else if diDalam2 {
		fmt.Println("Titik didalam lingkaran 2")
	} else {
		fmt.Println("Titik berada di luar kedua lingkaran")
	}

}