package main

import (
	"fmt"
	"os"
	"regexp"
	"github.com/go-git/go-git/v5"
)

// Clones the repo in ~/Project/github/{org}/{repo}
func main() {
	if len(os.Args) < 2 {
		fmt.Print("missing argument\n")
		os.Exit(1)
	}
	origin := os.Args[1]
	repoRegex := regexp.MustCompile("\\W*(\\w+)\\/(\\w+).git")
	matches := repoRegex.FindStringSubmatch(origin)
	org := matches[1]
	repoName := matches[2]

	home, _ := os.UserHomeDir()
	repoPath := fmt.Sprintf("%s/Projects/github/%s/%s", home, org, repoName)

	fmt.Printf("Cloning repo %s into %s\n", origin, repoPath)
	_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:      origin,
		Progress: os.Stdout,
	})
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Error: Already cloned\n")
			fmt.Printf("cd %s\n", repoPath)
			os.Exit(1)
		}
	}
	fmt.Printf("cd %s\n", repoPath)
}