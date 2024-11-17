package main

import (
	"fmt"
)

type HasilPertandingan string

const (
	MENANG HasilPertandingan = "MENANG"
	SERI   HasilPertandingan = "SERI"
	KALAH  HasilPertandingan = "KALAH"
)

type Klub struct {
	nama string
	skor int
}

type RekapPertandingan struct {
	klubA      Klub
	klubB      Klub
	pemenang   []string
	jumlahData int
}

func BuatRekapPertandingan() *RekapPertandingan {
	var namaKlubA, namaKlubB string

	fmt.Print("Klub A : ")
	fmt.Scan(&namaKlubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&namaKlubB)

	return &RekapPertandingan{
		klubA:      Klub{nama: namaKlubA},
		klubB:      Klub{nama: namaKlubB},
		pemenang:   make([]string, 0),
		jumlahData: 0,
	}
}

func (r *RekapPertandingan) tentukanPemenang(skorA, skorB int) string {
	if skorA > skorB {
		return r.klubA.nama
	} else if skorB > skorA {
		return r.klubB.nama
	}
	return "Draw"
}

func (r *RekapPertandingan) prosesInputSkor() {
	nomorPertandingan := 1

	for {
		fmt.Printf("Pertandingan %d : ", nomorPertandingan)
		fmt.Scan(&r.klubA.skor, &r.klubB.skor)

		if r.klubA.skor < 0 || r.klubB.skor < 0 {
			break
		}

		hasil := r.tentukanPemenang(r.klubA.skor, r.klubB.skor)
		if hasil != "Draw" {
			r.pemenang = append(r.pemenang, hasil)
		} else {
			r.pemenang = append(r.pemenang, "Draw")
		}

		nomorPertandingan++
		r.jumlahData++
	}
}

func (r *RekapPertandingan) tampilkanRekapPemenang() {
	for i, klub := range r.pemenang {
		fmt.Printf("Hasil %d : %s\n", i+1, klub)
	}
}

func main() {
	rekap := BuatRekapPertandingan()
	rekap.prosesInputSkor()
	rekap.tampilkanRekapPemenang()
	fmt.Println("Pertandingan Selesai")
}
