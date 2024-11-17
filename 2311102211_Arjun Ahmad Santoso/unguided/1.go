package main 

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}
type Lingkaran struct {
	c Titik
	r int
}

func main() {

	var l1 Lingkaran
	var l2 Lingkaran
	var t Titik

	fmt.Print("Masukkan koordinat x dan y titik pusat dan radius lingkaran 1: ")
	fmt.Scan(&l1.c.x, &l1.c.y, &l1.r)
	fmt.Print("Masukkan koordinat x dan y titik pusat dan radius lingkaran 2: ")
	fmt.Scan(&l2.c.x, &l2.c.y, &l2.r)
	fmt.Print("Masukkan koordinat x dan y titik sembarang: ")
	fmt.Scan(&t.x, &t.y)
	
	is_in_l1 := math.Sqrt(float64((l1.c.y - t.y)*(l1.c.y - t.y)) + float64((l1.c.x - t.x)*(l1.c.x - t.x))) <= float64(l1.r)
	is_in_l2 := math.Sqrt(float64((l2.c.y - t.y)*(l2.c.y - t.y)) + float64((l2.c.x - t.x)*(l2.c.x - t.x))) <= float64(l2.r)

	if is_in_l1 && is_in_l2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if is_in_l1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if is_in_l2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}