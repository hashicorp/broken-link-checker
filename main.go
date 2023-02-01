// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Set up the BrokenLinkChecker
	maxConnections := 0
	i, err := strconv.Atoi(os.Getenv("MAX_CONNECTIONS"))
	if err != nil {
		maxConnections = i
	}
	timeoutSeconds := 0
	i, err = strconv.Atoi(os.Getenv("TIMEOUT_SECONDS"))
	if err != nil {
		timeoutSeconds = i
	}
	brokenLinkChecker := &BrokenLinkChecker{
		Exclusions:     strings.Split(os.Getenv("EXCLUSIONS"), ","),
		MaxConnections: maxConnections,
		Timeout:        timeoutSeconds,
	}

	// Run the broken link checker on the page
	allPages, _ := brokenLinkChecker.Check(os.Args[1])
	filteredPages := filterErrors(allPages)

	// Print all the errors
	verbose, _ := strconv.ParseBool(os.Getenv("VERBOSE"))
	if verbose {
		prettyPrintErrors(allPages)
	} else {
		prettyPrintErrors(filteredPages)
	}

	// Exit with an error if there are any non-dubious errors
	if len(filteredPages) > 0 {
		os.Exit(1)
	}
}

// prettyPrintErrors prints out all of the broken link errors in a CLI friendly manner
func prettyPrintErrors(pages []Page) {
	for _, page := range pages {
		fmt.Fprintf(os.Stderr, "Page: %s\n", page.URL)
		for _, link := range page.BrokenLinks {
			fmt.Fprintf(os.Stderr, "  -> %s: %s\n", link.Error, link.URL)
		}
		fmt.Fprintf(os.Stderr, "\n")
	}
}

// filterErrors filters all dubious errors that don't necessisarily
// mean that there's an actual problem that needs to be addressed.
func filterErrors(pages []Page) []Page {
	filteredPages := []Page{}
	for _, page := range pages {
		filteredPageLinks := []BrokenLink{}
		for _, link := range page.BrokenLinks {
			// Too  many requests
			if link.Error == strconv.Itoa(http.StatusTooManyRequests) {
				continue
			}
			// A strange error code LinkedIn throws when it detects a bot
			if link.Error == "999" {
				continue
			}
			// Likely a large file download, which we are safe to ignore
			if link.Error == "timeout" {
				continue
			}
			filteredPageLinks = append(filteredPageLinks, link)
		}
		if len(filteredPageLinks) > 0 {
			filteredPages = append(filteredPages, page)
		}
	}
	return filteredPages
}
