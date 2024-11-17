//2311102018_Aryo_Tegar_Sukarno
package main
import (
	"fmt"
	"math"
)
type Circle struct {
	x, y, radius int
}
func calculateDistance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
}
func isInsideCircle(circle Circle, px, py int) bool {
	return calculateDistance(circle.x, circle.y, px, py) <= float64(circle.radius)
}

func main() {
	
	var circle1, circle2 Circle
	var px, py int

	fmt.Println("Masukkan koordinat dan radius lingkaran 1 (x y r):")
	fmt.Scan(&circle1.x, &circle1.y, &circle1.radius)

	fmt.Println("Masukkan koordinat dan radius lingkaran 2 (x y r):")
	fmt.Scan(&circle2.x, &circle2.y, &circle2.radius)

	fmt.Println("Masukkan koordinat titik sembarang (x y):")
	fmt.Scan(&px, &py)

	insideCircle1 := isInsideCircle(circle1, px, py)
	insideCircle2 := isInsideCircle(circle2, px, py)

	if insideCircle1 && insideCircle2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if insideCircle1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if insideCircle2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
