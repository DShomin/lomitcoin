package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"
)

func searcing(start, end int, target string, c chan string) {
	for nonce := start; nonce <= end; nonce++ {
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte("data"+fmt.Sprint(nonce))))
		if strings.HasPrefix(hash, target) {
			c <- hash
			break
		}
	}
	c <- ""
}

func goSearch(nonce, multi_core_num, step_size int, target string, start_time time.Time) {
	start := nonce
	for {
		c := make(chan string)
		for i := 1; i <= multi_core_num; i++ {
			end := start + step_size
			go searcing(start, end, target, c)
			start = end
		}
		for i := 0; i < multi_core_num; i++ {
			goval := <-c
			if goval != "" {
				fmt.Println(goval)
				elapsed := time.Since(start_time)
				fmt.Printf("Time : %s", elapsed)
				return
			}
		}
	}
}

func basicSearch(nonce int, target string, start_time time.Time) {
	for {
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte("data"+fmt.Sprint(nonce))))
		// fmt.Printf("Hash: %s\nTarget: %s\nNonce: %d\n\n", hash, target, nonce)
		if strings.HasPrefix(hash, target) {
			elapsed := time.Since(start_time)
			fmt.Printf("Time : %s", elapsed)
			return
		} else {
			nonce++
		}
	}
}

func main() {
	difficulty := 6
	nonce := 1
	target := strings.Repeat("0", difficulty)
	start_time := time.Now()

	// For go rutine for loop
	goSearch(nonce, 64, 100, target, start_time)

	// For basic loop
	// basicSearch(nonce, target, start_time)
}
