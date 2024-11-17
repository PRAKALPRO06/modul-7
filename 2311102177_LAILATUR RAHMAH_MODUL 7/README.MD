<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong> MODUL VII </strong></h2>
<h2 align="center"><strong> STRUCT DAN ARRAY </strong></h2>

<br>

<p align="center">

  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  LAILATUR RAHMAH / 2311102177<br>
  S1 IF 11 06
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Abednego Dwi Septiadi
  
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong><br>
  <strong>FAKULTAS INFORMATIKA</strong><br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong><br>
  <strong>2024</strong>
</p>

------

## Features

- [DASAR TEORI](#dasar-teori)
- [GUIDED](#guided)
- [UNGUIDED](#unguided)

## DASAR TEORI

### Array

Array adalah tipe data dasar yang digunakan untuk menyimpan sekumpulan elemen dengan tipe data yang sama. Di bahasa Go, array memiliki ukuran tetap yang ditentukan saat deklarasi.

**Karakteristik Array:**

1. Ukuran Tetap:

   Ukuran array di Go ditentukan saat deklarasi dan tidak dapat diubah selama runtime. Ini membuat array kurang fleksibel dibandingkan tipe data lain seperti slice.

2. Tipe Data Homogen:

   Semua elemen dalam array harus memiliki tipe data yang sama, seperti int, string, atau float64.

3. Indeks Dimulai dari Nol:

   Elemen pertama array berada pada indeks 0, elemen kedua di indeks 1, dan seterusnya.

4. Pengaksesan Elemen:

   Elemen dalam array diakses menggunakan indeks dengan notasi array[index].

**Deklarasi Array:**

Array dapat dideklarasikan dengan menentukan ukuran dan tipe datanya. Ada beberapa cara untuk mendeklarasikan dan menginisialisasi array di Go.

- Deklarasi Array Kosong:

```go
var numbers [5]int // array of 5 integers, initialized to zero values
```

- Inisialisasi Array dengan Nilai:

```go
var primes = [5]int{2, 3, 5, 7, 11} // array dengan panjang 5 dan nilai yang ditentukan
```

- Deklarasi Array dengan Panjang Otomatis:

Go juga memungkinkan kita untuk menentukan array tanpa panjang yang tetap dengan menggunakan [...], yang akan menyesuaikan panjang array sesuai dengan jumlah elemen yang diinisialisasi.

```go
primes := [...]int{2, 3, 5, 7, 11} // array panjang otomatis sesuai jumlah elemen
```

- Penggunaan Array:

Array dapat diakses dan dimodifikasi dengan mengacu pada indeksnya. 

```go
var numbers = [3]int{1, 2, 3}
fmt.Println(numbers[0]) // Output: 1
numbers[0] = 10         // Mengubah nilai elemen pertama menjadi 10
fmt.Println(numbers[0]) // Output: 10
```

- Array Multidimensi:

Go juga mendukung array multidimensi, yang berguna untuk representasi data seperti matriks atau tabel.

```go
var matrix [3][3]int // array 2D dengan 3 baris dan 3 kolom
matrix[0][1] = 5     // menetapkan nilai pada baris 0, kolom 1
```


### Struck

Struct adalah tipe data komposit yang memungkinkan kita mengelompokkan variabel dengan tipe yang berbeda dalam satu entitas. Struct sering digunakan untuk merepresentasikan objek atau entitas yang memiliki beberapa atribut.

**Karakteristik Struct:**
1. Tipe Data Heterogen:

   Field dalam struct dapat memiliki tipe data yang berbeda, memungkinkan penggunaan tipe data yang lebih kompleks.

2. Nama Field:

   Setiap field dalam struct memiliki nama yang unik.

3. Notasi Titik:

   Untuk mengakses atau memodifikasi field, kita menggunakan notasi titik (.).

4. Komposisi Struct:

   Struct mendukung komposisi (embedding), yang memungkinkan struct mengandung struct lain di dalamnya.

- Deklarasi Struct:

Struct dideklarasikan menggunakan kata kunci type, diikuti dengan nama struct, dan diakhiri dengan definisi field di dalam kurung kurawal.

```go
type Person struct {
    Name string
    Age  int
    City string
}
```

Dalam contoh ini, struct Person memiliki tiga field: Name dengan tipe string, Age dengan tipe int, dan City dengan tipe string.

**Inisialisasi Struct:**

- Menggunakan Literals:

```go
var person1 = Person{"Alice", 30, "Jakarta"}
```

- Menggunakan Nama Field (Pengisian Tertentu):

  Ini berguna jika kita ingin mengisi hanya sebagian field dari struct.

```go
person2 := Person{Name: "Bob", City: "Bandung"}
```

- Pointer Struct:

  Go mendukung penggunaan pointer pada struct, yang memungkinkan kita mengubah field struct melalui referensinya.

```go
person3 := &Person{"Charlie", 25, "Surabaya"}
person3.City = "Malang" // perubahan langsung pada alamat memori
```

- Mengakses Field Struct:
  
Field dalam struct diakses menggunakan notasi titik (.).

```go
fmt.Println("Nama:", person1.Name)
fmt.Println("Usia:", person1.Age)
```

- Embedded Structs (Komposisi Struct):

Go mendukung konsep embedding, yang memungkinkan struct memiliki struct lain di dalamnya.

```go
type Address struct {
    City, State string
}

type Person struct {
    Name    string
    Age     int
    Address // embedded struct
}

func main() {
    p := Person{Name: "Alice", Age: 30, Address: Address{City: "Jakarta", State: "DKI"}}
    fmt.Println("Kota:", p.City) // akses langsung ke field dari embedded struct
}
```

Dalam contoh di atas, Person memiliki embedded struct Address, sehingga City dan State dari Address dapat diakses langsung melalui objek Person.

### Perbandingan antara array dan struct:

| Fitur           | Array                                              | Struct                                              |
|-----------------|----------------------------------------------------|-----------------------------------------------------|
| Tipe Data       | Homogen, elemen memiliki tipe data yang sama       | Heterogen, field bisa memiliki tipe data berbeda    |
| Ukuran          | Tetap, ditentukan pada saat deklarasi	             | Dapat berubah, field dapat ditambah/diubah          |
| Penggunaan      | Menyimpan sekumpulan data linear atau sekuensial   | Merepresentasikan objek atau entitas yang kompleks  |
| Akses           | Menggunakan indeks                                 | Menggunakan nama field                              |
| Mutidimensi     | Mendukung array multidimensi                       | Mendukung komposisi atau embedding                  |

## GUIDED

### 1. Studi Case Guided 1

Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya.

Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.

Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".

**Contoh**
| No | Masukan                                                                | Keluaran                          |
|----|---------------------------------------------------------------------|---------------------------------|
| 1 | 1 1 5 <br> 8 8 4 <br> 2 2                                      | Titik di dalam lingkaran 1        |
| 2 | 1 2 3 <br> 4 5 6 <br> 7 8                                      | Titik di dalam lingkaran 2        |
| 3 | 5 10 15 <br> -15 4 20 <br> 0 0                                    | Titik di dalam lingkaran 1 dan 2  |
| 4 | 1 1 5 <br> 8 8 4 <br> 15 20                                    | Titik di luar lingkaran 1 dan 2 |

Fungsi untuk menghitung jarak titik (a, b) dan (c, d), dimana rumus jarak adalah:

```jarak = √((a-c)² + (b-d)²)```

dan juga fungsi untuk menentukan posisi sebuah titik sembarang berada di dalam suatu lingkaran atau tidak.

```go
function jarak(p, q : titik) -> real
{
    {Mengembalikan jarak antara titik p(x,y) dan titik q(x,y)}
}

function didalam(c:lingkaran, p:titik) -> boolean
{
    {Mengembalikan true apabila titik p(x,y) berada di dalam lingkaran c yang memiliki titik pusat (cx,cy) dan radius r}
}
```

**Catatan:** Lihat paket math dalam lampiran untuk menggunakan fungsi math.Sqrt() untuk menghitung akar kuadrat.

#### Source Code

```go
// LAILATUR RAHMAH
// 2311102177
// IF 11 06

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

// structur untuk menyimpan lingkaran dengan pusat dan radius
type Lingkaran struct {
	pusat  Titik
	radius int
}

// fungsi untuk mengitung jarak antara dua titik
func hitungJarak(a, b Titik) float64 {
	return math.Sqrt(float64((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)))
}

// fungsi untuk memeriksa apakah titik berada du dalam lingkaran
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

	/// input untuk titik sembarang
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

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/70bc4179-f9b1-4c35-9793-938c24de5cf9)

#### Deskripsi Program

Program ini mengimplementasikan dua buah lingkaran dan sebuah titik sembarang, kemudian memeriksa apakah titik tersebut berada di dalam salah satu atau kedua lingkaran. Struktur data yang digunakan terdiri dari Titik, yang menyimpan koordinat (x, y), dan Lingkaran, yang menyimpan pusat dan radius lingkaran. Program ini menerima input dari pengguna untuk menentukan pusat dan radius kedua lingkaran serta koordinat titik sembarang. Selanjutnya, program menghitung jarak antara titik dan pusat lingkaran menggunakan rumus jarak Euclidean, lalu memeriksa apakah jarak tersebut lebih kecil atau sama dengan radius lingkaran. Berdasarkan hasil perbandingan, program akan menampilkan apakah titik berada di dalam salah satu atau kedua lingkaran, atau di luar keduanya.

## Unguided

### 1. Studi Case Unguided 1

Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa informasi berikut:

  a. Menampilkan keseluruhan isi dari array.

  b. Menampilkan elemen-elemen array dengan indeks ganjil saja.

  c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indeks ke-0 adalah genap).

  d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan pengguna.

  e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid. Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil.

  f. Menampilkan rata-rata dari bilangan yang ada di dalam array.

  g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut.

  h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi tersebut.

#### Source Code

```go
// LAILATUR RAHMAH
// 2311102177
// IF 11 06

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

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/7c1ff33b-cebb-4629-b005-68ed53e46478)

#### Deskripsi Program

Program ini memungkinkan pengguna untuk melakukan berbagai operasi terhadap sebuah array bilangan integer yang dimasukkan melalui input. Pengguna dapat memasukkan jumlah elemen array dan kemudian mengisi array dengan elemen-elemen yang diinginkan. Program ini memiliki berbagai fungsi seperti menampilkan seluruh isi array, menampilkan elemen dengan indeks ganjil atau genap, menampilkan elemen berdasarkan kelipatan indeks tertentu, menghapus elemen pada indeks yang ditentukan, serta menghitung rata-rata dan standar deviasi dari elemen-elemen array. Selain itu, program ini juga dapat menghitung frekuensi kemunculan suatu bilangan dalam array. Hasil dari setiap operasi akan ditampilkan di layar sesuai dengan permintaan pengguna.

### 2. Studi Case Unguided 2

Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga.

Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja.

Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.   

Perhatikan sesi interaksi pada contoh berikut ini (teks bergaris bawah adalah Input/read):

```
Klub A : MU

Klub B : Inter

Pertandingan 1 : 2  0  // MU = 2 sedangkan Inter = 0
Pertandingan 2 : 1  2
Pertandingan 3 : 2  2
Pertandingan 4 : 0  1
Pertandingan 5 : 3  2
Pertandingan 6 : 1  0
Pertandingan 7 : 5  2
Pertandingan 8 : 2  3
Pertandingan 9 : -1  2

Hasil 1 : MU
Hasil 2 : Inter
Hasil 3 : Draw
Hasil 4 : Inter
Hasil 5 : MU
Hasil 6 : MU
Hasil 7 : MU
Hasil 8 : Inter

Pertandingan selesai
```

#### Source Code

```go
// LAILATUR RAHMAH
// 2311102177
// IF 11 06

package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil []string

	// input nama klub
	fmt.Print("Masukkan nama klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Masukkan nama klub B: ")
	fmt.Scanln(&klubB)

	// loop untuk input skor
	for pertandingan := 1; pertandingan <= 9; pertandingan++ {
		fmt.Printf("Pertandingan %d antara %s dan %s : ", pertandingan, klubA, klubB)
		_, err := fmt.Scanf("%d %d\n", &skorA, &skorB)

		// cek jika input tidak valid
		if err != nil {
			fmt.Println("Input tidak valid. Pastikan untuk memasukkan dua angka.")
			// mengosongkan input buffer
			var dummy string
			fmt.Scanln(&dummy)
			continue
		}

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

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/72a71bc6-4285-4919-97ef-901b5a38dfdb)

#### Deskripsi Program

Program ini digunakan untuk mencatat hasil pertandingan antara dua klub sepak bola yang dapat diinput oleh pengguna. Pengguna diminta untuk memasukkan nama klub A dan klub B, kemudian program akan melakukan input skor untuk masing-masing pertandingan antara kedua klub tersebut sebanyak 9 kali. Program juga memeriksa apakah input skor valid (dua angka yang dipisahkan oleh spasi) dan akan keluar dari loop jika salah satu skor yang dimasukkan adalah negatif. Setelah semua pertandingan selesai, hasil pertandingan akan ditampilkan. Jika pertandingan belum selesai, maka hasilnya akan ditampilkan sebagai "Belum selesai". Program ini juga menangani input yang tidak valid dengan meminta pengguna mengulangi input jika terjadi kesalahan. Pada akhir program, ditampilkan hasil dari setiap pertandingan, baik itu kemenangan klub A, klub B, atau hasil seri.

### 3. Studi Case Unguided 3

Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom.

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

Perhatikan sesi interaksi pada contoh berikut ini **(teks bergaris bawah adalah Input/read)**

Teks: **SENANG.**
Reverse teks: .GNANES

Teks: **KATAK.**
Reverse teks: KATAK.

Modifikasi program tersebut dengan menambahkan fungsi palindrom. Tambahkan instruksi untuk memanggil fungsi tersebut dan menampilkan hasilnya pada program utama.

**Palindrom adalah teks yang dibaca dari awal atau akhir adalah sama, contoh: KATAK, APA, KASUR_RUSAK**

```go
func palindrom(t tabel, n int) bool
/* Mengembalikan true apabila susunan karakter di dalam t membentuk palindrom,
dan false apabila sebaliknya. Petunjuk: Manfaatkan prosedur balikanArray */
```

Perhatikan sesi interaksi pada contoh berikut ini **(teks bergaris bawah adalah Input/read)**

Teks: **KATAK**
Palindrom: ? true

Teks: **SENANG**
Palindrom: ? false

#### Source Code

```go
// LAILATUR RAHMAH
// 2311102177
// IF 11 06

package main
import "fmt"
const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input rune
	fmt.Println("Masukkan karakter-karakter (akhir dengan titik '.'): ")
	for {
		fmt.Scanf("%c", &input)
		if input == '.' { // titik sebagai penghenti input
			break
		}
		(*t)[*n] = input
		*n++
	}
}

// prosedur untuk mencetak array
func cetakArray(t tabel, n int) {
	fmt.Print("Teks: ")
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]))
	}
	fmt.Println()
}

// prosedur untuk membalikkan urutan array
func balikanArray(t *tabel, n int) {
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
}

// fungsi untuk memeriksa apakah array membentuk palindrom
func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	// isi array tab dengan memanggil prosedur isiArray
	isiArray(&tab, &m)

	// cetak array tab sebelum dibalik
	cetakArray(tab, m)

	// balikan isi array tab dengan memanggil balikanArray
	balikanArray(&tab, m)

	// cetak array tab setelah dibalik
	fmt.Print("Reverse teks: ")
	cetakArray(tab, m)

	// Ccek apakah array membentuk palindrom
	isPalindrom := palindrom(tab, m)
	fmt.Printf("Palindrom: %v\n", isPalindrom)
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/920f0dca-400a-4ab4-8d55-1d64b5654367)

#### Deskripsi Program

Program ini memungkinkan pengguna untuk memasukkan serangkaian karakter yang disimpan dalam sebuah array, kemudian membalikkan urutan karakter tersebut dan memeriksa apakah urutan karakter tersebut membentuk sebuah palindrom. Pengguna diminta untuk menginput karakter satu per satu, dengan titik ('.') sebagai penanda akhir input. Setelah karakter dimasukkan, program akan menampilkan urutan asli, membalikkan urutannya, dan menampilkan hasilnya. Selain itu, program juga memeriksa apakah urutan karakter awal membentuk palindrom, yaitu urutan yang tetap sama saat dibaca dari depan dan belakang, dan menampilkan hasilnya sebagai true atau false.
