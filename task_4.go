package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	in := bufio.NewReader(os.Stdin)
	var itrs int
	fmt.Fscan(in, &itrs)
	for i := 0; i < itrs; i++ {
		var persns int
		fmt.Fscan(in, &persns)
		optimal := true
		min, max := 15, 30
		for j := 0; j < persns; j++ {
			var symb string
			var dig int
			//fmt.Fscanf(in, "%s %d", &symb, &dig)
			fmt.Fscan(in, &symb, &dig)
			if optimal {
				if min <= dig && dig <= max {
					if symb[0] == '>' {
						min = dig
						fmt.Fprintln(out, min)
					} else {
						max = dig
						fmt.Fprintln(out, min)
					}
				} else if dig > max && symb[0] == '<' {
					fmt.Fprintln(out, min)
				} else if dig < min && symb[0] == '>' {
					fmt.Fprintln(out, min)
				} else {
					optimal = false
					fmt.Fprintln(out, -1)
				}
			} else {
				fmt.Fprintln(out, -1)
			}
		}
		fmt.Fprintln(out, "")
	}
}

