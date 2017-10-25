package main

import strs "strings"
import "fmt"

var logStdOut = fmt.Println

func main() {
	logStdOut("Contains", strs.Contains("test", "es"))
	logStdOut("Count", strs.Count("test", "t"))
	logStdOut("Has prefix", strs.HasPrefix("test", "te"))
	logStdOut("Has suffix", strs.HasSuffix("test", "st"))
	logStdOut("Index", strs.Index("test", "e"))
	logStdOut("Joins", strs.Join([]string{"b", "a", "t", "u"}, "-"))
	logStdOut("Repeat", strs.Repeat("a", 5))
	logStdOut("Replace all", strs.Replace("foo", "o", "0", -1))
	logStdOut("Replace once", strs.Replace("foo", "o", "0", 1))
	logStdOut("Split", strs.Split("a,b,c,d,e", ","))
	logStdOut("ToLower", strs.ToLower("TEST"))
	logStdOut("ToUpper", strs.ToUpper("test"))
	logStdOut("Len", len("hello"))
	logStdOut("Char:", "hello"[1])
}
