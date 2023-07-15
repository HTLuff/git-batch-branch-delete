package git

import (
	"os/exec"
	"strings"
)

// GetBranches executes the 'git branch' command and returns the output.
func GetBranches(dir string) (string, error) {
	cmd := exec.Command("git", "branch")
	cmd.Dir = dir
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// ParseBranches parses the output of the 'git branch' command and returns a list of branch names.
func ParseBranches(output string) []string {
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

// DeleteBranch executes the 'git branch -D' command to delete the specified branch.
func DeleteBranch(dir, branch string) error {
	cmd := exec.Command("git", "branch", "-D", branch)
	cmd.Dir = dir
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
