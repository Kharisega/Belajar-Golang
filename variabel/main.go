package main

import "fmt"

func main()  {
	// * mendeklarasikan variabel string bersamaan dengan nilai didalamnya
	var firstname string = "Yosafat"

	// * mendeklarasikan variabel terlebih dahulu
	var lastname string

	// * memberi nilai pada variabel yang telah di deklarasikan
	lastname = "Pandu"

	// * mendeklarasikan variabel tanpa tipe data
	var namaawal string = "John"
	namaakhir := "Wick"

	// * perubahan value atau proses assignment tanpa tipe data
	akhirnama := "Wick"
	akhirnama = "Ethan"
	akhirnama = "Bourne"

	// * deklarasi multi variabel
	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	// * Variabel underscore (menampung nilai yang tidak bisa dipakai, tidak digunakan)
	_ = "Apapun itu secepatnya"

	// * Deklarasi Variabel menggunakan keyword new
	name := new(string)

	// TODO %s dibawah akan diganti dengan nilai dari variabel yang berada di parameter
	fmt.Printf("halo %s %s! \n", firstname, lastname)
	fmt.Printf("halo %s %s! \n", namaawal, namaakhir)
	fmt.Printf("halo %s! \n", akhirnama)
	fmt.Printf("halo %s %s %s! \n", seventh, eight, ninth)
	fmt.Println(*name)
}