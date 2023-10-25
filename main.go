package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/github"
)

func main() {
	var prNumber int
	flag.IntVar(&prNumber, "pr-number", 0, "The number of the Agent Pool controller workers.")
	flag.Parse()

	if prNumber == 0 {
		fmt.Println("PR number is not passed")
		os.Exit(1)
	}

	fmt.Println("PR number:", prNumber)

	r, ok := os.LookupEnv("GITHUB_REPOSITORY")
	if !ok {
		fmt.Println("Failed to get the repository via GITHUB_REPOSITORY")
		os.Exit(1)
	}
	rr := strings.Split(r, "/")

	client := github.NewClient(nil)
	files, _, err := client.PullRequests.ListFiles(context.TODO(), rr[0], rr[1], prNumber, &github.ListOptions{})
	if err != nil {
		fmt.Println(fmt.Errorf("Failed to get a file list: %s", err.Error()))
		os.Exit(1)
	}

	for _, f := range files {
		fmt.Println(*f.Filename)
	}
}
