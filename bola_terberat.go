package main

import "fmt"

// fungsi utama untuk menemukan bola terberat dari 4 bola tersebut
func main() {
	// Definisikan berat dari masing-masing bola 
	var B1, B2, B3, B4 int

	// Input berat dari masing-masing bola
	fmt.Println("Masukkan berat dari bola B1: ")
	fmt.Scan(&B1)
	fmt.Println("Masukkan berat dari bola B2: ")
	fmt.Scan(&B2)
	fmt.Println("Masukkan berat dari bola B3: ")
	fmt.Scan(&B3)
	fmt.Println("Masukkan berat dari bola B4: ")
	fmt.Scan(&B4)

	// Langkah 1: Bandingkan B1 dengan B2
	if B1 == B2 {
		// Jika B1 = B2, Periksalah B3 dengan B4
		if B3 > B4 {
			fmt.Println("Bola terberat adalah B3")
		} else {
			fmt.Println("Bola terberat adalah B4")
		}
	} else if B1 > B2 {
		// Jika B1 > B2, Bandingkan B1 dengan B3
		if B1 > B3 {
			fmt.Println("Bola terberat adalah B1")
		} else {
			fmt.Println("Bola terberat adalah B3")
		}
	} else {
		// Jika B1 < B2, Maka bandingkan B2 dengan B3
		if B2 > B3 {
			fmt.Println("Bola terberat adalah B2")
		} else {
			fmt.Println("Bola terberat adalah B3")
		}
	}
}