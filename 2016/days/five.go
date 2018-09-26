package days

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

//Run5 in which we look for more codes
func Run5(scanner *bufio.Scanner) string {
	secret := "ojvtpuvg"
	// secret := "abc"
	count := 0
	code := [8]string{"-", "-", "-", "-", "-", "-", "-", "-"}

	for i := 0; i < 40000000 && count < 8; i++ {
		data := []byte(secret + strconv.Itoa(i))
		digest := md5.Sum(data)
		hash := fmt.Sprintf("%x", digest)

		if hash[0:5] == "00000" {
			p, err := strconv.Atoi(string(hash[5]))
			if err != nil {
				continue
			}
			if p >= 0 && p < 8 {
				if string(code[p]) == "-" {
					code[p] = string(hash[6])
					count++
				}

			}
		}
	}
	return fmt.Sprintf("The code is: %s", strings.Join(code[:], ""))
}
