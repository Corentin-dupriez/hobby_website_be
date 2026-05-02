package main

import (
	"database/sql"
	"fmt"
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
	return db
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
}

func QueryDB(db *sql.DB, q string) {
	rows, err := db.Query("SELECT * from users")
	if err != nil {
		log.Fatal(err)
	}
	var res []User
	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, u)
	}
	for _, u := range res {
		fmt.Println(u)
	}
}
