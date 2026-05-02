package main

import (
	"database/sql"
	"log"
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

func ConnectToDB() *sql.DB {
	envFile, _ := godotenv.Read(".env")
	envHost := envFile["DB_HOST"]
	envPort := envFile["DB_PORT"]
	envUser := envFile["DB_USER"]
	envPassword := envFile["DB_PASSWORD"]

	dsn := "postgres://" + envUser + ":" + envPassword + "@" + envHost + ":" + envPort + "/postgres"

	// postgres[ql]://[user[:pwd]@][net-location][:port][/dbname][?param1=value1&...]

	c, err := pq.NewConnector(dsn)
	if err != nil {
		log.Fatal(err)
	}
	db := sql.OpenDB(c)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("Connected to database", "DB_HOST", envHost)

	res := structureExists(db)

	if !res {
		slog.Info("Table structure doesn't exist")
	}

	return db
}

func structureExists(d *sql.DB) bool {
	r := QueryDB(d, "SELECT tablename from pg_catalog.pg_tables")

	for r.Next() {
		var table string
		err := r.Scan(&table)
		if err != nil {
			log.Fatal(err)
		}
		if table == "recipes" {
			return true
		}
	}
	return false
}

func createDBStructure(d *sql.DB) {
}

func QueryDB(db *sql.DB, q string) *sql.Rows {
	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	return rows
	// var res []User
	// for rows.Next() {
	// 	var u User
	// 	err = rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	res = append(res, u)
	// }
	// for _, u := range res {
	// 	fmt.Println(u)
	// }
}
