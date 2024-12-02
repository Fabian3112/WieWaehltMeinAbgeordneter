package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init_db() {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "Abgeordnete",
		AllowNativePasswords: true,
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

}

func Users() []User {
	var users []User

	rows, err := db.Query("SELECT * FROM User")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			panic(err)
		}

		users = append(users, user)
	}
	return users
}

func UsersByName(name string) []User {
	var users []User

	rows, err := db.Query("SELECT * FROM User WHERE name = ?", name)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			panic(err)
		}

		users = append(users, user)
	}
	return users
}

func UserById(id int64) User {
	var user User
	rows, err := db.Query("SELECT * FROM User WHERE id = ?", id)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	rows.Next()
	err = rows.Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		panic(err)
	}

	return user
}

func PutUser(user User_put) {

	rows, err := db.Query("INSERT INTO User (name,email) VALUES (?,?)", user.Name, user.Email)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

}
