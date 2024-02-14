package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkProcess(sequence string) string {
	hasStarted := false
	hasFinished := false
	hasRestarted := false
	hasCanceled := false

	for _, event := range sequence {
		switch event {
		case 'M':
			if (hasStarted || hasRestarted) && !hasCanceled {
				return "NO"
			}
			hasStarted = true
			hasFinished, hasRestarted, hasCanceled = false, false, false
		case 'R':
			if !hasStarted || hasRestarted {
				return "NO"
			}
			hasRestarted = true
			hasStarted = false
		case 'C':
			if (!hasStarted && !hasRestarted) || hasCanceled {
				return "NO"
			}
			hasRestarted = false
			hasCanceled = true
		case 'D':
			if !hasStarted || hasRestarted || hasFinished {
				return "NO"
			}
			hasFinished = true
			hasStarted, hasRestarted, hasCanceled = false, false, false
		}
	}

	if !hasFinished || hasCanceled {
		return "NO"
	}

	return "YES"
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var num int
	fmt.Fscan(in, &num)

	for i := 0; i < num; i++ {
		var seq string
		fmt.Fscan(in, &seq)
		result := checkProcess(seq)
		fmt.Fprint(out, result)
		fmt.Fprint(out, "\n")
	}
}
