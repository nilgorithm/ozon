package main

import (
	"fmt"
)

func isValidDate(day, month, year int) string {
	daysInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		daysInMonth[1] = 29 // февраль високосный
	}

	if month < 1 || month > 12 {
		return "NO"
	}

	if day < 1 || day > daysInMonth[month-1] {
		return "NO"
	}

	return "YES"
}

func main() {
	var t, day, month, year int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&day, &month, &year)
		fmt.Println(isValidDate(day, month, year))
	}
}
