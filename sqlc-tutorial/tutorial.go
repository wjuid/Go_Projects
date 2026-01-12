package main

import (
	"context"
	"database/sql"
	"log"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"tutorial.sqlc.dev/app/tutorial"
)

func run() error {
	// Application logic goes here
	ctx := context.Background()
	db, err := sql.Open("mysql", "root:qnmd7456@tcp(192.168.0.105:3306)/wju_mysql?parseTime=true")
	if err != nil {
		return err
	}
	queries := tutorial.New(db)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	result, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{
		Name: "wju_mysql_remote project",
		Bio:  sql.NullString{String: "this section  mysql remote projects", Valid: true},
	})
	if err != nil {
		return err
	}

	insertedAuthorID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	log.Println(insertedAuthorID)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthorID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthorID, fetchedAuthor.ID))
	return nil
}
func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
