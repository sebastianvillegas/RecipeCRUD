package main

import (
	"database/sql"
	"recetas/handlers"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

//Create the connection to db.
func initDB(addr string) *sql.DB {
	db, err := sql.Open("postgres", addr)

	// Here we check for any db errors
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

//First SQL Instructions to create tables.
func migrate(db *sql.DB) {
	sql := `
	CREATE SEQUENCE IF NOT EXISTS customer_seq;

	CREATE TABLE IF NOT EXISTS Recipes(
		id INT PRIMARY KEY DEFAULT nextval('customer_seq'),
		name STRING NOT NULL,
    description STRING NOT NULL,
    ingredients STRING NOT NULL
	);
	`
	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}

func main() {
	addr := "postgresql://maxroach@localhost:26257/recipes?sslmode=disable"
	db := initDB(addr)
	migrate(db)

	//Use echo for the server administration
	e := echo.New()

	e.File("/", "public/index.html")                    //The file to load
	e.GET("/recipes", handlers.GetRecipes(db))          //Load the recipes
	e.PUT("/recipes", handlers.PutRecipe(db))           //create the recipes
	e.POST("/recipes/:id", handlers.UpdateRecipe(db))   //update the recipes
	e.DELETE("/recipes/:id", handlers.DeleteRecipe(db)) //delete the recipes

	e.Start(":8081") //Run the server on port 8080.

}
