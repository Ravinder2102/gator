# GATOR

## Description

RSS feed aggregator

Add RSS feeds from across the internet to be collected

Store the collected posts in a PostgreSQL database

Follow and unfollow RSS feeds that other users have added

View summaries of the aggregated posts in the terminal, with a link to the full post


## Installation 

### Prerequisites
You need to have Postgres and Go installed to run gator

To install gator use this in your terminal
```bash
go install github.com/Ravinder2102/gator@latest
```
Make a `~/.gatorconfig.json` file in your home directory
inside the config file add these contents
```json
{
  "db_url": "postgres://user:password@localhost:5432/database?sslmode=disable",
  "current_user_name": "username"
}
```
Replace `user` with `postgres` and `password` with your Postgres password

## Usage

Some of the commands you can use with gator are:

### login
```bash 
gator login <username>
```
Sets the current user in the config.

### register
```bash 
gator register <username>
```
Adds a new user to the database.

### reset
```bash 
gator reset
```
**Warning:** `gator reset` deletes all database data.
Reset the database.

### users
```bash
gator users
```
Lists all the users in the database.

### agg
```bash 
gator agg <timeBetweenRequests>
```
Aggregate and save posts from all feeds fetching on intervals specified by you.

### addfeed
```bash 
gator addfeed <feedName> <url>
```
Add feed to feeds database.

### feeds
```bash 
gator feeds
```
Lists all feeds in the feeds database.

### follow
```bash 
gator follow <url>
```
Follow a feed for the current user.

### following
```bash
gator following
```
Lists all the feeds being followed by current user.

### unfollow
```bash
gator unfollow <url>
```
Unfollow feed for current user.

### browse
```bash 
gator browse [limit]
```
List all the posts fetched for the feeds being followed by the current user in most recent order.
Specify limit to see a limited number of posts. 
If no limit is specified, the default limit is 2.
