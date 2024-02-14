package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func checkSwap(s1, s2 string) bool {
	ls1 := len(s1)
	if ls1 != len(s2) {
		return false
	}
	for i := 0; i < ls1; i++ {
		if s1[i] != s2[i] {
			if ls1-1-i == 1 {
				return s1[i] == s2[i+1] && s2[i] == s1[i+1]
			} else if ls1-i-1 > 1 {
				return s1[i] == s2[i+1] && s1[i+2:] == s2[i+2:] && s2[i] == s1[i+1]
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	var logins []string

	for i := 0; i < n; i++ {
		scanner.Scan()
		login := scanner.Text()
		logins = append(logins, login)
	}

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < m; i++ {
		scanner.Scan()
		log2 := scanner.Text()
		for l, log1 := range logins {
			if checkSwap(log1, log2) {
				fmt.Println(1)
				break
			} else if l == len(logins)-1 {
				fmt.Println(0)
			}
		}
	}
}
