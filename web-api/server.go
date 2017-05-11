package main

import (
    "encoding/json"
    "fmt"
    "database/sql"
    // "go-echo-vue/handlers"
    "log"
    "github.com/julienschmidt/httprouter"
    _ "github.com/mattn/go-sqlite3"
    "net/http"
)

// Responses
type ProductsResponse struct{
  Page   int      `json:"page"`
  Products []Product `json:"products"`
}
func (resp *ProductsResponse) AddProduct(prod Product) []Product{
    resp.Products = append(resp.Products, prod)
    return resp.Products
}


type Product struct {
	Id    int `db:"id"`
	Name string `db:"name"`
	Description  string `db:"description"`
  Type string `db:type`
  Price string `db:price`
  AlcoholPercentage string `db:alcoholPercentage`
  Availability string `db:availability`
  Year string `db:year`
  BottlesPerCase int `db:bottlesPerCase`
  CostPerBottle float32 `db:costPerBottle`
  Date  string `db:date`

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
    CREATE TABLE IF NOT EXISTS Products(id INTEGER PRIMARY KEY AUTOINCREMENT,name VARCHAR(100),description TEXT default 'none',type TEXT default 'N/A',price TEXT default '0.00',alcoholPercentage TEXT default '00.00', availability TEXT default 'available', year TEXT default '2017', bottlesPerCase INTEGER default 6, costPerBottle REAL deault '10.0',date TEXT default CURRENT_TIMESTAMP);
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


      // selectStmt, err := db.Prepare()
      rows, err := db.Query("SELECT * from Products")
      if err != nil {
        fmt.Println("Error while selecting: %s", err)
      }
      // defer rows.close()
      jsonResponse := ProductsResponse{Products:[]Product{} }
      var product Product

      for rows.Next() {
    		err = rows.Scan(&product.Id, &product.Name, &product.Description,&product.Type, &product.Price, &product.AlcoholPercentage, &product.Availability,&product.Year, &product.BottlesPerCase, &product.CostPerBottle, &product.Date)

    		if err != nil {
    			log.Fatal(err)
    		}
    		fmt.Println(product)
        jsonResponse.AddProduct(product)

    	}

        json.NewEncoder(w).Encode(jsonResponse)


    }
}

func postProducts (db *sql.DB) httprouter.Handle{
  return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
      w.Header().Set("Content-Type", "application/json")

      insertStmt, err := db.Prepare("INSERT INTO Products (id,name, description, type, price,alcoholPercentage,availability,year,bottlesPerCase, costPerBottle) VALUES (null,'Carberator salmon', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.', 'Red', '9.75', '11.00', 'available', '2017', 6, 9.75);")
      // insertStmt, err := db.Prepare("SELECT * from Products")
      _, err = insertStmt.Exec()
      if err != nil {
        fmt.Println("Error while selecting: %s", err)
      }

      // fmt.Fprint(w, "nigger2")
    }
}



func main() {
  db := initDB("MacedonianWinesAndSpirits.db")
  migrate(db)

  router := httprouter.New()

  router.GET("/products", getProducts(db))
  router.GET("/products/:id", getProducts(db))

  router.POST("/products", postProducts(db))


  log.Fatal(http.ListenAndServe(":8080", router))

}
