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

	for j := 0; j < t; j++ {
		var res float32
		var n, p int
		fmt.Fscan(in, &n, &p)
		s := 0
		for i := 0; i < n; i++ {
			var price int
			fmt.Fscan(in, &price)
			s += (price * p) % 100
		}
		res = float32(s) / 100
		fmt.Fprintf(out, "%.2f\n", res)
	}
}

