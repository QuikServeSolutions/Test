package stringme

import "strings"

func LeftMe(s string) string {
	return strings.TrimLeft(s, "te")
}
