package main

import "fmt"

func main()  {
	// TODO MATERI ARRAY

	// * Penggunaan Array biasa
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"
	// names[4] = "ez" //! jika melebihi index seperti disamping akan error

	fmt.Println(names[0], names[1], names[2], names[3]) // trafalgar d water law
	fmt.Println(names) // [trafalgar d water law]

	// * Inisialisasi Nilai Awal Array
	// cara horizontal
	var buah = [4]string{"apel", "mangga", "melon", "anggur"}

	// cara vertikal
	var buahh = [4]string{
		"apel",
		"mangga",
		"melon",
		"anggur",
	}

	fmt.Println("Jumlah elemen \t\t", len(buah))
	fmt.Println("Isi semua elemen \t", buahh)

	// * Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
	var angka = [...]int{1,2,3,4,5,3,2,3,1}

	fmt.Println("data array \t:", angka)
	fmt.Println("jumlah elemen \t:", len(angka))

	// * Array Multidimensi
	var numbers1 = [2][3]int{[3]int{3,2,3}, [3]int{3,4,5}}
	var numbers2 = [2][3]int{{ 3,2,3 }, {3,4,5}}

	fmt.Println("numbers1 = ", numbers1)
	fmt.Println("numbers2 = ", numbers2)

	// * Perulangan Elemen Array dengan for
	var sayur = [4]string{"bayam", "kangkung", "kol", "sawi"}

	for i := 0; i < len(sayur); i++ {
		fmt.Printf("elemen %d : %s\n", i, sayur[i])
	}

	// * Perulangan Elemen Array dengan for - range
	var sayurs = [4]string{"bayam", "kangkung", "kol", "sawi"}

	for i, sayuran := range sayurs {
		fmt.Printf("elemen %d : %s\n", i, sayuran)
	}

	// * penggunaan underscore di for - range (untuk menampung nilai yang tidak digunakan)
	var sayurku = [4]string{"bayam", "kangkung", "kol", "sawi"}

	for _, sayuran := range sayurku {
		fmt.Printf("elemen : %s\n", sayuran)
	}

	// * Alokasi elemen array menggunakan keyword make
	var obat = make([]string, 2)
	obat[0] = "paracetamol"
	obat[1] = "panadol"

	fmt.Println(obat)
}