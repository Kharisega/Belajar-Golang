package main

import "fmt"
import "strings"
import (
	"math/rand"
	"time"
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
}

func printMessage(message string, arr []string)  {
	nameString := strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	value := rand.Int() % (max - min + 1) + min
	return value
}