package main

import (
	"database/sql"
	"recetas/handlers"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

//Create the connection to db.
func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

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
	CREATE TABLE IF NOT EXISTS Recipes(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    ingredients VARCHAR NOT NULL
	);
	`
	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}

func main() {

	db := initDB("storage.db")
	migrate(db)

	//Use echo for the server administration
	e := echo.New()

	e.File("/", "public/index.html")                    //The file to load
	e.GET("/recipes", handlers.GetRecipes(db))          //Load the recipes
	e.PUT("/recipes", handlers.PutRecipe(db))           //create the recipes
	e.POST("/recipes/:id", handlers.UpdateRecipe(db))   //update the recipes
	e.DELETE("/recipes/:id", handlers.DeleteRecipe(db)) //delete the recipes

	e.Start(":8080") //Run the server on port 8080.

}
