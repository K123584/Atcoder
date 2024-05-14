// https://atcoder.jp/contests/abc128/tasks/abc128_a
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// P := A * 3
// applePie := 2 * P

var sc = bufio.NewScanner(os.Stdin)

func IntInput() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	A := IntInput()
	P := IntInput()

	addP := A * 3
	allP := P + addP
	applePie := allP / 2
	fmt.Println(applePie)
}
