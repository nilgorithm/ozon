package main

import (
	"fmt"
	"regexp"
	"strings"
)

func splitIntoPlates(s string) (string, bool) {
	plateRegex := regexp.MustCompile(`[A-Z][\d]{1,2}[A-Z]{2}`)
	matches := plateRegex.FindAllString(s, -1)
	concatenated := strings.Join(matches, "")

	if concatenated == s {
		return strings.Join(matches, " "), true
	}
	return "", false
 }

func main() {
	var its int
	fmt.Scan(&its)
	for i := 0; i < its; i++ {
		var st string
		fmt.Scan(&st)
		if result, valid := splitIntoPlates(st); valid {
			fmt.Println(result)
		} else {
			fmt.Println("-")
		}
	}
}
