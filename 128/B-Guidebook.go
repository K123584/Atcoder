package main

import (
	"fmt"
	"sort"
)

type ct struct {
	Rank  int
	Key   string
	Value int
}

func main() {
	var n int
	fmt.Scan(&n)

	var s string
	var p int
	c := make([]ct, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s, &p)
		c[i] = ct{i + 1, s, p}
	}

	sort.SliceStable(c, func(i, j int) bool {
		return c[i].Key < c[j].Key
	})

	sort.SliceStable(c, func(i, j int) bool {
		if c[i].Key == c[j].Key {
			return c[i].Value > c[j].Value
		}
		return false
	})

	for _, ct := range c {
		fmt.Println(ct.Rank)
	}
}
