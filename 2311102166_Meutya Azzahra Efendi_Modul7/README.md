![image](https://github.com/user-attachments/assets/b8a20d7d-d57e-41fa-a690-2da176d3ad90)<h1 align="center">LAPORAN PRAKTIKUM</h1>
<h2 align="center">ALGORITMA DAN PEMROGRAMAN 2</h2>

<h3 align="center">MODUL 7</h3>
<h4 align="center">STRUCK & ARRAY</h4>

<p align="center">
  <img src="telkomuniv.png" alt="Logo Telkom University" width="200">
</p>
<p align="center">
    <strong>Disusun Oleh:</strong><br>
    Meutya Azzahra Efendi / 2311102166<br>
</p>

<p align="center">
    <strong>Dosen Pengampu:</strong><br>
    Abednego Dwi Septiadi
</p>

<p align="center">
    PROGRAM STUDI S1 TEKNIK INFORMATIKA<br>
    FAKULTAS INFORMATIKA<br>
    TELKOM UNIVERSITY PURWOKERTO<br>
    2024
</p>


### I. DASAR TEORI

### 7.1 Tipe Bentukan
Tipe bentukan memungkinkan pemrograman untuk mendefinisikan suatu tipe data baru pada suatu bahasa pemrograman. Tipe bentukan ini dapat dibedakan atas dua jenis, yaitu Alias dan Struct.
1. Alias (Type)
Bahasa pemrograman pada umumnya mengizinkan pemrograman untuk mengubah nama suatu tipe data dengan nama baru yang lebih ringkas dan familiar. Sebagai contoh "integer" dapat dirubah dengan nama alias "bilangan". Caranya dengan menggunakan kata kunci "type".

![Screenshot 2024-11-16 111423](https://github.com/user-attachments/assets/470c617c-7c77-4330-bbb3-71badcfb2eb3)

2. Struct atau Record
Structure memungkinkan pemrograman untuk mengelompokkan beberapa data atau nilai yang memiliki relasi atau keterkaitaan tertentu menjadi suatu kesatuan. Masing-masing nilai tersimpan dalam field dan structure tersebut.

![Screenshot 2024-11-16 111713](https://github.com/user-attachments/assets/f81235e6-3b7c-4911-9c32-8aa686a99664)

Berbeda dengan bahasa pemrograman lain, kesamaan tipe dari dua variabel berjenis structure bukan karena namanya tetapai karena strukturnya. Dua variabel dengan nama-nama field dan tipe field yang sama (dan salam urutan yang sama) dianggap mempunyai tipe yang sama. Tentunya akan lebih memudahkan jika structure tersebut didefinisikan sebagai sebuah tipe baru, sehingga deklarasi structure tidak perlu lagi seluruh fieldnya ditulis ulang berkali-kali.

![Screenshot 2024-11-16 112019](https://github.com/user-attachments/assets/082520cc-4736-4bb9-b105-757bc864bde7)

### 7.2 Array
Array mempunyai ukuran (jumlah elemen) yang tetap (statis) selama eksekusi program, sehingga jumlah elemen array menjadi bagian dari deklarasi variabel dengan tipe array.

![Screenshot 2024-11-16 112440](https://github.com/user-attachments/assets/dd44d9ee-f72e-4be1-b123-0521870657e6)

Jumlah elemen array dapat diminta dengan fungsi len yang tersedia. Sebagai contoh len(arr) akan menghasilkan 73 untuk contoh di atas.
Indeks array dimulai dari 0, sehingga indeks arr pada contoh adalah 0, 1.. len(arr)-1

![Screenshot 2024-11-16 112716](https://github.com/user-attachments/assets/5389218f-6045-420f-99aa-b040b5a3d7bd)

- Slice (Array dinamik)
  
Array dalam Go juga dapat mempunyai ukuran dinamik. (Tidak digunakan di kelas algoritma pemrograman). Deklarasinya mirip dengan deklarasi array, tetapi jumlah elemennya dikosongkan.

![Screenshot 2024-11-16 112934](https://github.com/user-attachments/assets/e4a0c9bc-a9d6-42d5-b16e-3484b59fe11a)

Fungsi built-in len dapat digunakan untuk mengetahui ukuran slice. Fungsi lain, cap dapat digunakan untuk mengetahui total tempat yang disediakan untuk slice tersebut.

![Screenshot 2024-11-16 113120](https://github.com/user-attachments/assets/63a83498-f1eb-44de-8ff1-a5e4557914ff)

Fungsi built-in append dapat digunakan untuk menambahlan elemen ke suatu slice dan bila perlu memperbesar tempat untuk slice tersebut.

![Screenshot 2024-11-16 113255](https://github.com/user-attachments/assets/34f05888-04d0-4970-88d0-49da4c0e8877)

Sebuah slice baru juga dapat terbentuk dengan mengambik slice dari suatu array atau slice yang lain.

![Screenshot 2024-11-16 113409](https://github.com/user-attachments/assets/bb241370-3c94-4a54-853c-295221429dfc)

- Map
  
Tipe array lain, sebuah array dinamik, indeksnya (disini disebut kunci) tidak harus berbentuk integer. Indeks dapat berasal dari tipe apa saja. Struktur ini disebut map.

![Screenshot 2024-11-16 162539](https://github.com/user-attachments/assets/e960760f-70cd-4bd1-b0fd-f1cef9da326f)

![Screenshot 2024-11-16 162613](https://github.com/user-attachments/assets/dec6d365-4120-415a-b002-3ddfdff4333f)

### II. GUIDED

# 1. Soal study Case
Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x,y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya.

Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.

Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik diluar lingkaran 1 dan 2".

Contoh:

![Screenshot 2024-11-16 193058](https://github.com/user-attachments/assets/3b168230-9918-4ad1-a019-8ffba46ca5d9)

Fungsi untuk menghitung jarak titik (a,b) dan (c,d) dimana rumus jarak adalah:

![Screenshot 2024-11-16 193306](https://github.com/user-attachments/assets/bf3f5141-8c04-4946-86a4-b3804c1d0f46)

dan juga fungsi untuk menentukan posisi sebuah titik sembarang berada di dalam suatu lingkaran atau tidak.

![Screenshot 2024-11-16 193403](https://github.com/user-attachments/assets/dab874cb-7d58-41dd-970a-6a080d0eff5f)

# Source Code
```go
package main

import (
	"fmt"
	"math"
)

// struktur untuk menyimpan titik dengan koordinat (x, y)
type Titik struct {
	x int
	y int
}

// struktur untuk menyimpan lingkaran dengan pusat dan radius
type Lingkaran struct {
	pusat  Titik
	radius int
}

// fungsi untuk menghitung jarak antara dua titik
func hitungJarak(a, b Titik) float64 {
	return math.Sqrt(float64((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)))
}

// fungsi untuk memeriksa apakah titik berada di dalam lingkaran
func titikDiDalamLingkaran(t Titik, l Lingkaran) bool {
	jarak := hitungJarak(t, l.pusat)
	return jarak <= float64(l.radius)
}

func main() {
	//input untuk lingkaran 1
	var cx1, cy1, r1 int
	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 1 (cx1 cy1 r1): ")
	fmt.Scanln(&cx1, &cy1, &r1)
	lingkaran1 := Lingkaran{pusat: Titik{x: cx1, y: cy1}, radius: r1}

	// input untuk lingkaran 2
	var cx2, cy2, r2 int
	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 2 (cx2 cy2 r2): ")
	fmt.Scanln(&cx2, &cy2, &r2)
	lingkaran2 := Lingkaran{pusat: Titik{x: cx2, y: cy2}, radius: r2}

	// input untuk titik sembarang
	var x, y int
	fmt.Print("Masukkan koordinat titik sembarang (x y): ")
	fmt.Scanln(&x, &y)
	titik := Titik{x: x, y: y}

	// pengecekan posisi titik
	diDalamL1 := titikDiDalamLingkaran(titik, lingkaran1)
	diDalamL2 := titikDiDalamLingkaran(titik, lingkaran2)

	// menampilkan hasil sesuai kondisi
	if diDalamL1 && diDalamL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalamL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalamL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik berada di luar lingkaran")
	}
}
```
# Screenshot Output

![image](https://github.com/user-attachments/assets/9f4dab11-f9cb-4f93-afe2-e51da027f684)

# Deskripsi Program
Program di atas adalah aplikasi dalam bahasa Go yang memeriksa apakah sebuah titik berada di dalam satu atau lebih lingkaran. Pengguna memasukkan koordinat pusat dan radius untuk dua lingkaran serta koordinat sebuah titik. Program menghitung jarak titik ke pusat lingkaran menggunakan rumus Euclidean dan membandingkannya dengan radius lingkaran. Hasilnya menunjukkan apakah titik tersebut berada di dalam lingkaran pertama, lingkaran kedua, keduanya, atau di luar keduanya.

### II. UNGUIDED

# 1. Soal study Case
Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa informasi berikut:
a.  Menampilkan keseluruhan isi array
b.  Menampilkan elemen-elemen array dengan indeks ganjil saja
c.  Menampilkan elemen-elemen array dengan indeks genap saja (asumsikan indek ke-0 adalah genap)
d.  Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan pengguna.
e.  Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid. Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil
f.  Menampilkan rata-rata dari bilangan yang ada di dalam array
g.  Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut.

# Source Code
```go
// Meutya Azzahra Efendi
// 2311102166
package main

import (
	"fmt"
	"math"
)

func tampilSemua(array []int) {
	fmt.Println("Isi array:", array)
}

func tampilGanjil(array []int) {
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < len(array); i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}

func tampilGenap(array []int) {
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < len(array); i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}

func tampilKelipatan(array []int, x int) {
	fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
	for i := x; i < len(array); i += x {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}

func hapusIndeks(array []int, indeks int) []int {
	fmt.Printf("Menghapus elemen pada indeks %d\n", indeks)
	return append(array[:indeks], array[indeks+1:]...)
}

func rataRata(array []int) float64 {
	total := 0
	for _, val := range array {
		total += val
	}
	return float64(total) / float64(len(array))
}

func standarDeviasi(array []int) float64 {
	mean := rataRata(array)
	var sum float64
	for _, val := range array {
		sum += math.Pow(float64(val)-mean, 2)
	}
	return math.Sqrt(sum / float64(len(array)))
}

func frekuensi(array []int, nilai int) int {
	count := 0
	for _, val := range array {
		if val == nilai {
			count++
		}
	}
	return count
}

func main() {
	var N, x, hapusIdx, cariFrekuensi int

	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&N)

	array := make([]int, N)

	fmt.Println("Masukkan elemen array:")
	for i := 0; i < N; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}

	tampilSemua(array)
	tampilGanjil(array)
	tampilGenap(array)

	fmt.Print("Masukkan nilai x untuk menampilkan elemen dengan indeks kelipatan x: ")
	fmt.Scan(&x)
	tampilKelipatan(array, x)

	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&hapusIdx)
	if hapusIdx >= 0 && hapusIdx < len(array) {
		array = hapusIndeks(array, hapusIdx)
		tampilSemua(array)
	} else {
		fmt.Println("Indeks yang dimasukkan tidak valid!")
	}

	fmt.Printf("Rata-rata elemen array: %.2f\n", rataRata(array))

	fmt.Printf("Standar deviasi elemen array: %.2f\n", standarDeviasi(array))

	fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&cariFrekuensi)
	fmt.Printf("Frekuensi %d di dalam array: %d\n", cariFrekuensi, frekuensi(array, cariFrekuensi))
}
```

# Screenshot Output

![image](https://github.com/user-attachments/assets/d2ad9d24-ba03-46da-8dc8-8fa3df642bb2)

# Deskripsi Program
Program di atas adalah aplikasi dalam bahasa Go untuk memproses array integer dengan berbagai fungsi analitik dan manipulasi. Berikut deskripsi singkat fungsionalitasnya:

1. **Tampil Semua Elemen**: Menampilkan seluruh elemen array.  
2. **Tampil Elemen dengan Indeks Ganjil**: Menampilkan elemen array pada indeks ganjil.  
3. **Tampil Elemen dengan Indeks Genap**: Menampilkan elemen array pada indeks genap.  
4. **Tampil Elemen pada Indeks Kelipatan X**: Menampilkan elemen array pada indeks yang merupakan kelipatan bilangan \( x \).  
5. **Hapus Elemen pada Indeks Tertentu**: Menghapus elemen array berdasarkan indeks yang diberikan.  
6. **Rata-Rata Elemen**: Menghitung rata-rata dari seluruh elemen array.  
7. **Standar Deviasi**: Menghitung standar deviasi elemen array untuk mengukur sebarannya.  
8. **Frekuensi Bilangan**: Menghitung jumlah kemunculan sebuah nilai tertentu dalam array.  

Program menerima input dari pengguna untuk menentukan jumlah elemen, nilai-nilai elemen, indeks yang akan dihapus, bilangan untuk frekuensi, serta \( x \) untuk operasi kelipatan. Outputnya berupa hasil operasi sesuai fungsi yang dijalankan.

# 2. Soal Study Case
Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. BUatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga.
Pertama-ttama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja. Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.
Perhantikan sesi interaksi pada contoh berikut ini (teks bergaris bawah adalah input/read).

![image](https://github.com/user-attachments/assets/50c86098-5dd0-423c-8ec6-457fd4f7cc01)

# Source Code
```go
// Meutya Azzahra Efendi
// 2311102166

package main

import "fmt"

func inputNamaKlub() (string, string) {
	var klubA, klubB string
	fmt.Print("Masukkan nama klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Masukkan nama klub B: ")
	fmt.Scanln(&klubB)
	return klubA, klubB
}

func inputSkor(pertandingan int, klubA, klubB string) (int, int) {
	var skorA, skorB int
	fmt.Printf("Pertandingan %d antara %s dan %s: ", pertandingan, klubA, klubB)
	_, err := fmt.Scanf("%d %d\n", &skorA, &skorB)

	// cek jika input tidak valid
	if err != nil {
		fmt.Println("Input tidak valid. Pastikan untuk memasukkan dua angka.")
		// mengosongkan input buffer
		var dummy string
		fmt.Scanln(&dummy)
		return -1, -1 // mengembalikan nilai tidak valid
	}
	return skorA, skorB
}

func main() {
	klubA, klubB := inputNamaKlub()
	var hasil []string

	// loop untuk input skor
	for pertandingan := 1; pertandingan <= 9; pertandingan++ {
		skorA, skorB := inputSkor(pertandingan, klubA, klubB)

		// break loop jika salah satu skor negatif
		if skorA < 0 || skorB < 0 {
			break
		}

		// tentukan hasil pertandingan
		if skorA > skorB {
			hasil = append(hasil, klubA) // klubA menang
		} else if skorA < skorB {
			hasil = append(hasil, klubB) // klubB menang
		} else {
			hasil = append(hasil, "Draw") // seri
		}
	}

	// tampilkan hasil
	fmt.Println("Hasil pertandingan:")
	for i := 0; i < 9; i++ {
		if i < len(hasil) {
			fmt.Printf("Hasil Pertandingan %d: %s\n", i+1, hasil[i])
		} else {
			// jika belum ada hasil untuk pertandingan tersebut
			fmt.Printf("Hasil Pertandingan %d: Belum selesai\n", i+1)
		}
	}
	fmt.Println("Pertandingan selesai")
}
```

# Screenshot Output

![image](https://github.com/user-attachments/assets/e6414ad0-436c-4af0-b082-e4580ef2c5b0)

# Deskripsi Program
Program di atas adalah aplikasi dalam bahasa Go untuk mencatat hasil pertandingan antara dua klub sepak bola dalam maksimal 9 pertandingan. Berikut adalah deskripsi singkatnya:

Input Nama Klub: Pengguna memasukkan nama untuk Klub A dan Klub B.
Input Skor Pertandingan: Pengguna memasukkan skor kedua klub untuk setiap pertandingan secara berulang hingga 9 kali atau hingga skor negatif dimasukkan untuk menghentikan input.
Validasi Input: Jika input tidak valid (bukan angka), program meminta pengguna untuk mengulangi input skor.
Penentuan Hasil:
Jika skor Klub A lebih besar, Klub A dinyatakan menang.
Jika skor Klub B lebih besar, Klub B dinyatakan menang.
Jika skor sama, hasilnya dinyatakan sebagai "Draw".
Output Hasil Pertandingan: Program menampilkan hasil dari setiap pertandingan yang telah tercatat. Jika kurang dari 9 pertandingan selesai, sisanya akan ditampilkan sebagai "Belum selesai".
Fitur Utama:

Mencatat hasil hingga 9 pertandingan.
Memvalidasi input agar hanya menerima skor berupa angka.
Menghentikan input jika ada skor negatif.

# 3. Soal Study Case
Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan mebalikkan urutan isi array dan memeriksa apakah membentuk palindrom.

Lengkapi potongan algoritma berikut ini!
```go
package main
import "fmt"

const NMAX int = 127
type tabel [NMAX] rune
  tab: tabel
  m: integer

func isiArray(t *tabel, n *int)
/*I.S. Data tersedia dalam piranti masukan
  F.S. Array t berisi sejumlah n karakter yang dimasukkan user,
  Proses input selama karakter bukanlah TITIK dan n <= NMAX */
func cetakArray(t tabel, n int)
/*I.S. Terdefinisi array t yang berisi sejumlah n karakter
  F.S. n karakter dalam array muncul di layar */
func balikanArray(t *tabel, n int)
/*I.S. Terdefinisi array t yang berisi sejumlah n karakter
  F.S. Urutan isi array t terbalik */
func main(){
    var tab tabel
    var m int

    // Isi array tab dengan memanggil prosedur isiArray

    // Balikan isi array tab dengan memanggil balikanArray

    // Cetak isi array tab
}
```

![image](https://github.com/user-attachments/assets/10918798-f669-4e10-8dd4-6cf825081fd5)

![image](https://github.com/user-attachments/assets/47edac33-b30c-45d8-b4e4-44c308a42ffb)

Perhatikan sesi interaksi pada contoh berikut ini (teks bergaris bawah adalah input/read)

![image](https://github.com/user-attachments/assets/a5246f74-c38b-42b6-9d50-1f00e0e416a3)

Modifikasi program tersebut dengan menambahkan fungsi palindrom. Tambahkan instruksi untuk memanggil fungsi tersebut dan menampilkan hasilnya pada program utama.

![image](https://github.com/user-attachments/assets/d5e68679-8123-4482-9bb6-060deb293873)

# Source Code
```go
// Meutya Azzahra Efendi
// 2311102166

package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

// Fungsi untuk membaca input karakter hingga titik
func bacaInputKarakter(t *tabel) int {
	var input rune
	var n int
	fmt.Println("Masukkan karakter-karakter (akhir dengan titik '.'): ")
	for {
		fmt.Scanf("%c", &input)
		if input == '.' { // titik sebagai penghenti input
			break
		}
		if n < NMAX {
			(*t)[n] = input
			n++
		}
	}
	return n
}

// Prosedur untuk mencetak array
func cetakArray(t tabel, n int) {
	fmt.Print("Teks: ")
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]))
	}
	fmt.Println()
}

// Prosedur untuk membalikkan urutan array
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

// Fungsi untuk memeriksa apakah array membentuk palindrom
func adalahPalindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel

	// Isi array tab dengan memanggil fungsi bacaInputKarakter
	m := bacaInputKarakter(&tab)

	// Cetak array tab sebelum dibalik
	cetakArray(tab, m)

	// Balikan isi array tab
	balikanArray(&tab, m)

	// Cetak array tab setelah dibalik
	fmt.Print("Reverse teks: ")
	cetakArray(tab, m)

	// Cek apakah array membentuk palindrom
	isPalindrom := adalahPalindrom(tab, m)
	fmt.Printf("Palindrom: %v\n", isPalindrom)
}
```

# Screenshot Output

![image](https://github.com/user-attachments/assets/12f1d043-30dd-4add-845a-f385d1f11ee0)

# Deskripsi Program
Program di atas adalah aplikasi dalam bahasa Go untuk membaca, memproses, dan menganalisis string yang dimasukkan pengguna hingga tanda titik (.). Berikut adalah deskripsi singkatnya:

Membaca Input: Program membaca karakter satu per satu dari pengguna hingga tanda titik (.) dimasukkan. Karakter disimpan dalam array tabel.

Menampilkan String: Program menampilkan string yang telah dimasukkan.

Membalik String: Program membalik urutan karakter dalam string menggunakan prosedur balikanArray.

Memeriksa Palindrom: Program memeriksa apakah string merupakan palindrom, yaitu string yang tetap sama ketika dibaca terbalik.

Fitur Utama:

Input Karakter Dinamis: Membaca hingga 127 karakter sebelum tanda titik.
Balik String: Membalik string dan menampilkan hasilnya.
Pemeriksaan Palindrom: Menentukan apakah string tersebut simetris.
