package gitbranchbatchdelete

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteBranches(t *testing.T) {
	// Write test cases to verify the behavior of DeleteBranches function
	branches := []string{"branch1", "branch2"}
	err := DeleteBranches(branches)
	assert.NoError(t, err)
}
