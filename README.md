# sql.rs
Golang DB Interface For Queries - Returns a slice of map[string][string]

# Usage (optional)
```bash
$ export MYSQL_DSN="<user>:<password>@tcp(<host>:<port>)/<dbname>
```

AND/OR:

```bash
$ export POSTGRES_DSN="user=<user> password=<password host=<host> port=<port> dbname=<dbname> sslmode=disable"
```

THEN:
```bash
$ go run main.go
```

This will print 10 rows from every table in the instance that <user> has access to.


