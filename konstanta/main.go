package main

import "fmt"

func main()  {
	// * Konstanta sama seperti variabel namun nilai didalamnya tidak bisa berubah atau permanen
	// ! penggunaan konstanta sama seperti variabel
	const firstname string = "Sandy"

	// ! deklarasi konstanta dengan teknik type inference
	const lastname = "Willy"

	fmt.Printf("Halo %s!\n", firstname)
	fmt.Printf("Nice to meet you, %s!\n", lastname)
}