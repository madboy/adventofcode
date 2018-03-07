package days

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"strconv"
)

// Run4 is the fourth day of christmas
func Run4(scanner *bufio.Scanner) string {
	secret := "ckczppom"
	return fmt.Sprintf("%s\n%s", numberFiveZeros(secret), numberSixZeros(secret))
}

func numberFiveZeros(secret string) string {
	for i := 0; i < 2000000; i++ {
		data := []byte(secret + strconv.Itoa(i))
		digest := md5.Sum(data)
		hash := fmt.Sprintf("%x", digest)

		if hash[0:5] == "00000" {
			return fmt.Sprintf("Number is %d and gives the digest %x", i, digest)
		}
	}
	return ""
}

func numberSixZeros(secret string) string {
	for i := 0; i < 10000000; i++ {
		data := []byte(secret + strconv.Itoa(i))
		digest := md5.Sum(data)
		hash := fmt.Sprintf("%x", digest)

		if hash[0:6] == "000000" {
			return fmt.Sprintf("Number is %d and gives the digest %x", i, digest)
		}
	}
	return ""
}
