package main

import (
	"fmt"
	"math"
)

// Struktur untuk menyimpan titik dengan koordinat (x, y)
type Titik struct {
	x int
	y int
}

// Struktur untuk menyimpan lingkaran dengan pusat dan radius
type Lingkaran struct {
	pusat  Titik
	radius int
}

// Fungsi untuk menghitung jarak antara dua titik
func hitungJarak(a, b Titik) float64 {
	return math.Sqrt(float64((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)))
}

// Fungsi untuk memeriksa apakah titik berada di dalam lingkaran
func titikDalamLingkaran(t Titik, l Lingkaran) bool {
	jarak := hitungJarak(t, l.pusat)
	return jarak <= float64(l.radius)
}

func main() {
	// Input untuk lingkaran 1
	var cx1, cy1, r1 int
	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 1 (cx1 cy1 r1): ")
	fmt.Scanln(&cx1, &cy1, &r1)
	lingkaran1 := Lingkaran{pusat: Titik{x: cx1, y: cy1}, radius: r1}

	// Input untuk lingkaran 2
	var cx2, cy2, r2 int
	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 2 (cx2 cy2 r2): ")
	fmt.Scanln(&cx2, &cy2, &r2)
	lingkaran2 := Lingkaran{pusat: Titik{x: cx2, y: cy2}, radius: r2}

	// Input untuk titik sembarang
	var x, y int
	fmt.Print("Masukkan koordinat titik sembarang (x y): ")
	fmt.Scanln(&x, &y)
	titik := Titik{x: x, y: y}

	// Pengecekan posisi titik
	diDalam1 := titikDalamLingkaran(titik, lingkaran1)
	diDalam2 := titikDalamLingkaran(titik, lingkaran2)

	// Menampilkan hasil sesuai kondisi
	if diDalam1 && diDalam2 {
		fmt.Println("Titik berada di dalam lingkaran 1 dan 2")
	} else if diDalam1 {
		fmt.Println("Titik berada di dalam lingkaran 1")
	} else if diDalam2 {
		fmt.Println("Titik berada di dalam lingkaran 2")
	} else {
		fmt.Println("Titik berada di luar kedua lingkaran")
	}
}
