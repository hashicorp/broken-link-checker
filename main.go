package main

import (
	"fmt"
)

func main() {
	brokenLinkChecker := &BrokenLinkChecker{
		MaxConnections: 10,
	}

	pages, _ := brokenLinkChecker.Check("https://waypointproject.io")
	fmt.Println(pages)
}
