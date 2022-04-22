package main

import "fmt"
import "strings"
import (
	"math/rand"
	"time"
	"math"
)

func main()  {
	name := []string{"John", "Wick"}
	printMessage("Halo", name)

	rand.Seed(time.Now().Unix())
    var randomValue int

    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)
    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)
    randomValue = randomWithRange(2, 30)
    fmt.Println("random number:", randomValue)

	data := time.Now().Minute()
	fmt.Println(data)

	var diameter float64 = 15
    var area, circumference = calculate(diameter)

    fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
    fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

	var avg = hitung(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
    var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
    fmt.Println(msg)

	// ! Fungsi Closure sebagai variabel
	// ? Closure adalah fungsi tanpa nama
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n{
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	// ? %v digunakan untuk menampilkan segala jenis data. Bisa array, int, float, bool, dan lainnya
	numbers := []int{2, 3, 5, 3, 4, 2, 3}
	min, max := getMinMax(numbers)
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)
}

func printMessage(message string, arr []string)  {
	nameString := strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	value := rand.Int() % (max - min + 1) + min
	return value
}

// * Fungsi yang mengembalikan lebih dari satu nilai dan predefined return value
// ? fungsi math.Pow() digunakan untuk memangkat nilai
// ? Konstanta math.Pi mempresentasikan nilai pi atau memiliki nilai 22/7
func calculate(d float64) (area float64, circumference float64) {
		area = math.Pi * math.Pow(d / 2, 2)
		circumference = math.Pi * d
	
		return
	}

// ! Fungsi variadic atau fungsi dengan parameter sejenis yang tak terbatas
// ? Untuk konversi tipe data cukup memberi tipedata tanda kurung maka sudah menjadi fungsi untuk mengonversi nilai
func hitung(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}
