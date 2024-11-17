package main

import (
	"fmt"
	"math"
)

//Struktur untuk menyimpan titik dengan koordinat (x,y)
type Titik struct {
	x int
	y int
}

// Struktur untuk menyimpan lingkaran dengan pusat dan radius
type Lingkaran struct {
	pusat Titik
	radius int
}

// Fungsi unwtuk mengitung jarak titik antara dua titik
func hitungJarak(a, b Titik) float64 {
	return math.Sqrt(float64((a.x - b.x)*(a.x - b.x) + (a.y - b.y)*(a.y - b.y)))
}

// Fungsi untuk memeriksa apakah titik berada di dalam lingkaran
func titikDiDalamLingkaran(t Titik, l Lingkaran) bool {
	jarak := hitungJarak(t, l.pusat)
	return jarak <= float64(l.radius)
}

func main(){
	// Input untuk lingkaran l
	var cx1, cy1, r1 int
	fmt.Print("Masukan koordinat pusat dan radius lingkaran 1 (cx1, cy1, r1): ")
	fmt.Scanln(&cx1, &cy1, &r1)
	lingkaran1 := Lingkaran{pusat: Titik{x: cx1, y: cy1}, radius: r1}

	// Input untuk lingkaran 2
	var cx2, cy2, r2 int
	fmt.Print("Masukan koordinat pusat dan radius lingkaran 2 (cx2, cy2, r2): ")
	fmt.Scanln(&cx2, &cy2, &r2)
	lingkaran2 := Lingkaran{pusat: Titik{x: cx2, y: cy2}, radius: r2}

	var x, y int
	fmt.Print("Masukan koordinat titik sembarang (x, y): ")
	fmt.Scanln(&x, &y)
	titik := Titik{x: x, y: y}

	// Pengecekan posisi titik
	diDalam1 := titikDiDalamLingkaran(titik, lingkaran1)
	diDalam2 := titikDiDalamLingkaran(titik, lingkaran2)

	// Menampilkan hasil sesua kondisi
	if diDalam1 && diDalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar kedua lingkaran")
	}
}