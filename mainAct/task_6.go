package main
import (
    "fmt"
	//"regexp"
    //"strings"
	"os"
	"bufio"
)



func main() {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int

	fmt.Fscan(in, &t)
	defer out.Flush()

	for g := 0; g < t; g++ {
		var n, m int
		fmt.Fscan(in, &n, &m)
		arr := make([][]int, n)
		row_val := make([][]int, n)
		clm_val := make([][]int, m)
		all_val := make([]int, 5)
		for i := 0; i < n; i++ {
			arr[i] = make([]int, m)
			row_val[i] = make([]int, 5)
		}
		for j := 0; j < m; j++ {
			clm_val[j] = make([]int, 5)
		}

		for i := 0; i < n; i++ {
			var tmp string
			fmt.Fscan(in, &tmp)

			for j := 0; j < m; j++ {
				val := int(tmp[j] - '1')
				arr[i][j] = val
				clm_val[j][val]++
				row_val[i][val]++
				all_val[val]++
			}
		}

		var max_row, max_clm int
		max_val := -1

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				for val := 0; val < 5; val++ {
					if all_val[val] == 0 {
						continue
					}
					curr := 0
					if arr[i][j] == val {
						curr = 1
					}
					if row_val[i][val] + clm_val[j][val] - curr == all_val[val] {
						//fmt.Fprintf(out, "%d : [%d, %d] %d + %d - %d = %d\n", val, i, j, row_val[i][val], clm_val[j][val], curr, all_val[val])
						if max_val < val {
							max_val = val
							max_row = i
							max_clm = j
						}
					} else {
						break
					}
				}
			}
		}


		//for j := 0; j < m; j++ {
			//fmt.Fprintf(out, "%d ", clm_sum[j])
		//}
		//fmt.Fprintln(out, "")
		//fmt.Fprintln(out, "")


		//for i := 0; i < n; i++ {
			//for j := 0; j < m; j++ {
				//fmt.Fprintf(out, "%d ", arr[i][j])
			//}
			//fmt.Fprintln(out, "")
		//}
		fmt.Fprintf(out, "%d %d\n", max_row+1, max_clm+1)
	}
}

