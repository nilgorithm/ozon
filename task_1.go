package main

import (
	"bufio"
	"fmt"
	"os"
)

// решение регрессией
func main() {
	r := bufio.NewReader(os.Stdin)
	var its int //количество итераций
	fmt.Scan(&its)
	tsh := 4 //количество типов кораблей
	for i := 0; i < its; i++ {
		sum := make(map[int]int)
		var temp int
		for i := 0; i < 10; i++ {
			fmt.Fscan(r, &temp)
			sum[temp] += 1
		}
		for i := 1; i < tsh+1; i++ {
			_, exists := sum[i]
			if exists && sum[i] == tsh-i+1 {
				if i == tsh {
					fmt.Printf("YES")
				}
				continue
			} else {
				fmt.Printf("NO")
				break
			}
		}

	}
}
