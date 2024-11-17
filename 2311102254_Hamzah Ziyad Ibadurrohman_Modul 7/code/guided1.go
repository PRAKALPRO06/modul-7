package main

import (
	"fmt"
	"math"
)

// menyimpan titik dengan koordinat (x,y)
type titik struct {
	x int
	y int
}

// menyimpan lingkaran dengan pusat dan radius
type lingkaran struct {
	pusat  titik
	radius int
}

// fungsi untuk menghitung jarak antara dua titik
func hitungjarak(a, b titik) float64 {
	return math.Sqrt(float64((a.x - b.x) + (a.y-b.y)*(a.y-b.y)))
}

// fungsi untuk memeriksa apakah titik berada di dalam lingkaran
func titikDiDalamLingkaran(t titik, l lingkaran) bool {
	jarak := hitungjarak(t, l.pusat)
	return jarak <= float64(l.radius)
}

func main() {
	// input untuk lingkaran 1
	var cx1, cy1, r1 int
	fmt.Print(" Masukkan koordinat pusat dan radius lingkaran 1 (cx1, cy2, r1): ")
	fmt.Scanln(&cx1, &cy1, &r1)
	lingkaran1 := lingkaran{pusat: titik{x: cx1, y: cy1}, radius: r1}

	//input untuk lingkaran 2
	var cx2, cy2, r2 int
	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 2 (cx2, cy2, r2): ")
	fmt.Scanln(&cx2, &cy2, &r2)
	lingkaran2 := lingkaran{pusat: titik{x: cx2, y: cy2}, radius: r2}

	//input untuk titik sembarang
	var x, y int
	fmt.Print("Masukkan koordinat titik sembarang(x y): ")
	fmt.Scanln(&x, &y)
	titik := titik{x: x, y: y}

	// pengecekan posisi titik
	diDalam1 := titikDiDalamLingkaran(titik, lingkaran1)
	diDalam2 := titikDiDalamLingkaran(titik, lingkaran2)

	//menampilkan hasil sesuai kondisi
	if diDalam1 && diDalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik berada di luar lingkaran")
	}

}
