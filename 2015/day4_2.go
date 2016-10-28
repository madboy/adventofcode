package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func leadingZeros(start []byte) bool {
	if fmt.Sprintf("%x", start[0:3]) == "000000" {
		return true
	} else {
		return false
	}
}

func getDigest(secret string, number int) [16]byte {
	data := []byte(secret + strconv.Itoa(number))
	return md5.Sum(data)
}

func main() {
	secret := "iwrupvqb"
	for i := 0; i < 10000000; i++ {
		digest := getDigest(secret, i)
		if leadingZeros(digest[0:3]) {
			fmt.Printf("number is %d and gives the digest %x\n", i, digest)
			break
		}
	}
}
