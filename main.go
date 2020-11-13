package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	brokenLinkChecker := &BrokenLinkChecker{
		MaxConnections: 5,
	}

	pages, _ := brokenLinkChecker.Check(os.Args[1])
	b, _ := json.Marshal(pages)
	fmt.Println(string(b))
}
