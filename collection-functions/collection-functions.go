package main

import (
	"fmt"
	"strings"
)

func Index(t []string, e string) int {
	for i, v := range t {
		if v == e {
			return i
		}
	}
	return -1
}

func Include(vs []string, v string) bool {
	return Index(vs, v) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	var stats bool = false
	for _, e := range vs {
		if f(e) {
			stats = true
			break
		}
	}
	return stats
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	fvs := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			fvs = append(fvs, v)
		}
	}
	return fvs
}

func Map(vs []string, f func(string) string) []string {
	mvs := make([]string, len(vs))

	for i, v := range vs {
		mvs[i] = f(v)
	}
	return mvs
}

func main() {

	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println("apple index :", Index(strs, "apple"))

	fmt.Println("Include plum:", Include(strs, "plum"))

	fmt.Println(Any(strs, func (v string) bool {
		return strings.HasPrefix(v, "z")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.Contains(v, "p")
	}))

	fmt.Println(Filter(strs, func(i string) bool {
		return strings.Contains(i, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))
}
