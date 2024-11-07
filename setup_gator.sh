#!/bin/bash

# Prompt for the Github username
echo "Please enter your Github username (for your fork of gator):"
read github_username

current_user_name=""
db_url=""
next=false
# Prompt for username
while [ "$next" = false ]; do
    echo "Please Enter your username:"
    read current_user_name
    echo "you entered: $current_user_name. Accept [y]/[n]?"
    read confirm
    if [ "$confirm" = "y" ]; then
        next=true
    fi
done

# Prompt for database URL
next=false
while [ "$next" = false ]; do
    echo "Please enter your postgres connection string (e.g., postgres://your-username:@localhost:5432/gator):"
    read db_url
    echo "You entered: $db_url. Accept [y]/[n]?"
    read confirm
    if [ "$confirm" = "y" ]; then
        next=true
    fi
done

# Save the configuration to ~/.gatorconfig.json
echo "{\"db_url\":\"${db_url}?sslmode=disable\",\"current_user_name\":\"${current_user_name}\"}" > ~/.gatorconfig.json
if [ -f ~/.gatorconfig.json ]; then
    echo "Configuration saved to ~/.gatorconfig.json"
else
    echo "Error: Configuration file could not be saved."
    exit 1
fi


#Install gator command-line tool from user's fork
echo "Installing gator command-line tool from github.com/${github_username}/gator."
go install github.com/${github_username}/gator@latest

# Ensure installation was successful
if [ $? -ne 0 ]; then
    echo "Error: gator installation failed. Please check your Go setup."
    exit 1
fi

# Ensure installation was successful
if ! command -v gator &> /dev/null; then
    echo "Error: gator could not be installed or is not in your PATH. Please check your Go installation."
    exit 1
fi

# Register user in their local database
echo "Registering ${current_user_name} in your local database"
gator register "${current_user_name}"

#Login user
echo "Logging in ${current_user_name}"
gator login "${current_user_name}"
