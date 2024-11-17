package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x int
	y int
}

type Lingkaran struct {
	pusat Titik
	radius int
}

func hitungJarak(a, b Titik) float64 {
	return math.Sqrt(float64 ((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)))
}

func titikDidalamLingkaran(t Titik, l Lingkaran) bool {
	jarak := hitungJarak(t, l.pusat)
	return jarak <= float64(l.radius)
}

func main(){
	var cx1, cy1, r1 int
	fmt.Print("Masukan koordinat pusat dan radius lingkaran 1 (cx1 cv1 r1): ")
	fmt.Scanln(&cx1, &cy1, &r1)
	lingkaran1 := Lingkaran{pusat: Titik{x: cx1, y: cy1}, radius: r1}

	var cx2, cy2, r2 int
	fmt.Print("Masukan koordinat pusat dan radius lingkaran 1 (cx1 cv1 r1): ")
	fmt.Scanln(&cx2, &cy2, &r2)
	lingkaran2 := Lingkaran{pusat: Titik{x: cx2, y: cy2}, radius: r2}

	var x, y int
	fmt.Print("Masukan koordinat titik sembarang (x y): ")
	fmt.Scanln(&x, &y)
	titik := Titik{x: x, y: y}

	diDalamL1 := titikDidalamLingkaran(titik, lingkaran1)
	diDalamL2 := titikDidalamLingkaran(titik, lingkaran2)

	if diDalamL1 && diDalamL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalamL1{
		fmt.Println("Titik didalam lingkaran 1")
	} else if diDalamL2 {
		fmt.Println("Titik didalam lingkaran 2")
	} else {
		fmt.Println("Titik berada di luar kedua lingkaran")
	}
}