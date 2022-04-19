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

	// * Kondisi dengan Switch Case
	skor := 6

	switch skor {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// * Switch case dengan beberapa kondisi
	switch skor {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// * kurung kurawal pada case dan default
	// ! jika kondisi menggunakan proses setelah switch tidak perlu diberi variabel
	switch {
	case skor == 8:
		fmt.Println("perfect")
	case (skor < 8) && (skor > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// * fallthrough dalam switch case
	// ? fallthrough digunakan untuk memaksa proses pengecekan diteruskan dari satu case ke case selanjutnya dengan tanpa menghiraukan kondisinya, jadi satu case di pengecekan selanjutnya tersebut selalu bersikap benar meskipun hasilnya salah

	poin := 7

	switch {
	case poin == 8:
		fmt.Println("perfect")
	case (poin < 8) && (poin > 3):
		fmt.Println("awesome")
		fallthrough
	case poin < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// * Kondisi bersarang
	biji := 10

	if biji > 7 {
		switch biji {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if biji == 5 {
			fmt.Println("not bad")
		} else if biji == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if biji == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}