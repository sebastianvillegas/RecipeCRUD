package handlers

import (
	"database/sql"
	"net/http"
	"recetas/models"
	"strconv"

	"github.com/labstack/echo"
)

type H map[string]interface{}

//It calls the respective model to bring the recipes from db and return a JSON response
func GetRecipes(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetRecipes(db))
	}
}

//It calls the respective model to create the recipe in db and return a JSON response
func PutRecipe(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new recipe
		var recipe models.Recipe
		// Map imcoming JSON body to the new Recipe
		c.Bind(&recipe)
		// Add a recipe using our new model
		id, err := models.PutRecipe(db, recipe.Name, recipe.Description, recipe.Ingredients)
		// Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
			// Handle any errors
		} else {
			return err
		}
	}
}

//It calls the respective model to update a recipe in db and return a JSON response
func UpdateRecipe(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Instantiate a new recipe
		var recipe models.Recipe
		// Map imcoming JSON body to the new Recipe
		c.Bind(&recipe)
		//Read the id from the request
		id, _ := strconv.Atoi(c.Param("id"))
		//update a recipe using the model
		_, err := models.UpdateRecipe(db, id, recipe.Name, recipe.Description, recipe.Ingredients)
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"updated": id,
				"name":    recipe.Name,
			})
			// Handle errors
		} else {
			return err
		}
	}
}

//It calls the respective model to delete a recipe in db and return a JSON response
func DeleteRecipe(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		////Read the id from the request
		id, _ := strconv.Atoi(c.Param("id"))
		// Use our new model to delete a recipe
		_, err := models.DeleteRecipe(db, id)
		// Return a JSON response on success
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
			// Handle errors
		} else {
			return err
		}
	}
}
