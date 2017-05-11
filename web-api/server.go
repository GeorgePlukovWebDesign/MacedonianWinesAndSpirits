package main

import (
    "fmt"
    "database/sql"
    // "go-echo-vue/handlers"
    "log"
    "github.com/julienschmidt/httprouter"
    _ "github.com/mattn/go-sqlite3"
    "net/http"
)

type Product struct {
	Id    int
	Name string
	Description  string
  Type string
  Price string
  AlcoholPercentage string
  availability string
  year string
  BottlesPerCase int
  CostPerBottle float32
  Date  string

}

func main() {
  db := initDB("MacedonianWinesAndSpirits.db")
  migrate(db)

  // // Create and execute a insert statement
  // insertStmt, err := conn.Prepare("INSERT INTO Products (name, description, type, price,alcoholPercentage,availability,year,bottlesPerCase, costPerBottle) VALUES ('Carberator salmon', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.', 'Red', '9.75', '11.00', 'available', '2017', 6, 9.75);")
  // err = selectStmt.Exec()
  // if err != nil {
  //   fmt.Println("Error while inserting: %s", err)
  // }

  router := httprouter.New()

    router.GET("/products", getProducts(db))
    router.POST("/products", postProducts(db))

    // router.NotFound = &DefaultHandler{}



    log.Fatal(http.ListenAndServe(":8080", router))

  }
