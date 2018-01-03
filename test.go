package main

import (
	"context"
	"database/sql"
	"time"

	"sync"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1)/test")
	if err != nil {
		panic(err)
	}

	go func() {
		<-time.After(100 * time.Millisecond)
		cancel()
	}()

	var w sync.WaitGroup
	for i := 0; i < 100; i++ {
		w.Add(1)
		go func(ci int) {

			defer w.Done()
			rows, err := db.QueryContext(ctx, "select email from authors")
			if err != nil {
				log.Println(err)
				return
			}

			defer rows.Close()

			for rows.Next() {
				var email string
				err = rows.Scan(&email)
				if err != nil {
					log.Println(err)
					return
				}
			}

			if rows.Err() != nil {
				log.Print(rows.Err())
			}
		}(i)

	}

	w.Wait()

}
