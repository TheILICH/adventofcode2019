package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// "io"
	// "sort"
	// "strings"
	"strconv"
	// "container/list"
)

const (
	inf int = 1e18
	mod int = 998244353
	N   int = 500
	M   int = 500
)

var (

)

func solve() {

	var a []int
	end := false
	for !end {
		line, err := r.ReadString('\n')
		if err != nil && err.Error() == "EOF" {
			end = true
		}

		n, _ := strconv.Atoi(strings.TrimSpace(line))
		a = append(a, n)
	}

	res := 0
	for _, n := range a {
		res += n / 3 - 2
	}

	fprintf("%d\n", res)

}

func star() {

	var a []int
	end := false
	for !end {
		line, err := r.ReadString('\n')
		if err != nil && err.Error() == "EOF" {
			end = true
		}

		n, _ := strconv.Atoi(strings.TrimSpace(line))
		a = append(a, n)
	}

	res := 0
	for _, n := range a {
		for n > 0 {
			n = n / 3 - 2
			if n <= 0 {
				break
			}
			res += n
		}
	}

	fprintf("%d\n", res)

}

var (
	w *bufio.Writer
	r *bufio.Reader
)

func fscan(a ...any) {
	fmt.Fscan(r, a...)
}

func fscanf(format string, a ...any) {
	fmt.Fscanf(r, format, a...)
}

func fprintf(format string, a ...any) {
	fmt.Fprintf(w, format, a...)
}

func main() {

	// r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)

	fin, _ := os.Open("f.in")
	defer fin.Close()
	r = bufio.NewReader(fin)

	// fout, _ := os.Open("f.out")
	// defer fout.Close()
	// w = bufio.NewWriter(fout)

	defer w.Flush()

	tt := 1
	// fscan(&tt)

	for i := 0; i < tt; i++ {
		// solve()
		star()
	}

}
