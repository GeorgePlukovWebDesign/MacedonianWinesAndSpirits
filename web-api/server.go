package main

import (
    "encoding/json"
    "fmt"
    "strconv"
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
        jsonResponse.AddProduct(product)

    	}

        json.NewEncoder(w).Encode(jsonResponse)


    }
}

func getSingleProduct (db *sql.DB) httprouter.Handle{
  return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
      w.Header().Set("Content-Type", "application/json")

      // selectStmt, err := db.Prepare()
      rows, err := db.Query("SELECT * from Products where id=?",ps.ByName("id"))
      if err != nil {
        fmt.Println("Error while selecting: %s", err)
      }
      // defer rows.close()
      // jsonResponse := ProductsResponse{Products:[]Product{} }
      var product Product

      for rows.Next() {
    		err = rows.Scan(&product.Id, &product.Name, &product.Description,&product.Type, &product.Price, &product.AlcoholPercentage, &product.Availability,&product.Year, &product.BottlesPerCase, &product.CostPerBottle, &product.Date)



    		if err != nil {
    			log.Fatal(err)
    		}
        json.NewEncoder(w).Encode(product)

    	}



    }
}

// Adds a product to the database and then redirects to the new items URL /products/:id
func postProducts (db *sql.DB) httprouter.Handle{
  return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

    	name := r.PostFormValue("name")
    	description := r.PostFormValue("description")
      typ := r.PostFormValue("typ")
      price := r.PostFormValue("price")
      alcoholPercentage := r.PostFormValue("alcoholPercentage")
      availability := r.PostFormValue("availability")
      year := r.PostFormValue("year")
      bottlesPerCase := r.PostFormValue("bottlesPerCase")
      costPerBottle := r.PostFormValue("costPerBottle")


      _, err := db.Exec("INSERT INTO Products (name, description, type, price,alcoholPercentage,availability,year,bottlesPerCase, costPerBottle) VALUES (?, ?, ?, ?, ?, ?, ?, ?,?);", name,description,typ,price,alcoholPercentage,availability,year,bottlesPerCase,costPerBottle)

      if err != nil {
        log.Fatal(err)
      }

      // Get the new product with the highest ID
      rows, err := db.Query("SELECT id from Products where id=(select max(id) from Products);")
      if err != nil {
        log.Fatal(err)
      }

      var product Product

      // Redirect to the product id page
      for rows.Next() {
    		err = rows.Scan(&product.Id)
    		if err != nil {
    			log.Fatal(err)
    		}
        http.Redirect(w, r, "/products/" + strconv.Itoa(product.Id), 300)
    	}

    }
}
// TODO: Send back some sort of response
func deleteSingleProduct (db *sql.DB) httprouter.Handle{
  return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

      // selectStmt, err := db.Prepare()
      _, err := db.Exec("DELETE FROM Products where id=?",ps.ByName("id"))
      if err != nil {
        log.Fatal(err)
      }
      // w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
      //x w.WriteHeader(http.StatusOK)

    }
}


func main() {
  db := initDB("MacedonianWinesAndSpirits.db")
  migrate(db)

  router := httprouter.New()


  // Get products
  router.GET("/products/", getProducts(db))
  router.GET("/products/:id", getSingleProduct(db))

  // Post products
  router.POST("/products/", postProducts(db))

  // Delete product
  router.DELETE("/products/:id", deleteSingleProduct(db))


  log.Fatal(http.ListenAndServe(":8080", router))

}
