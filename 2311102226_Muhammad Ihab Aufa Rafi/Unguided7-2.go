package main

import (
	"fmt"
	"math"
)

func tampilArray(arr []int, n int) {
    fmt.Print("Isi array: ")
    for i := 0; i < n; i++ {
        fmt.Printf("%d ", arr[i])
    }
    fmt.Println()
}

func tampilIndeksGanjil(arr []int, n int) {
    fmt.Print("Elemen dengan indeks ganjil: ")
    for i := 1; i < n; i += 2 {
        fmt.Printf("%d ", arr[i])
    }
    fmt.Println()
}

func tampilIndeksGenap(arr []int, n int) {
    fmt.Print("Elemen dengan indeks genap: ")
    for i := 0; i < n; i += 2 {
        fmt.Printf("%d ", arr[i])
    }
    fmt.Println()
}

func tampilKelipatanX(arr []int, n int, x int) {
    fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
    for i := 0; i < n; i++ {
        if i%x == 0 {
            fmt.Printf("%d ", arr[i])
        }
    }
    fmt.Println()
}

func hapusElemen(arr []int, n *int, indeks int) {
    for i := indeks; i < *n-1; i++ {
        arr[i] = arr[i+1]
    }
    *n--
}

func hitungRataRata(arr []int, n int) float64 {
    sum := 0
    for i := 0; i < n; i++ {
        sum += arr[i]
    }
    return float64(sum) / float64(n)
}

func hitungStandarDeviasi(arr []int, n int) float64 {
    mean := hitungRataRata(arr, n)
    sumSqDiff := 0.0
    
    for i := 0; i < n; i++ {
        diff := float64(arr[i]) - mean
        sumSqDiff += diff * diff
    }
    return math.Sqrt(sumSqDiff / float64(n))
}

func hitungFrekuensi(arr []int, n int, bilangan int) int {
    frekuensi := 0
    for i := 0; i < n; i++ {
        if arr[i] == bilangan {
            frekuensi++
        }
    }
    return frekuensi
}

func main() {
    var n, x, indeksHapus, cariBilangan int
    var arr [100]int
    
    fmt.Print("Masukkan jumlah elemen array (N): ")
    fmt.Scan(&n)
    
    fmt.Println("Masukkan elemen array:")
    for i := 0; i < n; i++ {
        fmt.Printf("Elemen ke-%d: ", i)
        fmt.Scan(&arr[i])
    }

    tampilArray(arr[:], n)
    tampilIndeksGanjil(arr[:], n)
    tampilIndeksGenap(arr[:], n)

    fmt.Print("Masukkan nilai X untuk menampilkan kelipatan: ")
    fmt.Scan(&x)
    tampilKelipatanX(arr[:], n, x)

    fmt.Print("Masukkan indeks yang ingin dihapus: ")
    fmt.Scan(&indeksHapus)
    hapusElemen(arr[:], &n, indeksHapus)
	
    fmt.Println("Array setelah penghapusan:")
    tampilArray(arr[:], n)
    
    fmt.Printf("Rata-rata: %.2f\n", hitungRataRata(arr[:], n))
    
    fmt.Printf("Standar deviasi: %.2f\n", hitungStandarDeviasi(arr[:], n))
    
    fmt.Print("Masukkan bilangan yang ingin dihitung frekuensinya: ")
    fmt.Scan(&cariBilangan)
    frek := hitungFrekuensi(arr[:], n, cariBilangan)
    fmt.Printf("Frekuensi bilangan %d: %d\n", cariBilangan, frek)
}