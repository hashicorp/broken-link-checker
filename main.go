package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	brokenLinkChecker := &BrokenLinkChecker{
		MaxConnections: 5,
	}

	pages, _ := brokenLinkChecker.Check("https://waypointproject.io")

	b, _ := json.Marshal(pages)
	fmt.Println(string(b))
}
