package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/gumper23/sql/qm"
)

func main() {
	fmt.Printf("Howdy\n")

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

		var cols []string
		rows, cols, err := qm.QueryRows(db, "select * from pg_stat_activity")
		if err != nil {
			panic(err)
		}
		printQueryMap(rows, cols)

		rows, cols, err = qm.QueryRows(db, "select * from ints")
		if err != nil {
			panic(err)
		}
		printQueryMap(rows, cols)
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

		rows, cols, err := qm.QueryRows(db, "select * from information_schema.processlist")
		if err != nil {
			panic(err)
		}
		printQueryMap(rows, cols)

		rows, cols, err = qm.QueryRows(db, "select * from qm.ints")
		if err != nil {
			panic(err)
		}
		printQueryMap(rows, cols)

		rows, cols, err = qm.QueryRows(db, "select * from qm.dates")
		if err != nil {
			panic(err)
		}
		printQueryMap(rows, cols)
	}
}

func printQueryMap(qm []map[string]string, cols []string) {
	for _, row := range qm {
		for _, col := range cols {
			fmt.Printf("%-20s\t%s\n", col+":", row[col])
		}
		fmt.Println()
	}
}
