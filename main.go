package main

import (
	"database/sql"
	"fmt"
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

		err = rs.QueryRows(db, "select table_schema as table_schema, table_name as table_name from information_schema.tables")
		if err != nil {
			panic(err)
		}

		for _, row := range rs.Rows {
			err := rs.QueryRows(db, fmt.Sprintf("select * from %s.%s limit 10", row["table_schema"], row["table_name"]))
			if err != nil {
				panic(err)
			}
			rs.Print()
		}
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

		err = rs.QueryRows(db, "select table_schema as table_schema, table_name as table_name from information_schema.tables")
		if err != nil {
			panic(err)
		}

		for _, row := range rs.Rows {
			err := rs.QueryRows(db, fmt.Sprintf("select * from %s.%s limit 10", row["table_schema"], row["table_name"]))
			if err != nil {
				panic(err)
			}
			rs.Print()
		}
	}
}
