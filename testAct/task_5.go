package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	repeat, _ := strconv.Atoi(strings.TrimSpace(line))
	for j := 0; j < repeat; j++ {
		_, _ = reader.ReadString('\n')
		var res []string
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		tokens := strings.Split(text, " ")
		var startNumb int
		var prevNumb int
		var count int
		for i, token := range tokens {
			currNumb, _ := strconv.Atoi(token)
			if i > 0 {
				if prevNumb-currNumb == -1 && currNumb > startNumb {
					count += 1
				} else if prevNumb-currNumb == 1 && currNumb < startNumb {
					count -= 1
				} else {
					res = append(res, strconv.Itoa(count), strconv.Itoa(currNumb))
					count = 0
					startNumb = currNumb
				}
			} else {
				res = append(res, strconv.Itoa(currNumb))
				startNumb = currNumb
			}
			prevNumb = currNumb
		}
		result := strings.Join(res, " ")
		fmt.Println(len(res) + 1)
		fmt.Println(result + " " + strconv.Itoa(count))
	}
}
