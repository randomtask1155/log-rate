package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	logLines int
	logSize  int
)

func init() {
	l := os.Getenv("LOG_LINES")
	s := os.Getenv("LOG_SIZE")

	if l == "" || s == "" {
		panic("LOG_LINES and LOG_SIZE environment variables required")
	}

	var lerr, serr error
	logLines, lerr = strconv.Atoi(l)
	logSize, serr = strconv.Atoi(s)
	if lerr != nil || serr != nil {
		fmt.Printf("logLines: %s\nlogSize: %s\n", lerr, serr)
		panic("can not convert env vars from string to int")
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func emitLine() {
	fmt.Fprintf(os.Stdout, "%s\n", RandStringRunes(logSize))
}

func main() {

	lineInterval := (time.Second / time.Duration(int64(logLines)))
	for {
		for i := lineInterval; i <= time.Second; i += lineInterval {
			emitLine()
			time.Sleep(lineInterval)
		}
	}

}
