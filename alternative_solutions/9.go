package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func getIndexes(st string) [][]int {
	re := regexp.MustCompile(`\*{2,}`)
	matches := re.FindAllStringIndex(st, -1)
	if len(matches) > 0 {
		return matches
	}
	return nil
}

//func removeExisted() {
//	for i, pair := range m[key] {
//		fmt.Println(pair)
//		if pair == pairToRemove {
//			m[key] = append(m[key][:i], m[key][i+1:]...)
//			break
//		}
//	}
//
//}

func main() {
	r := bufio.NewReader(os.Stdin)
	var its int
	fmt.Fscan(r, &its)
	for g := 0; g < its; g++ {
		var y, x int
		fmt.Fscan(r, &y, &x)
		m := make([][]int, 0)
		//m := make(map[int][][]int)
		for l := 0; l < y; l++ {
			var cs string
			fmt.Fscan(r, &cs)
			if indexes := getIndexes(cs); indexes != nil {
				m = append(m, indexes...)
			}

		}
		fmt.Println(m)
		fmt.Println(len(m))
	}
}
