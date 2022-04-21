package main

import "fmt"

func main()  {
	// TODO Materi Slice ( Masih satu kesatuan dengan )

	// * Relasi antara Array dengan Slice
	var buah = []string{"apel", "stroberi", "melon", "pisang"}
	fmt.Println(buah[2]) // melon

	var buah1 = []string{"apel", "stroberi", "melon", "pisang"} // slice
	var buah2 = [4]string{"apel", "stroberi", "melon", "pisang"} // array
	var buah3 = [...]string{"apel", "stroberi", "melon", "pisang"} // array
	fmt.Println(buah1)
	fmt.Println(buah2)
	fmt.Println(buah3)

	// * operasi Slice
	var buahBaru = buah[0:2]
	fmt.Println(buahBaru)

	// * Slice Merupakan Tipe Data Reference
	// ? Slice merupakan tipe data reference atau referensi. Artinya jika ada slice baru yang terbentuk dari slice lama, maka data elemen slice yang baru akan memiliki alamat memori yang sama dengan elemen slice lama. Setiap perubahan yang terjadi di elemen slice baru, akan berdampak juga pada elemen slice lama yang memiliki referensi yang sama.

	var sayur = []string{"kol", "sawi", "kangkung", "bayam"}

	var sayurs = sayur[0:3]
	var sayuran = sayur[1:4]

	var asayurs = sayurs[1:2]
	var asayuran = sayuran[0:1]

	fmt.Println(sayur) // [kol sawi kangkung bayam]
	fmt.Println(sayurs) // [kol sawi kangkung]
	fmt.Println(sayuran) // [sawi kangkung bayam]
	fmt.Println(asayurs) // [sawi]
	fmt.Println(asayuran) // [sawi]

	//! sayur sawi diganti menjadi brokoli
	asayuran[0] = "brokoli"

	fmt.Println(sayur) // [kol brokoli kangkung bayam]
	fmt.Println(sayurs) // [kol brokoli kangkung]
	fmt.Println(sayuran) // [brokoli kangkung bayam]
	fmt.Println(asayurs) // [brokoli]
	fmt.Println(asayuran) // [brokoli]

	// * Fungsi cap()
	// ? digunakan untuk melihat kapasitas dari slice yang ada
	// ! bukan jumlah elemen tetapi kapasitas
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits))  // len: 4
	fmt.Println(cap(fruits))  // cap: 4
	
	var aFruits = fruits[0:3]
	fmt.Println(aFruits)
	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4
	
	var bFruits = fruits[1:4]
	fmt.Println(bFruits)
	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3

	// * fungsi append()
	var fruit = []string{"apple", "grape", "banana"}
	var cFruits = append(fruit, "papaya")
	
	fmt.Println(fruit)  // ["apple", "grape", "banana"]
	fmt.Println(cFruits) // ["apple", "grape", "banana", "papaya"]

	var dFruits = fruit[0:2]

	fmt.Println(cap(dFruits)) // 3
	fmt.Println(len(dFruits)) // 2
	
	fmt.Println(fruit)  // ["apple", "grape", "banana"]
	fmt.Println(dFruits) // ["apple", "grape"]
	
	var eFruits = append(dFruits, "papaya")
	
	fmt.Println(fruit)  // ["apple", "grape", "papaya"]
	fmt.Println(dFruits) // ["apple", "grape"]
	fmt.Println(eFruits) // ["apple", "grape", "papaya"]

	//  * fungsi copy()
	// ! dst berisi elemen yang terecopy dari src, jika dst dicetak maka akan keluar n tetapi jika n maka yang muncul hasil len(dst)
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)
	
	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3

	// * Pengaksesan elemen slice dengan 3 index
	// ? index yang ketiga digunakan untuk menentukan kapasitas
	var buahh = []string{"apple", "grape", "banana"}
	var abuahh = buahh[0:2]
	var bbuahh = buahh[0:2:2]
	
	fmt.Println(buahh)      // ["apple", "grape", "banana"]
	fmt.Println(len(buahh)) // len: 3
	fmt.Println(cap(buahh)) // cap: 3
	
	fmt.Println(abuahh)      // ["apple", "grape"]
	fmt.Println(len(abuahh)) // len: 2
	fmt.Println(cap(abuahh)) // cap: 3
	
	fmt.Println(bbuahh)      // ["apple", "grape"]
	fmt.Println(len(bbuahh)) // len: 2
	fmt.Println(cap(bbuahh)) // cap: 2
}