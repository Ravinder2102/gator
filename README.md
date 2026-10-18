# GATOR

A command-line RSS feed aggregator written in Go. Gator collects osts from RSS feeds across the internet, stores them in a PostgreSQL database, and lets you read posts in your terminal.



## Features

- Add RSS feeds from the internet
- Collect and store posts in a PostgreSQL database
- Follow and unfollow feeds that other users have added
- Browse post in the terminal, with a link to the full post
- Multiple users, each with their own followed feeds

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or later
- [PostgreSQL](https://www.postgresql.org/download/) 15 or later
- [goose](https://github.com/pressly/goose) for database migrations:


```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

## Installation

Install gator with:
 
```bash
go install github.com/Ravinder2102/gator@latest
```

## Setup

### 1. Create the database
 
```bash
psql -U postgres -c "CREATE DATABASE gator;"
```
 
### 2. Run the migrations
 
From the root of the cloned repository:
 
```bash
git clone https://github.com/Ravinder2102/gator.git
cd gator
goose -dir sql/schema postgres "postgres://postgres:YOUR_PASSWORD@localhost:5432/gator?sslmode=disable" up
```

### 3. Create the config file
 
Create a file named `.gatorconfig.json` in your home directory (`~/.gatorconfig.json`):
 
```json
{
  "db_url": "postgres://postgres:YOUR_PASSWORD@localhost:5432/gator?sslmode=disable"
}
```
 
Replace `postgres` with your Postgres username if it differs, and `YOUR_PASSWORD` with your Postgres password. The `current_user_name` field is added automatically when you run `gator register` or `gator login`.

## Quick Start
 
```bash
# Create your user (this also logs you in)
gator register alice
 
# Add a feed
gator addfeed "Boot.dev Blog" https://blog.boot.dev/index.xml
 
# In a second terminal, start collecting posts every minute
gator agg 1m
 
# Back in the first terminal, read the latest posts
gator browse 5
```

## Commands

### login
```bash 
gator login <username>
```
Switch the current user

### register
```bash 
gator register <username>
```
Create a new user and log in as them

### users
```bash
gator users
```
List all users

### agg
```bash 
gator agg <timeBetweenRequests>
```
Fetch and save posts from all feeds on an interval (e.g. 30s, 1m, 1h)

### addfeed
```bash 
gator addfeed <feedName> <url>
```
Add a feed and follow it.

### feeds
```bash 
gator feeds
```
List all feeds in the database

### follow
```bash 
gator follow <url>
```
Follow an existing feed as the current user

### following
```bash
gator following
```
List the feeds the current user follows

### unfollow
```bash
gator unfollow <url>
```
Unfollow a feed as the current user.

### browse
```bash 
gator browse [limit]
```
Show the most recent posts from followed feeds (default limit: 2)

### reset
```bash 
gator reset
```
Delete all data from the database

**Warning:** `gator reset` permanently deletes all users, feeds, and posts. Use it only in development.