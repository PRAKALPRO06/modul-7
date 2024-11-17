package main

import "fmt"

func main() {
    var klubA, klubB string
    var skorA, skorB int
    var hasil_226 []string
    
    fmt.Print("Klub A : ")
    fmt.Scan(&klubA)
    fmt.Print("Klub B : ")
    fmt.Scan(&klubB)
    
    for i := 1; ; i++ {
        fmt.Printf("Pertandingan %d : ", i)
        fmt.Scan(&skorA, &skorB)
        
        if skorA < 0 || skorB < 0 {
            break
        }
        
        if skorA > skorB {
            hasil_226 = append(hasil_226, klubA)
        } else if skorB > skorA {
            hasil_226 = append(hasil_226, klubB)
        } else {
            hasil_226 = append(hasil_226, "Draw")
        }
    }
    
    fmt.Println()
    for i, h := range hasil_226 {
        fmt.Printf("Hasil %d : %s\n", i+1, h)
    }
    fmt.Println("Pertandingan selesai")
}