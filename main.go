package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type Series struct {
	SeriesId   string `json:"series_id"`
	SeriesName string `json:"series_name"`
}

func main() {
	e := echo.New()

	e.GET("/series", func(c echo.Context) error {
		db, err := sql.Open("mysql", "hblake:password@tcp(127.0.0.1:3306)/books")

		if err != nil {
			panic(err.Error())
		}

		defer db.Close()

		var series Series

		results, err := db.Query("SELECT * FROM series")
		if err != nil {
			fmt.Println("This is the actual error HEATHER: ", err.Error())
			fmt.Println("Press cmd-f and type `main.go:` to highlight the lines of code")
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		for results.Next() {
			var s Series
			// for each row, scan the result into our tag composite object
			err = results.Scan(&s.SeriesId, &s.SeriesName)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			log.Printf(s.SeriesId)
		}

		return c.JSON(http.StatusOK, series)
	})

	port := os.Getenv("BOOKS_PORT")
	portString := ":" + port
	e.Logger.Fatal(e.Start(portString))
}
