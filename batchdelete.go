package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Get the current working directory.
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get the current working directory: %v", err)
	}

	// Execute the Git command to get the list of branches.
	cmd := exec.Command("git", "branch")
	cmd.Dir = dir
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to execute the 'git branch' command: %v", err)
	}

	// Parse the output and display the list of branches.
	branches := parseBranches(string(output))
	if len(branches) == 0 {
		fmt.Println("No branches found.")
	} else {
		fmt.Println("Branches:")
		for _, branch := range branches {
			fmt.Println(branch)
		}
	}
}

// parseBranches parses the output of the 'git branch' command and returns a list of branch names.
func parseBranches(output string) []string {
	var branches []string

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		branchName := strings.TrimPrefix(line, "*")
		branches = append(branches, strings.TrimSpace(branchName))
	}

	return branches
}
