package main

import "fmt"

func avg(a, b float64) {
	fmt.Println("Среднее число: ", (a + b) / 2)
}

func main() {
	var (
		firstDigit  int
		secondDigit int
	)
	fmt.Scan(&firstDigit, &secondDigit) //сканирование ввода для первой и второй переменной по-порядку
	avg(float64(firstDigit), float64(secondDigit))
}
