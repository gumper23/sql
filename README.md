# sql.rs
Golang DB helper that allows exploration of arbitrary queries. Results are converted to strings.

# Example:

```go
package main

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gumper23/sql/rs"
)

func main() {
	var rs rs.Resultset
	dsn, ok := os.LookupEnv("MYSQL_DSN")
	if ok {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		err = rs.QueryRows(db, "select * from information_schema.tables")
		if err != nil {
			panic(err)
		}
		rs.Hprint()
	}
}

 ```

# Testing (optional)
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
	
Tested with Postgres and MySQL DBMSes. 

