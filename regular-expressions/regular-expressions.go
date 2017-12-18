package main

import (
	"regexp"
	"fmt"
	"bytes"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println("Match", match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Printf("%t\n", r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))

	var str = "peach punch"
	loc := r.FindStringIndex(str)
	substr := str[loc[0]:loc[1]]
	fmt.Println(substr)

	fmt.Println(r.FindStringSubmatch(str))

	fmt.Println(r.FindStringSubmatchIndex(str))

	fmt.Println(r.FindAllString(str+" pinch", -1))
	fmt.Println(r.FindAllString(str+" pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch") // constant regular exp
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)

	fmt.Println(string(out))
}
