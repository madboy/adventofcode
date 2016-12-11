package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	secret := "reyedfim"
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
	fmt.Printf("The code is: %s\n", code)
}
