# Gator - RSS Feed Aggrigator

Gator is a simple RSS feed aggrigator written in Go.

## Prerequisites

Gator has 2 dependencies
- Postgres version 15+
- Go version 1.23+

Once you have these installed you can install gator using the go install command: ```go install github.com/anthony81799/gator@latest```

To use the application create a ```.gatorconfig.json``` file in your home directory. The cofiguration requires the database url stored like this:
```json
{
    "db_url":<database_connection_string>
}
```

Once the configuration is set up ypu can run gator with the ```gator``` command and any sub command such as:
- ```register <username>```
- ```login <username>```
- ```reset```
- ```users```
- ```agg <time_between_reqs>```
- ```addfeed <name> <url>```
- ```feeds```
- ```follow <url>```
- ```following```
- ```unfollow <url>```
- ```browse <optional_limit>```