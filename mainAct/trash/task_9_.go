package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func getInd(st, rg string) [][]int {
	re := regexp.MustCompile(rg)
	matches := re.FindAllStringIndex(st, -1)
	if len(matches) > 0 {
		return matches
	}
	return nil
}

func decrStr(s string, i1, i2 int) string {
	var builder strings.Builder
	for i, _ := range s {
		if i != i1 && i != i2-1 {
			builder.WriteByte(s[i])
		}
	}
	return builder.String()
}

func checkSeq(st string, l int) string {
	pairs := []string{`Y.*Z`, `X.*Z`, `X.*Y`}
	for _, p := range pairs {
		for {
			inds := getInd(st, p)
			if inds == nil { // Проверяем, было ли найдено совпадение
				break
			}
			i1, i2 := inds[0][0], inds[0][1]
			//println(i1, i2, st)
			st = decrStr(st, i1, i2)
			//println(st)
			//fmt.Println(st)
		}
	}
	if len(st) > 0 {
		return "No"
	}
	return "Yes"
}

func main() {
	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var in *bufio.Reader
	in = bufio.NewReader(os.Stdin)
	var num int
	fmt.Fscan(in, &num)

	for i := 0; i < num; i++ {
		var slen int
		fmt.Fscan(in, &slen)
		var inp string
		fmt.Fscan(in, &inp)
		res := checkSeq(inp, slen)
		fmt.Fprint(out, res)
		fmt.Fprint(out, "\n")

	}
}
