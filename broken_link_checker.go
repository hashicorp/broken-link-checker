package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Page struct {
	URL         string       `json:"url"`
	BrokenLinks []BrokenLink `json:"links"`
}

type BrokenLink struct {
	Error string `json:"error"`
	URL   string `json:"url"`
}

// BrokenLinkChecker is a simple wrapper around the muffet CLI, as muffet
// does not provide a public go interface. Not all muffet parameters are
// exposed here, as this is an intentional simplification.
//
// https://github.com/raviqqe/muffet
type BrokenLinkChecker struct {
	Exclusions     []string
	MaxConnections int
	Timeout        int
}

// Check checks for broken links against a specific URL
func (c *BrokenLinkChecker) Check(url string) ([]Page, error) {
	// Generate the Params
	var params = []string{}
	if c.MaxConnections != 0 {
		params = append(params, fmt.Sprintf("--max-connections=%d", c.MaxConnections))
	}
	if c.Timeout != 0 {
		params = append(params, fmt.Sprintf("--timeout=%d", c.Timeout))
	}
	for _, exclusion := range c.Exclusions {
		params = append(params, fmt.Sprintf("--exclude=.*%s.*", exclusion))
	}
	params = append(params, "--json")
	params = append(params, url)

	// Execute the command
	commandOutput, err := exec.Command("muffet", params...).Output()
	if err != nil {
		// Consume an error exit status, this is expected and a non-error at this point
		if err.Error() != "exit status 1" {
			return nil, err
		}
	}

	// Parse out the pages from the response
	pageArray := []Page{}
	err = json.Unmarshal(commandOutput, &pageArray)
	if err != nil {
		return nil, err
	}
	return pageArray, nil
}
