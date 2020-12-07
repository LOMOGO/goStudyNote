package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"log"
	"time"
)

func main() {
	t1 := time.Now()
	salt := "Qjwzg|_|"
	dk, err := scrypt.Key([]byte("some password"), []byte(salt), 8192, 8, 1, 16)
	if err != nil {
		log.Fatal(err)
	}
	dk1 := hex.EncodeToString(dk)
	elapsed := time.Since(t1)

	t2 := time.Now()
	sha1Hash := sha1.New()
	sha1Hash.Write([]byte("some password"))
	dk2 := hex.EncodeToString(sha1Hash.Sum(nil))
	elapsed2 := time.Since(t2)
	fmt.Println(string(dk1), "    time: ", elapsed)
	fmt.Println(string(dk2), "    time: ", elapsed2)
}
