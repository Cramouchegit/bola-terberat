# Deskripsi Program:

Program ini dibuat untuk menemukan bola terberat di antara 4 buah bola yang diberikan oleh pengguna. 
Setiap bola memiliki berat yang berbeda dan tidak diketahui sebelumnya. 
Program menggunakan metode pembandingan berat untuk menentukan bola terberat secara efisien.

Program akan membandingkan berat bola satu per satu dan mencetak bola mana yang memiliki berat terbesar. 
Logika yang digunakan adalah dengan memanfaatkan serangkaian pernyataan if-else untuk menentukan hasil akhir.

## Fitur Program:
1. Input dari Pengguna: Program meminta pengguna untuk memasukkan berat masing-masing bola: B1, B2, B3, dan B4.
2. Pembandingan Berat Bola: Program membandingkan berat masing-masing bola satu per satu menggunakan logika if-else.
3. Output: Program menampilkan bola mana yang paling berat di antara keempat bola tersebut.

# Struktur Program:
1. Definisi Variabel:
   - B1, B2, B3, B4 digunakan untuk menyimpan berat dari masing-masing bola.
2. Input dari Pengguna:
   - Program meminta masukan dari pengguna untuk nilai berat keempat bola tersebut secara berurutan.
3. Logika Perbandingan:
   - Program pertama-tama membandingkan B1 dengan B2. Jika mereka memiliki berat yang sama, maka pembandingan dilanjutkan ke B3 dan B4.
   - Jika B1 lebih besar dari B2, B1 dibandingkan dengan B3 untuk menemukan yang lebih berat.
   - Jika B1 lebih kecil dari B2, B2 dibandingkan dengan B3.
4. Output:
   - Program menampilkan bola terberat yang ditemukan berdasarkan perbandingan tersebut.

# Contoh Penggunaan:
Misalnya, jika pengguna memasukkan berat sebagai berikut:
  - B1 = 2
  - B2 = 1
  - B3 = 6
  - B4 = 5

Maka, program akan membandingkan bola B1 dan B2 terlebih dahulu. Karena B2 lebih berat dari B1, program akan melanjutkan pembandingan B2 dengan B3. Hasil akhirnya adalah B3 yang memiliki berat terbesar dibandingkan bola lainnya.

```go
package main

import (
	"fmt"
)

// Fungsi utama untuk menemukan bola terberat dari 4 bola
func main() {
	// Definisikan berat dari masing-masing bola untuk keperluan simulasi
	var B1, B2, B3, B4 int

	// Input berat masing-masing bola dari pengguna
	fmt.Println("Masukkan berat bola B1: ")
	fmt.Scan(&B1)
	fmt.Println("Masukkan berat bola B2: ")
	fmt.Scan(&B2)
	fmt.Println("Masukkan berat bola B3: ")
	fmt.Scan(&B3)
	fmt.Println("Masukkan berat bola B4: ")
	fmt.Scan(&B4)

	// Langkah 1: Bandingkan B1 dengan B2
	if B1 == B2 {
		// Jika B1 = B2, periksa B3 dengan B4
		if B3 > B4 {
			fmt.Println("Bola terberat adalah B3")
		} else {
			fmt.Println("Bola terberat adalah B4")
		}
	} else if B1 > B2 {
		// Jika B1 > B2, bandingkan B1 dengan B3
		if B1 > B3 {
			fmt.Println("Bola terberat adalah B1")
		} else {
			fmt.Println("Bola terberat adalah B3")
		}
	} else {
		// Jika B1 < B2, bandingkan B2 dengan B3
		if B2 > B3 {
			fmt.Println("Bola terberat adalah B2")
		} else {
			fmt.Println("Bola terberat adalah B3")
		}
	}
}
```
