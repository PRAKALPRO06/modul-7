package main

import (
	"fmt"
	"math"
)

//struktur koordinat
type Titik struct {
	x int
	y int
}
//struktur lingkaran
type Lingkaran226 struct {
	pusat Titik
	radius int
}

//fungsi euclidean
func hitungJarak(a,b Titik) float64 {
	return math.Sqrt(float64((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)))
}

func titikDalamLingkaran(t Titik, l Lingkaran226) bool {
	jarak := hitungJarak(t, l.pusat)
	return jarak <= float64(l.radius)
}

func main() {
	//input untuk lingkaran 1
	var cx1, cy1, r1 int
	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 1 (cx1, cy1, r1): ")
	fmt.Scanln(&cx1, &cy1, &r1)
	lingkaran1 := Lingkaran226{pusat: Titik{x: cx1, y:cy1}, radius: r1}

	//input untuk lingkaran 2
	var cx2, cy2, r2 int
	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 2 (cx2, cy2, r2): ")
	fmt.Scanln(&cx2, &cy2, &r2)
	lingkaran2 := Lingkaran226{pusat: Titik{x: cx2, y:cy2}, radius: r2}

	//input titik sembarang
	var x, y int
	fmt.Print("Masukkan koordinat titik sembarang (x y): ")
	fmt.Scanln(&x, &y)
	titik := Titik{x: x, y: y}

	//pengecekan posisi titik
	diDalamL1 := titikDalamLingkaran(titik, lingkaran1)
	diDalamL2 := titikDalamLingkaran(titik, lingkaran2)

	//menampilkan hasil sesuai kondisi
	if diDalamL1 && diDalamL2{
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	}else if diDalamL1{
		fmt.Println("Titik di dalam lingkaran 1")
	}else if diDalamL2{
		fmt.Println("Titik di dalam lingkaran 2")
	}else {
		fmt.Println("Titik berada di luar kedua lingkaran")
	}
}
