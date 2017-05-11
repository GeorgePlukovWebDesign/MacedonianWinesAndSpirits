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

func initDB(filepath string) *sql.DB {
    db, err := sql.Open("sqlite3", filepath)

    // Here we check for any db errors then exit
    if err != nil {
        panic(err)
    }

    // If we don't get any errors but somehow still don't get a db connection
    // we exit as well
    if db == nil {
        panic("db nil")
    }
    return db
}

func migrate(db *sql.DB) {
    sql := `
    CREATE TABLE IF NOT EXISTS Products(id INTEGER PRIMARY KEY AUTOINCREMENT,name VARCHAR(100),description TEXT default 'none',type TEXT default 'N/A',price TEXT default '0.00',alcoholPercentage TEXT default '00.00', availability TEXT default 'available', year TEXT default '2017', bottlesPerCase INTEGER default 6, costPerBottle REAL deault '10.0',date TEXT );
    `
    _, err := db.Exec(sql)
    // Exit if something goes wrong with our SQL statement above
    if err != nil {
        panic(err)
    }
}


func getProducts (db *sql.DB) httprouter.Handle{
  return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
      w.Header().Set("Content-Type", "application/json")

      selectStmt, err := db.Prepare("SELECT * from Products")
      _, err = selectStmt.Exec()
      if err != nil {
        fmt.Println("Error while selecting: %s", err)
      }

      fmt.Fprint(w, "nigger")
    }
}

func postProducts (db *sql.DB) httprouter.Handle{
  return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
      w.Header().Set("Content-Type", "application/json")


      selectStmt, err := db.Prepare("SELECT * from Products")
      _, err = selectStmt.Exec()
      if err != nil {
        fmt.Println("Error while selecting: %s", err)
      }

      fmt.Fprint(w, "nigger2")
    }
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
