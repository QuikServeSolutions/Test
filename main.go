package main

import (
	"fmt"

	"github.com/terryberlin/test/stringme"
)

// "crypto/sha256"
// "encoding/hex"
// func hash(c string) string {
// 	h := sha256.New()
// 	fmt.Fprint(h, c)
// 	return hex.EncodeToString(h.Sum(nil))
//  fmt.Println(hash("tjb"))
// }

func main() {
	fmt.Println(stringme.UpperMe("terry"))
	fmt.Println(stringme.LowerMe("terry"))
	fmt.Println(stringme.LeftMe("terry"))
}
