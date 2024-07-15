package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	length := flag.Int("l", 32, "length of tokens to generate")
	number := flag.Int("n", 1, "number of tokens to generate")
	flag.Parse()

	for i := 0; i < *number; i++ {
		v := randomBase64String(*length)
		fmt.Println(v)
	}
}

func randomBase64String(length int) string {
	b := make([]byte, int(math.Ceil(float64(length*4)/3.0)))
	for {
		if _, err := rand.Read(b); err != nil {
			log.Fatal("failed to generate random bytes", err)
		}
		s := base64.RawStdEncoding.EncodeToString(b)[:length]
		if !strings.ContainsAny(s, "+/") {
			return s
		}
	}
}
