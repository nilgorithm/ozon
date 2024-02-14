package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Pair struct {
	Start string
	End   string
}

func findUniquePairs(input string, pairs []Pair) string {
	//	[][2]int {
	var result [][2]int
	usedZ := make(map[[2]int]bool)
	for _, p := range pairs {
		yRegex := regexp.MustCompile(p.Start)
		zRegex := regexp.MustCompile(p.End)
		yMatches := yRegex.FindAllStringIndex(input, -1)
		zMatches := zRegex.FindAllStringIndex(input, -1)

		for _, yMatch := range yMatches {
			for _, zMatch := range zMatches {
				carr := [2]int{yMatch[0], zMatch[0]}
				if !usedZ[carr] {
					result = append(result, carr)
					usedZ[carr] = true
					break
				}
			}
		}
	}
	fmt.Println("Уникальные пары: ", result)
	fmt.Println("Уникальные пары: ", usedZ)
	if len(result) > 0 && len(result) == len(input)/2 {
		return "Yes"
	}
	return "No"
}

func main() {
	pairs := []Pair{
		{Start: "Y", End: "Z"},
		{Start: "X", End: "Y"},
		{Start: "X", End: "Z"},
	}
	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var in *bufio.Reader
	in = bufio.NewReader(os.Stdin)
	var num int
	fmt.Fscan(in, &num)

	for i := 0; i < num; i++ {
		var slen string
		fmt.Fscan(in, &slen)
		var inp string
		fmt.Fscan(in, &inp)
		res := findUniquePairs(inp, pairs)
		fmt.Fprint(out, res)
		fmt.Fprint(out, "\n")
	}
}

//func getIndexes(st string) [][]int {
//	re := regexp.MustCompile(`Y.*Z`)
//	matches := re.FindAllStringIndex(st, -1)
//	if len(matches) > 0 {
//		return matches
//	}
//	return nil
//}
