# Gator CLI

## Introduction
Gator CLI is a command-line tool for managing and aggregating blog feeds. It allows users to register, log in, add feeds, follow/unfollow feeds, and browse posts from their followed feeds.

## Setup

### 1. Install Go
- **Using Homebrew**:
    ```sh
    brew install go
    ```
- **Using Official Sites**:
    - Download the Go installer from the [official Go website](https://golang.org/dl/).
    - Follow the installation instructions for your operating system.

### 2. Install PostgreSQL
- **Using Homebrew**:
    ```sh
    brew install postgresql
    ```
- **Using Official Sites**:
    - Download the PostgreSQL installer from the [official PostgreSQL website](https://www.postgresql.org/download/).
    - Follow the installation instructions for your operating system.

### 3. Clone and Configure the Repository
1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/gator-cli.git
    cd gator-cli
    ```
2. Set up the PostgreSQL database:
    ```sh
    psql -U postgres -c "CREATE DATABASE gator;"
    ```
3. Run the database migrations using Goose:
    ```sh
    goose -dir sql/schema postgres "user=your_postgres_username password=your_postgres_password dbname=gator sslmode=disable" up
    ```

### 4. Automatic Configuration Script
To make setup easier, you can run the following script to configure the Gator CLI tool, set up the required database, and register/login a user in your local database:

```bash
#!/bin/bash

current_user_name=""
db_url=""
next=false

# Prompt for username
while [ "$next" = false ]; do
    echo "Please enter your username:"
    read current_user_name
    echo "You entered: $current_user_name. Accept [y]/[n]?"
    read confirm
    if [ "$confirm" = "y" ]; then
        next=true
    fi
done

# Prompt for database URL
next=false
while [ "$next" = false ]; do
    echo "Please enter your PostgreSQL connection string (e.g., postgres://your-username:@localhost:5432/gator):"
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

# Install Gator command-line tool
echo "Installing gator command-line tool."
go install github.com/Brent-the-carpenter/gator@latest

# Ensure installation was successful
if ! command -v gator &> /dev/null; then
    echo "Error: gator could not be installed or is not in your PATH. Please check your Go installation."
    exit 1
fi

# Register user in the local database
echo "Registering ${current_user_name} in your local database"
gator register "${current_user_name}"

# Log in user
echo "Logging in ${current_user_name}"
gator login "${current_user_name}"
```
Make sure to make the script executable by running:
```sh
chmod +x setup_gator.sh
```
## Build and Run

If you didn’t run the setup script above, you’ll need to manually build the Gator CLI tool before running it. 

1. **Build the CLI Tool**:
    ```sh
    go build -o gator
    ```
2. **Run the CLI Tool**:
    ```sh
    ./gator
    ```

Alternatively, if you have Go installed and prefer installing `gator` directly, you can use:
```sh
go install github.com/Brent-the-carpenter/gator@latest
```
## Usage

### Available Commands
- **User Management**:
  - `register <username>`: Register a new user.
  - `login <username>`: Log in as an existing user.
  - `users`: List all registered users.
  - `reset`: Reset all user data in the database.

- **Feed Management**:
  - `addfeed <name> <url>`: Add a new feed and follow it.
  - `feeds`: List all available feeds.
  - `follow <url>`: Follow an existing feed.
  - `following`: List all feeds followed by the current user.
  - `unfollow <url>`: Unfollow a feed.

- **Post Aggregation**:
  - `agg <interval>`: Aggregate posts from RSS feeds at specified intervals (e.g., `10m` for 10 minutes).
  - `browse <limit>`: Browse posts from feeds followed by the current user, with an optional limit.

### Command Examples
- **Register a new user**:
    ```sh
    ./gator register username
    ```
- **Log in as an existing user**:
    ```sh
    ./gator login username
    ```
- **Add a new feed and follow it**:
    ```sh
    ./gator addfeed "feedname" "feedurl"
    ```
- **List all feeds**:
    ```sh
    ./gator feeds
    ```
- **Follow an existing feed**:
    ```sh
    ./gator follow feedurl
    ```
- **List all feeds followed by the current user**:
    ```sh
    ./gator following
    ```
- **Unfollow a feed**:
    ```sh
    ./gator unfollow feedurl
    ```
- **Aggregate posts from RSS feeds at intervals**:
    ```sh
    ./gator agg 10m
    ```
- **Browse posts for the current user**:
    ```sh
    ./gator browse
    ```
- **Reset database users**:
    ```sh
    ./gator reset
    ```

## License
This project is licensed under the MIT License.
