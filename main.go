package main

import "fmt"

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

func main() {
	fmt.Println("Hello GoReleaser")
	fmt.Printf("Version: %s, commit: %s, date: %s, by: %s\n", version, commit, date, builtBy)
}
