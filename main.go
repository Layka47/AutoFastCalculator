package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var k = 0
	for i := 0; i < 3; i++ {
		fmt.Println("Братанчик, дождись")
	}
	for true {
		a := rand.Intn(10000)
		b := rand.Intn(10000)
		sleepDuration := time.Duration(1000-k) * time.Millisecond

		sum, multiply := namedSumAndMultiply(a, b)
		if sleepDuration <= 70*time.Millisecond {
			sleepDuration = 70 * time.Millisecond
			fmt.Println("Чё хавать будем?")
		}
		time.Sleep(sleepDuration)
		if sleepDuration > 71*time.Millisecond {
			fmt.Printf("Сумма(a,b) = %v; Произведение(a,b) = %v; a = %v; b = %v; Задержка: %v\n", sum, multiply, a, b, sleepDuration)
			k += 20
		}
	}

}

func namedSumAndMultiply(first, second int) (sum, multiply int) {
	sum = first + second
	multiply = first * second
	return sum, multiply
}
