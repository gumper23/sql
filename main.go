package main

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gumper23/sql/rs"
	_ "github.com/lib/pq"
)

func main() {
	var rs rs.Resultset

	dsn, ok := os.LookupEnv("POSTGRES_DSN")
	if ok {
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		err = rs.QueryRows(db, "select * from pg_stat_activity")
		if err != nil {
			panic(err)
		}
		rs.Print()

		err = rs.QueryRows(db, "select * from ints")
		if err != nil {
			panic(err)
		}
		rs.Print()
	}

	dsn, ok = os.LookupEnv("MYSQL_DSN")
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

		err = rs.QueryRows(db, "select * from information_schema.processlist")
		if err != nil {
			panic(err)
		}
		rs.Print()

		err = rs.QueryRows(db, "select * from qm.ints")
		if err != nil {
			panic(err)
		}
		rs.Print()

		err = rs.QueryRows(db, "select * from qm.dates")
		if err != nil {
			panic(err)
		}
		rs.Print()
	}
}
