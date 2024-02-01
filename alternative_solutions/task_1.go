package main

import (
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"bufio"
	"os"
)
//решение регуляркой с расширенной версией библиотеки regexp поддерживающей positive lookahead assertion
func main() {
	var its int
	fmt.Scan(&its)

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	re := regexp.MustCompile(`^(?=(?:[^1]*1){4}[^1]*$)(?=(?:[^2]*2){3}[^2]*$)(?=(?:[^3]*3){2}[^3]*$)(?=(?:[^4]*4){1}[^4]*$)[1-4\s]+$`, regexp.None)
	for j := 0; j < its; j++ {
		st, _ := in.ReadString('\n')
		match, _ := re.MatchString(string(st))
		if match {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}
