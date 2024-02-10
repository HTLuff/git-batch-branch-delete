package batchdelete

import (
	"log"
	"os"

	"github.com/HTLuff/git-branch-batch-delete/internal/git"
	"github.com/HTLuff/git-branch-batch-delete/internal/prompt"
)

// DeleteBranches prompts the user to select branches for deletion and deletes them.
func DeleteBranches() {
	// Get the current working directory.
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get the current working directory: %v", err)
	}

	// Execute the Git command to get the list of branches.
	branchOutput, err := git.GetBranches(dir)
	if err != nil {
		log.Fatalf("Failed to execute the 'git branch' command: %v", err)
	}

	// Parse the output and display the list of branches.
	branches := git.ParseBranches(branchOutput)
	if len(branches) == 0 {
		log.Println("No branches found.")
		return
	}

	// Prompt the user to select branches for deletion.
	branchSelection, err := prompt.SelectBranches(branches)
	if err != nil {
		log.Fatalf("Failed to read user input: %v", err)
	}

	if len(branchSelection) == 0 {
		log.Println("No branches selected.")
		return
	}

	// Confirm branch deletion.
	confirm := prompt.ConfirmDeletion(branchSelection)
	if !confirm {
		log.Println("Deletion canceled.")
		return
	}

	// Delete the selected branches.
	for _, branch := range branchSelection {
		err := git.DeleteBranch(dir, branch)
		if err != nil {
			log.Printf("Failed to delete branch '%s': %v", branch, err)
		} else {
			log.Printf("Deleted branch '%s'.", branch)
		}
	}
}
