package models

import (
	"database/sql"
)

//Struct modeling a recipe with its fields
type Recipe struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Ingredients string `json:"ingredients"`
}

//Struct with array of recipes
type RecipeCollection struct {
	Recipes []Recipe `json:"items"`
}

// GetRecipes from the DB, make the query to get the recipes
func GetRecipes(db *sql.DB) RecipeCollection {
	sql := "SELECT * FROM Recipes"
	//Executes the SQL Query
	rows, err := db.Query(sql)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()
	//initialize the struct to send in JSON
	result := RecipeCollection{}
	//Iterate all rows from the SQL Query
	for rows.Next() {
		//Initiliaze a new recipe
		recipe := Recipe{}
		//put the data inside the recipe
		err2 := rows.Scan(&recipe.ID, &recipe.Name, &recipe.Description, &recipe.Ingredients)
		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}
		//we add the recipe to the array from the struct result
		result.Recipes = append(result.Recipes, recipe)
	}
	return result
}

// PutRecipe into DB, make the query to insert the recipe
func PutRecipe(db *sql.DB, name string, description string, ingredients string) (int64, error) {
	sql := "INSERT INTO Recipes(name, description, ingredients) VALUES(?,?,?)"
	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name, description and ingredients'
	//Execute the query
	result, err2 := stmt.Exec(name, description, ingredients)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}
	return result.LastInsertId()
}

// UpdateRecipe into DB, make the query to update a recipe
func UpdateRecipe(db *sql.DB, id int, name string, description string, ingredients string) (int64, error) {
	sql := "UPDATE Recipes SET name = ?, description = ?, ingredients = ? WHERE id = ?"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name, description and ingredients'
	//Execute the query
	result, err2 := stmt.Exec(name, description, ingredients, id)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}
	return result.RowsAffected()

}

// DeleteRecipe into DB, make the query to delete a recipe
func DeleteRecipe(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM Recipes WHERE id = ?"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}

	// Replace the '?' in our prepared statement with 'id'
	// Execute the Instruction
	result, err2 := stmt.Exec(id)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
