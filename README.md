# The Blogegator 

The Blogegator is a CLI tool that will aggregate RSS feeds, allowing multiple users for follow desired feeds.

## Requirements

In order to use The Blogegator, you will need to download and install Go and PostgreSQL.

To install Go follow the instructions [here](https://go.dev/doc/install).

(*PostgreSQL installation and set up instructions adopted from [Boot.dev](https://www.boot.dev/lessons/74bea1f2-19cd-4ea9-966e-e2ca9dd1dfa9))

To use PostgreSQL, you will install the server and start it. Then, you can connect to it using a client like [psql](https://www.postgresql.org/docs/current/app-psql.html#:~:text=psql%20is%20a%20terminal%2Dbased,or%20from%20command%20line%20arguments.) or [PGAdmin](https://www.pgadmin.org/).

**Mac OS with brew**

```
brew install postgresql@15
```

**Linux / WSL (Debian). Here are the docs from Microsoft, but simply:**

```
sudo apt update
sudo apt install postgresql postgresql-contrib
```

Ensure the installation worked. The `psql` command-line utility is the default client for Postgres. Use it to make sure you're on version 15+ of Postgres:
```
psql --version
```

(Linux only) Update postgres password:
```
sudo passwd postgres
```

Enter a password, and be sure you won't forget it. You can just use something easy like `postgres`.

## Set Up

Clone this repo and start a terminal in its root directory.

### PostgreSQL 
Start the Postgres server in the background
- Mac: brew services start postgresql
- Linux: sudo service postgresql start

Connect to the server. Enter the psql shell:

- Mac: `psql postgres`
- Linux: `sudo -u postgres psql`

You should see a new prompt that looks like this:

```
postgres=#
```

Create a new database. I called mine gator:
```
CREATE DATABASE gator;
```

Connect to the new database:
```
\c gator
```

You should see a new prompt that looks like this:
```
gator=#
```

Set the user password (Linux only)
```
ALTER USER postgres PASSWORD 'your_password';
```

### Schema

Install goose to handle establishing the database schema
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Determine your connection string to the Postgres database. The format is:
```
protocol://username:password@host:port/database
```

It should look something like this 
```
postgres://postgres:your_password@localhost:5432/gator
```

From the root directory of the Blogegator direcotry, `cd` into the `sql/schema` directory and run:
```
goose postgres <connection_string> up
```

### Configuration

In your home directory, create the file `.gatorconfig.json`. In the file, create a JSON object with the single field `db_url`. It's value will be the database connection string from above with an additional `sslmode=disable` query string.
```
protocol://username:password@host:port/database?sslmode=disable
```

The complete file should look like the following:
```
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
}
```

## Usage

The Blogegator has the following commands:

| Syntax        | Description |
| :---          | :---        |
|`register <username>`| Register a new user.|
|`login <username>`   | Switch the current user. User must be registered.|
|`reset`              | Reset the list of registered users.|
|`users`              | Show a list of registered users.|
|`addfeed <name> <url>`      | Add a feed to be aggregated. Add to user list of followed feeds|
|`feeds`              | Show a list of feeds being aggregated.|
|`follow <url>`       | Follow the specified feed. Must be added first.|
|`unfollow <url>`     | Remove feed from following list for user.|
|`following`          | Show a list of feeds being followed by the current user.|
|`browse [limit]`     | Show the most \[limit]recent posts. If not specified \[limit] is 2|
|`agg <time_between_reqs>`    | Begin aggregating feeds, checking the next feed every <time_between_reqs>. This command should be started and left running on a separate terminal|

