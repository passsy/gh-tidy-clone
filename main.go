package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/cli/go-gh"
)

// Clones the repo in ~/Project/github/{org}/{repo}
func main() {
	if len(os.Args) < 2 {
		fmt.Print("missing argument\n")
		os.Exit(1)
	}
	var origin = os.Args[1]
	if !strings.HasSuffix(origin, ".git") {
		origin = fmt.Sprintf("%s.git", origin)
	}
	repoRegex := regexp.MustCompile("\\S*\\/(\\S+)\\/(\\S+)\\.git")
	matches := repoRegex.FindStringSubmatch(origin)
	org := matches[1]
	var repoName = matches[2]
	// remove trailing slash
	repoName = strings.ReplaceAll(repoName, "/", "")

	home, _ := os.UserHomeDir()
	repoPath := fmt.Sprintf("%s/Projects/github/%s/%s", home, org, repoName)

	fmt.Printf("Cloning repo %s into %s\n", origin, repoPath)
	stdOut, _, err := gh.Exec("repo", "clone", fmt.Sprintf("%s/%s", org, repoName), repoPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(stdOut.String())
	fmt.Printf("cd %s\n", repoPath)
}
