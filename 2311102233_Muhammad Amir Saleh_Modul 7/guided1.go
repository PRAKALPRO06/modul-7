package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	center titik
	radius int
}

func jarak(p, q titik) float64 {
	dx := float64(p.x - q.x)
	dy := float64(p.y - q.y)
	return math.Sqrt(dx*dx + dy*dy)
}

func dalam(c lingkaran, p titik) bool {
	return jarak(c.center, p) <= float64(c.radius)
}

func posisi(c1, c2 lingkaran, p titik) string {
	dlmLingkaran1 := dalam(c1, p)
	dlmLingkaran2 := dalam(c2, p)

	if dlmLingkaran1 && dlmLingkaran2 {
		return "Titik di dalam lingkaran 1 dan 2"
	} else if dlmLingkaran1 {
		return "Titik di dalam lingkaran 1"
	} else if dlmLingkaran2 {
		return "Titik di dalam lingkaran 2"
	}
	return "Titik di luar lingkaran 1 dan 2"
}

func main() {
	var c1x, c1y, r1 int
	fmt.Scan(&c1x, &c1y, &r1)
	lingkaran1 := lingkaran{titik{c1x, c1y}, r1}

	var c2x, c2y, r2 int
	fmt.Scan(&c2x, &c2y, &r2)
	lingkaran2 := lingkaran{titik{c2x, c2y}, r2}

	var px, py int
	fmt.Scan(&px, &py)
	titik := titik{px, py}

	result := posisi(lingkaran1, lingkaran2, titik)
	fmt.Println(result)
}
