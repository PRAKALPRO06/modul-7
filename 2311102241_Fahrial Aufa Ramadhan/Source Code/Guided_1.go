package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	pusat  titik
	radius int
}

func hitungjarak(a, b titik) float64 {
	return math.Sqrt(float64((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)))
}

func titikdalamlingkaran(t titik, I lingkaran) bool {
	jarak := hitungjarak(t, I.pusat)
	return jarak <= float64(I.radius)
}

func main() {
	var cx1, cy1, r1 int
	fmt.Print("Masukan kordinat pusat dan radius lingkaran 1 (cx1 cy1 r1): ")
	fmt.Scanln(&cx1, &cy1, &r1)
	lingkaran1 := lingkaran{pusat: titik{x: cx1, y: cy1}, radius: r1}

	var cx2, cy2, r2 int
	fmt.Print("Masukan kordinat pusat dan radius lingkaran 2 (cx2,cy2 ,r2): ")
	fmt.Scanln(&cx2, &cy2, &r2)
	lingkaran2 := lingkaran{pusat: titik{x: cx2, y: cy2}, radius: r2}

	var x, y int
	fmt.Print("Masukan kordinat titik sembarang (x,y): ")
	fmt.Scanln(&x, &y)
	titik := titik{x: x, y: y}

	diDalamL1 := titikdalamlingkaran(titik, lingkaran1)
	diDalamL2 := titikdalamlingkaran(titik, lingkaran2)

	if diDalamL1 && diDalamL2 {
		fmt.Println("titik dalam lingkaran 1 dan 2")
	} else if diDalamL1 {
		fmt.Println("titik dalam lingkaran 1")
	} else if diDalamL2 {
		fmt.Println("titik dalam lingkaran 2")
	} else {
		fmt.Println("titik di luar lingkaran")
	}

}
