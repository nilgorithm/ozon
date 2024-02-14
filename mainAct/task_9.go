//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"regexp"
//)
//
//func momt(inp string) {
//
//	xRegex := regexp.MustCompile("X")
//	yRegex := regexp.MustCompile("Y")
//	zRegex := regexp.MustCompile("Z")
//
//	xMatches := xRegex.FindAllStringIndex(inp, -1)
//	yMatches := yRegex.FindAllStringIndex(inp, -1)
//	zMatches := zRegex.FindAllStringIndex(inp, -1)
//	fmt.Println(xMatches)
//	fmt.Println(yMatches)
//	fmt.Println(zMatches)
//}
//
//func main() {
//	var out *bufio.Writer
//	out = bufio.NewWriter(os.Stdout)
//	defer out.Flush()
//	var in *bufio.Reader
//	in = bufio.NewReader(os.Stdin)
//	var num int
//	fmt.Fscan(in, &num)
//
//	for i := 0; i < num; i++ {
//		var slen string
//		fmt.Fscan(in, &slen)
//		var inp string
//		fmt.Fscan(in, &inp)
//		momt(inp)
//	}
//}

package main

import (
	"fmt"
)

func canFormPairs(events string) bool {
	countX, countY, countZ := 0, 0, 0

	// Подсчитываем количество каждого события
	for _, event := range events {
		switch event {
		case 'X':
			countX++
		case 'Y':
			countY++
		case 'Z':
			countZ++
		}
	}

	// Проверяем, что количество каждого события позволяет сформировать пары без остатка
	// и что общее количество событий четное
	//if (countX+countY+countZ)%2 != 0 {
	//	return false
	//}

	// Проверяем, что количество событий позволяет сформировать только правильные пары
	// Пары XY и XZ требуют наличия X, а пара YZ требует Y и Z
	for countY > 0 || countX > 0 {
		if countZ > 0 && countY > 0 {
			countY--
			countZ--
		} else if countZ > 0 && countX > 0 {
			countX--
			countZ--
		} else if countX > 0 && countY > 0 {
			countX--
			countY--
		} else {
			return false
		}

	}

	// Если остались события Y или Z, они должны образовать пару YZ
	return countX == countZ || countX+countY == countZ
}

func main() {
	inputs := []string{"YXYYZZ", "ZYXZ", "XYZXZY"}
	for _, events := range inputs {
		if canFormPairs(events) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
