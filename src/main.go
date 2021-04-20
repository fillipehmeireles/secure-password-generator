package main

/*
	Secure Random Password Generator

	Author: Fillipe Meireles
	https://github.com/fillipehmeireles/secure-password-generator
*/

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println(generate())
}

func generate() string {
	seed := rand.NewSource(time.Now().UnixNano())
	randNum := rand.New(seed)

	return encodeIt(randNum.Int63())
}

func encodeIt(randomNum int64) string {
	numStr := strconv.Itoa(int(randomNum))
	halfPass := numStr[0:4]
	halfEncrPass := newBase64(halfPass)
	sHanfEncrPass := newMd5((halfPass + numStr[4:len(numStr)-1]))

	return halfEncrPass + sHanfEncrPass
}

func newBase64(data string) string {
	encr := base64.StdEncoding.EncodeToString([]byte(data))
	return encr[0:4]
}

func newMd5(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
