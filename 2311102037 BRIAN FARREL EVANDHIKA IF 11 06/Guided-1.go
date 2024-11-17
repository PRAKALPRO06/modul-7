//2311102037_BRIAN FARREL EVANDHIKA_IF 11 06
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
	pusat  titik
	radius int
}

func hitungjarak(a, b titik) float64 {
	return math.Sqrt(float64((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)))
}

func titikdalamlingkaran(t titik, l lingkaran) bool {
	jarak := hitungjarak(t, l.pusat)
	return jarak <= float64(l.radius)
}

func main() {
	var cx1, cy1, r1 int
	fmt.Print("Masukkan kordinat pusat dan radius lingkaran l (cx1 cy1 r1): ")
	fmt.Scanln(&cx1, &cy1, &r1)

	// Membuat objek lingkaran1
	lingkaran1 := lingkaran{pusat: titik{x: cx1, y: cy1}, radius: r1}

	var cx2, cy2, r2 int
	fmt.Print("Masukkan kordinat pusat dan radius lingkaran l2 (cx2 cy2 r2): ")
	fmt.Scanln(&cx2, &cy2, &r2)

	// Membuat objek lingkaran2
	lingkaran2 := lingkaran{pusat: titik{x: cx2, y: cy2}, radius: r2}

	var x, y int
	fmt.Print("Masukkan kordinat titik t (x y): ")
	fmt.Scanln(&x, &y)

	// Membuat objek titik yang akan diperiksa
	titik2 := titik{x: x, y: y}

	// Memeriksa apakah titik berada di dalam lingkaran1 dan lingkaran2
	didalaml1 := titikdalamlingkaran(titik2, lingkaran1)
	didalaml2 := titikdalamlingkaran(titik2, lingkaran2)

	if didalaml1 && didalaml2 {
		fmt.Println("Titik berada di dalam kedua lingkaran.")
	} else if didalaml1 {
		fmt.Println("Titik berada di dalam lingkaran 1 saja.")
	} else if didalaml2 {
		fmt.Println("Titik berada di dalam lingkaran 2 saja.")
	} else {
		fmt.Println("Titik berada di luar kedua lingkaran.")
	}
}