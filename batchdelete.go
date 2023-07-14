package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	survey "github.com/AlecAivazis/survey/v2"
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
		return
	}

	// Prepare the prompt for branch selection.
	prompt := &survey.MultiSelect{
		Message: "Select branches to delete:",
		Options: branches,
	}

	// Prompt the user to select branches for deletion.
	var branchSelection []string
	if err := survey.AskOne(prompt, &branchSelection); err != nil {
		log.Fatalf("Failed to read user input: %v", err)
	}

	if len(branchSelection) == 0 {
		fmt.Println("No branches selected.")
		return
	}

	// Confirm branch deletion.
	confirm := false
	confirmPrompt := &survey.Confirm{
		Message: fmt.Sprintf("Are you sure you want to delete the selected branches: %s?", strings.Join(branchSelection, ", ")),
	}
	if err := survey.AskOne(confirmPrompt, &confirm); err != nil {
		log.Fatalf("Failed to read user input: %v", err)
	}

	if !confirm {
		fmt.Println("Deletion canceled.")
		return
	}

	// Delete the selected branches.
	for _, branch := range branchSelection {
		cmd := exec.Command("git", "branch", "-D", branch)
		cmd.Dir = dir
		if err := cmd.Run(); err != nil {
			log.Printf("Failed to delete branch '%s': %v", branch, err)
		} else {
			fmt.Printf("Deleted branch '%s'.\n", branch)
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
