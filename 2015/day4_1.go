package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func leadingZeros(start []byte) bool {
	tails := fmt.Sprintf("0x%x", start[2])
	tail, _ := strconv.ParseInt(tails, 0, 64)
	if (fmt.Sprintf("%x", start[0:2]) == "0000") && (tail < 10) {
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
	for i := 0; i < 2000000; i++ {
		digest := getDigest(secret, i)
		if leadingZeros(digest[0:3]) {
			fmt.Printf("number is %d and gives the digest %x\n", i, digest)
			break
		}
	}
}
