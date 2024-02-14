package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var repeat int
	fmt.Scan(&repeat)

	scanner := bufio.NewScanner(os.Stdin)

	for j := 0; j < repeat; j++ {
		var sckip int
		fmt.Scan(&sckip)
		var cont []string
		for i := 0; i < sckip; i++ {
			if scanner.Scan() {
				line := scanner.Text()
				cont = append(cont, line)
			}
		}
		fmt.Println(len(cont), cont)
		cont = nil
	}
}
