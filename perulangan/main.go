package main

import "fmt"

func main()  {
	// TODO Materi Perulangan atau for
	// * Di Go perulangan hanya menggunakan for saja, tapi gabungan dari for, foreach, dan while

	// * Penggunaan for biasa
	for i := 0; i < 5; i++ {
		fmt.Printf("Angka %d\n", i)
	}

	// * Penggunaan for dengan argumen hanya kondisi
	u := 0

	for u < 5 {
		fmt.Printf("Angka %d\n", u)
		u++
	}

	// * Penggunaan for tanpa argumen
	o := 0

	for {
		fmt.Printf("Angka %d\n", o)

		o++
		if o == 5 {
			break
		}
	}

	// * Penggunaan keyword break dan continue
	for i := 0; i <= 10; i++ {
		if i % 2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Printf("Angka %d\n", i)
	}

	// * Perulangan bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	// * perulangan dengan memanfaatkan label
	// ? label ini digunakan oleh keyword break dan continue untuk menuju ke perulangan dengan label tertentu

	outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Printf("Matriks [%d][%d]\n", i, j)
		}
	}
}