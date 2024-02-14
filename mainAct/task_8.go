package main
import (
    "fmt"
	//"regexp"
    //"strings"
	"os"
	"bufio"
)

func getOpt(a, b int) int {
	if a == b {
		return -1
	} else if a < b {
		return 0
	}
	return 1
}

//func check(rm, lm, val, k int) bool {
	//rm = rm << (k*2)
	//lm = lm << (k*2)
	//return rm & val == rm || lm & val == lm
//}

func check(rm, val, k, j int, out *bufio.Writer) bool {
	rm = rm << (k*2*j)
	tmpm := 0
	for i := 0; i < 2*k; i++ {
		tmpm = (tmpm << 1) | 1
	}
	tmpm = tmpm << (k*2*j)
	tmp := val & tmpm
	//fmt.Fprintf(out, "%d & %d/(%d|%d)\n", rm, tmp, val, tmpm)
	return rm & tmp == tmp && rm & tmp == rm
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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

	for g := 0; g < t; g++ {
		var n int
		fmt.Fscan(in, &n)

		max_l := make([]int, n)
		opts := make([]int, n-1)
		prev := -1
		for i := 0; i < n; i++ {
			var curr int
			fmt.Fscan(in, &curr)
			if prev >= 0 {
				opts[i-1] = getOpt(prev, curr)
			}

			prev = curr
		}
		if n == 1 {
			fmt.Fprintln(out, "0")
			continue
		}

		//fmt.Fprintln(out, "opts")
		for i := 0; i < n - 1; i++ {
			//fmt.Fprintf(out, "%d", opts[i])
			if i < n - 2 {
				//fmt.Fprintf(out, " ")
			}
		}
		//fmt.Fprintln(out, "")
		//fmt.Fprintln(out, "")


		ks := make([][]int, n/2)
		for i := range ks {
			//k := i+1
			ks[i] = make([]int, 0)
		}

		zeros_found := 0
		curr_i := -1
		ones_found := 0

		for i := 0; i < n-1; i++ {
			//fmt.Fprintf(out, "%d: (%d) %d [%d, %d]\n", i, opts[i], curr_i, zeros_found, ones_found)
			if opts[i] < 0 {
				k := min(ones_found, zeros_found)
				if k > 0 && curr_i >= 0 {
					//fmt.Fprintf(out, "+%d: (%d) [%d, %d] - k=%d\n", curr_i, opts[i], zeros_found, ones_found, k)
					ks[k-1] = append(ks[k-1], curr_i)
				}
				zeros_found = 0
				ones_found = 0
				curr_i = -1
			} else if opts[i] == 0 && curr_i < 0 {
				zeros_found++
				ones_found = 0
			} else if opts[i] == 0 {
				if ones_found > 0 {
					//ks[ones_found] = append(ks[ones_found], curr_i)
					k := min(ones_found, zeros_found)
					//fmt.Fprintf(out, "+%d: (%d) [%d, %d] - k=%d\n", curr_i, opts[i], zeros_found, ones_found, k)
					ks[k-1] = append(ks[k-1], curr_i)
				}
				zeros_found = 1
				ones_found = 0
				curr_i = -1
			} else if opts[i] == 1 && curr_i < 0 {
				if zeros_found > 0 {
					curr_i = i-1
					ones_found = 1
				} else {
					ones_found = 0
				}
			} else if opts[i] == 1 {
				k := min(ones_found, zeros_found)
				if k > 0 {
					max_l[k-1] = 1
				}
				if ones_found <= zeros_found {
					ones_found++
				} else {
					if curr_i >= 0 {
						k := min(ones_found, zeros_found)
						//fmt.Fprintf(out, "+%d: (%d) [%d, %d] - k=%d\n", curr_i, opts[i], zeros_found, ones_found, k)
						ks[k-1] = append(ks[k-1], curr_i)
					}
					zeros_found = 0
					ones_found = 0
					curr_i = -1
				}
			}
		}
		if ones_found > 0 {
			k := min(ones_found, zeros_found)
			//fmt.Fprintf(out, "+%d: (end) [%d, %d] - k=%d\n", curr_i, zeros_found, ones_found, k)
			ks[k-1] = append(ks[k-1], curr_i)
		}

		for i := range ks {
			k := i+1
			inter := -1
			count := 0
			for j := range ks[i] {
				pos := ks[i][j]
				//fmt.Fprintf(out, "%d: [%d] = %d /+ %d\n", k, j, pos, count)
				if pos < 0 {
					break
				}
				if inter < 0 {
					count++
				} else if inter == pos - k {
					count++
				} else {
					count = 1
				}
				inter = pos + k
				if count > max_l[i] {
					max_l[i] = count
				}
			}
		}



		for i := 0; i < n; i++ {
			fmt.Fprintf(out, "%d", max_l[i])
			if i < n - 1 {
				fmt.Fprintf(out, " ")
			}
		}
		fmt.Fprintln(out, "")
	}
}

