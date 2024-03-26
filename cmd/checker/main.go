package main

import (
	"fmt"
	"log"
)

const (
	bannerText = `
__        __   _        ____ _               _             
\ \      / /__| |__    / ___| |__   ___  ___| | _____ _ __ 
 \ \ /\ / / _ \ '_ \  | |   | '_ \ / _ \/ __| |/ / _ \ '__|
  \ V  V /  __/ |_) | | |___| | | |  __/ (__|   <  __/ |   
   \_/\_/ \___|_.__/   \____|_| |_|\___|\___|_|\_\___|_|   
`
	usageText = `
Usage:
  -url
       HTTP server URL to make requests (required)
  -n
       Number of requests to make
  -c
       Concurrency level
`
)

func banner() string { return bannerText[1:] }
func usage() string  { return usageText[1:] }

func main() {
	fmt.Println("Checker entrypoint!")

	var f flags
	if err := f.parse(); err != nil {
		fmt.Println(usage())
		log.Fatal(err)
	}
	fmt.Println(banner())
	fmt.Printf("Making %d requests to %s with a concurrency level of %d.\n",
		f.n, f.url, f.c)
}
