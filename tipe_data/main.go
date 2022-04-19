package main

import "fmt"

func main() {
	var angka1 uint8 = 255
	var angka2 uint16 = 65535
	var angka3 uint32 = 4294967295
	var angka4 uint64 = 18446744073709551615

	// * sama dengan uint32 atau uint64 tergantung dari nilai yang ada
	var angka5 uint = 18446744073709551615

	// * sama dengan uint8
	var angka6 byte = 255

	var angka7 int8 = -128
	var angka8 int16 = -32768
	var angka9 int32 = -2147483648
	var angka10 int64 = -9223372036854775808

	//  * sana dengan int32 atau int64 tergantung dari nilai yang ada
	var angka11 int = -9223372036854775808

	// * sama dengan int32
	var angka12 rune = -2147483648

	// ! %d pada fmt.Printf() digunakan untuk memformat data numerik non-desimal.
	fmt.Printf("bilangan : %d\n", angka1)
	fmt.Printf("bilangan : %d\n", angka2)
	fmt.Printf("bilangan : %d\n", angka3)
	fmt.Printf("bilangan : %d\n", angka4)
	fmt.Printf("bilangan : %d\n", angka5)
	fmt.Printf("bilangan : %d\n", angka6)
	fmt.Printf("bilangan : %d\n", angka7)
	fmt.Printf("bilangan : %d\n", angka8)
	fmt.Printf("bilangan : %d\n", angka9)
	fmt.Printf("bilangan : %d\n", angka10)
	fmt.Printf("bilangan : %d\n", angka11)
	fmt.Printf("bilangan : %d\n", angka12)

	// * tipe data bilangan desimal atau float
	var desimal = 2.62

	// ! $f dibawah digunakan untuk memformat data numerik desimal menjadi string dengan desimal 6 digit
	// ! $.3f dibawah digunakan untuk memformat data numerik desimal menjadi string dengan desimal 3 digit
	fmt.Printf("bilangan : %f\n", desimal)
	fmt.Printf("bilangan : %.3f\n", desimal)

	// * tipe data boolean
	var exist bool = true

	// ! %t dibawah digunakan untuk memformat data boolean
	fmt.Printf("Apakah ada? %t \n", exist)

	// * tipe data string
	var hadir string = "Ada"

	// ! %s dibawah digunakan untuk memformat data string
	fmt.Printf("Apakah ada? %s \n", hadir)
	
	// ? nil bukan merupakan tipe data, melainkan sebuah nilai. Variabel yang isi nilainya nil berarti memiliki nilai kosong

}