package main

import "fmt"

func main()  {
	// ** Materi Kondisi

	point := 8

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	// ! Variabel Temporary Pada if - else
	// ? Variabel temporary adalah variabel yang hanya bisa digunakan pada blok seleksi kondisi di mana ia ditempatkan saja. (variabel scope)

	nilai := 8840.0

	if percent := nilai / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
}