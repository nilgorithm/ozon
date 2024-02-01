package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func getIndexes(st string) [][]int {
	re := regexp.MustCompile(`\*{2,}`)
	matches := re.FindAllStringIndex(st, -1)
	if len(matches) > 0 {
		return matches
	}
	return nil
}

func findPairs(indexes [][]int, m *map[int][][]int, res *[]int) [][]int {
	var newIndexes [][]int
	for _, rpair := range indexes {
		found := false
		for k, arr := range *m {
			for ii, pair := range arr {
				if pair[0] == rpair[0] && pair[1] == rpair[1] {
					(*m)[k] = append((*m)[k][:ii], (*m)[k][ii+1:]...)
					*res = append(*res, k)
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			newIndexes = append(newIndexes, rpair)
		}
	}

	return newIndexes
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var its int
	fmt.Fscan(r, &its)
	for g := 0; g < its; g++ {
		var y, x int
		fmt.Fscan(r, &y, &x)
		m := make(map[int][][]int)
		var res []int
		for l := 0; l < y; l++ {
			var cs string
			fmt.Fscan(r, &cs)
			if indexes := getIndexes(cs); indexes != nil {
				if len(m) > 0 {
					indexes = findPairs(indexes, &m, &res)
					fmt.Println(m, indexes)
					for _, rpair := range indexes {
						lstatus := 0
						for _, arr := range m {
							for _, pair := range arr {
								if rpair[0] > pair[0] && pair[1] > rpair[1] {
									lstatus += 1
								}
							}
						}
						m[lstatus] = append(m[lstatus], rpair)
					}

				} else {
					m[0] = append(m[0], indexes...)
				}
			}

		}
		sort.Ints(res)
		stringSlice := make([]string, len(res))
		for i, num := range res {
			stringSlice[i] = strconv.Itoa(num)
		}
		resultString := strings.Join(stringSlice, " ")
		fmt.Println(resultString)
	}
}
