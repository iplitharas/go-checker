package main

import (
	"flag"
	"runtime"
)

type flags struct {
	url  string
	n, c int
}

type parseFunc func(string) error

func (f *flags) parse() error {
	var (
		u = flag.String("url", "", "HTTP server URL to make requests. (required)")
		n = flag.Int("n", 100, "Number of requests to make")
		c = flag.Int("c", runtime.NumCPU(), "Concurrency level")
	)
	flag.Parse()
	f.url = *u
	f.n = *n
	f.c = *c
	return nil

}
