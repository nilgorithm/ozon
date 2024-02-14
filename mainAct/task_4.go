package main
import (
    "fmt"
	//"regexp"
    "strings"
	"os"
	"bufio"
)


func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func replace(str string, replacement string, index int) string {
    return str[:index] + replacement + str[index+1:]
}

func main() {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int

	fmt.Fscan(in, &t)
	defer out.Flush()

	for j := 0; j < t; j++ {
		var n, m int
		fmt.Fscan(in, &n, &m)
		var arr []string = make([]string, n)
		x_tl, y_tl := -1, -1 //tl - top left
		c_tl := "x"
		x_br, y_br := -1, -1 //br - bottom right
		c_br := "x"
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])

			ind_a := strings.Index(arr[i], "A")
			ind_b := strings.Index(arr[i], "B")

			x_max := max(ind_a, ind_b)
			x_min := min(ind_a, ind_b)
			if x_max >= 0 {
				if x_min >= 0 {
					y_tl = i
					x_tl = x_min
					c_tl = arr[y_tl][x_tl:x_tl+1]

					y_br = i
					x_br = x_max
					c_br = arr[y_br][x_br:x_br+1]
				}
				if x_tl < 0 {
					y_tl = i
					x_tl = x_max
					c_tl = arr[y_tl][x_tl:x_tl+1]
				} else {
					y_br = i
					x_br = x_max
					c_br = arr[y_br][x_br:x_br+1]
				}
			}
		}
		c_tl = strings.ToLower(c_tl)
		c_br = strings.ToLower(c_br)
		if x_tl % 2 == 1 {
			arr[y_tl] = replace(arr[y_tl], c_tl, x_tl-1)

			x_tl -= 1
		}
		for y := y_tl-1; y >= 0; y-- {
			arr[y] = replace(arr[y], c_tl, x_tl)
		}
		for x := x_tl-1; x >= 0; x-- {
			arr[0] = replace(arr[0], c_tl, x)
		}

		if x_br % 2 == 1 {
			arr[y_br] = replace(arr[y_br], c_br, x_br+1)

			x_br += 1
		}
		for y := y_br+1; y < n; y++ {
			arr[y] = replace(arr[y], c_br, x_br)
		}
		for x := x_br+1; x < m; x++ {
			arr[n-1] = replace(arr[n-1], c_br, x)
		}

		for i := 0; i < n; i++ {
			fmt.Fprintln(out, arr[i])
		}
		//fmt.Fprintf(out, "tl %d %d - %s\nbr %d %d - %s\n", x_tl, y_tl, c_tl, x_br, y_br, c_br)
	}
}

