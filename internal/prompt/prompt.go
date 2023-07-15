package prompt

import (
	"fmt"

	survey "github.com/AlecAivazis/survey/v2"
)

// SelectBranches prompts the user to select branches from a checklist.
func SelectBranches(branches []string) ([]string, error) {
	prompt := &survey.MultiSelect{
		Message: "Select branches to delete:",
		Options: branches,
	}
	var branchSelection []string
	if err := survey.AskOne(prompt, &branchSelection); err != nil {
		return nil, err
	}
	return branchSelection, nil
}

// ConfirmDeletion prompts the user to confirm branch deletion.
func ConfirmDeletion(branches []string) bool {
	confirm := false
	confirmPrompt := &survey.Confirm{
		Message: fmt.Sprintf("Are you sure you want to delete the selected branches: %s?", branches),
	}
	if err := survey.AskOne(confirmPrompt, &confirm); err != nil {
		return false
	}
	return confirm
}
