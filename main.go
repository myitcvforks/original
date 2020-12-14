package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		panic("Could not read build info")
	}
	fmt.Printf("Hello, world! We are version %v\nBut I added some new features!\n", bi.Main.Version)
}
