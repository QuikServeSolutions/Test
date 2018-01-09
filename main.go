package main

import (
	"fmt"

	"github.com/terryberlin/test/upperme"
)

//import "github.com/terryberlin/test/upperme"
//"github.com/quikServesolutions/test/upperme"

// "crypto/sha256"
// "encoding/hex"
// func hash(c string) string {
// 	h := sha256.New()
// 	fmt.Fprint(h, c)
// 	return hex.EncodeToString(h.Sum(nil))
//  fmt.Println(hash("tjb"))
// }

func main() {
	fmt.Println(upperme.UpperMe("terry"))
	fmt.Println(upperme.LowerMe("terry"))
}
