#!/bin/bash

# Clone the gitbatchdelete repository
git clone https://github.com/htluff/gitbatchdelete.git

# Install the module using go get
go get -u github.com/htluff/gitbatchdelete

# Add the Go binary directory to the PATH
if [[ "$OSTYPE" == "darwin"* ]]; then
  echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bash_profile
elif [[ "$OSTYPE" == "linux"* ]]; then
  echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
elif [[ "$OSTYPE" == "msys" ]]; then
  echo 'export PATH=$PATH:/c/Go/bin' >> ~/.bashrc
fi

# Print the installation completion message
echo "gitbatchdelete installed successfully. Please restart your terminal or run 'source ~/.bashrc' to update the environment variables."
