package stringme

import (
	"strings"
)

func UpperMe(s string) string {
	// fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	// fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	// fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	// fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	return strings.ToUpper(s)

}
