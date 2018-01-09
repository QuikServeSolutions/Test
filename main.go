 package main

 import (
     "crypto/sha256"
     "fmt"
     "encoding/hex"
 )

 func hash(c string) string {
         h := sha256.New()
         fmt.Fprint(h, c)
         return hex.EncodeToString(h.Sum(nil))
 }


 func main() {
     fmt.Println(hash("tjb"))
 }