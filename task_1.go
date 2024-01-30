package main

import (
  "os"
  "fmt"
  "bufio"
)

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
			if exists && sum[i] == i + (tsh-(1+(i-1)*2)) {
				if i==tsh {
					fmt.Println("Yes")
				}
				continue
			} else {
				fmt.Println("No")
				break
			}
		}

	}
}
