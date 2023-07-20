#!/bin/bash

# Clone the gitbatchdelete repository
git clone https://github.com/HTLuff/git-branch-batch-delete.git

# Change to the project directory
cd git-branch-batch-delete

# Build the project and create the binary
go build -o gitbranchbatchdelete ./cmd/main.go

# Move the binary to a location in the $PATH (e.g., /usr/local/bin)
sudo mv gitbranchbatchdelete /usr/local/bin/

# Print installation completion message
echo "Installation completed. You can now use 'gitbranchbatchdelete' in your terminal."
