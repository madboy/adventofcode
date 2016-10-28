package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	secret := "iwrupvqb"

	for i := 0; i < 2000000; i++ {
		data := []byte(secret + strconv.Itoa(i))
		digest := md5.Sum(data)
		hash := fmt.Sprintf("%x", digest)

		if hash[0:5] == "00000" {
			fmt.Printf("number is %d and gives the digest %x\n", i, digest)
			break
		}
	}
}
