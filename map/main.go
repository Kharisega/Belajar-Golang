package main

import "fmt"

func main()  {
	// TODO Materi Map
	// ? Map dalam golang seperti object dalam bahasa lain

	// * Penggunaan Map
	var chicken = map[string]int{}
	
	chicken["januari"] = 50
	chicken["februari"] = 40
	
	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei",     chicken["mei"])     // mei 0

	// * Pemanggilan item map menggunakan for - range
	// ! Output yang ditampilkan muncul secara random
	var ayam = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}
	
	for key, val := range ayam {
		fmt.Println(key, "  \t:", val)
	}

	// * Menghapus item map
	var pitik = map[string]int{"januari": 50, "februari": 40}

	fmt.Println(len(pitik)) // 2
	fmt.Println(pitik)
	
	delete(pitik, "januari")
	
	fmt.Println(len(pitik)) // 1
	fmt.Println(pitik)

	// * Mengecek keberadaan item dengan key tertentu
	var jago = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = jago["mei"]
	
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	// * Kombinasi Slice dan Map

	// ! Penggunaan Map jika lebih dari satu data (biasa)
	var kampung = []map[string]string{
		map[string]string{"name": "chicken blue",   "gender": "male"},
		map[string]string{"name": "chicken red",    "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}
	
	for _, chicken := range kampung {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	// ! Penggunaan jika lebih dari satu data secara ringkas (jika go sudah versi terbaru)
	var pitikan = []map[string]string{
		{"name": "chicken blue",   "gender": "male"},
		{"name": "chicken red",    "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	fmt.Println(pitikan)

	// ! Tiap elemen dalam map bisa memiliki key yang berbeda-beda
	var data = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}

	fmt.Println(data)
}