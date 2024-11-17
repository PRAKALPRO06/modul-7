package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]int

func isiArray(t *tabel, n *int) {
    var kar int
    *n = 0
    fmt.Print("Teks         : ")
    fmt.Scanf("%c", &kar)
    for kar != '.' && *n < NMAX {
        if kar != ' ' {
            t[*n] = kar
            *n++
        }
        fmt.Scanf("%c", &kar)
    }
}

func cetakArray(t tabel, n int, label string) {
    fmt.Print(label, " : ")
    for i := 0; i < n; i++ {
        fmt.Printf("%c", t[i])
    }
    fmt.Println()
}

func balikkanArray(t *tabel, n int) {
    for i := 0; i < n/2; i++ {
        t[i], t[n-1-i] = t[n-1-i], t[i]
    }
}

func palindrom(t tabel, n int) bool {
    for i := 0; i < n/2; i++ {
        if t[i] != t[n-1-i] {
            return false
        }
    }
    return true
}

func main() {
    var tab_226 tabel
    var m_226 int

    isiArray(&tab_226, &m_226)
        
    var tabReverse tabel = tab_226
    balikkanArray(&tabReverse, m_226)
    
    fmt.Printf("Palindrom    ? %t\n", palindrom(tab_226, m_226))
}