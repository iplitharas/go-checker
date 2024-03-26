package main

import "fmt"

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
	fmt.Println(banner())
	fmt.Println(usage())
}
