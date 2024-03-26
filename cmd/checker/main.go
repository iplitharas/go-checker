package main

import "fmt"

var bannerText = `
__        __   _        ____ _               _             
\ \      / /__| |__    / ___| |__   ___  ___| | _____ _ __ 
 \ \ /\ / / _ \ '_ \  | |   | '_ \ / _ \/ __| |/ / _ \ '__|
  \ V  V /  __/ |_) | | |___| | | |  __/ (__|   <  __/ |   
   \_/\_/ \___|_.__/   \____|_| |_|\___|\___|_|\_\___|_|   
`[1:]

func main() {
	fmt.Println("Checker entrypoint!")
	fmt.Println(bannerText)
}
