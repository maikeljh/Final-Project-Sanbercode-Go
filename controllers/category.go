package controllers

import (
	"database/sql"
	"final-project-sanbercode-go-batch-41/database"
	"final-project-sanbercode-go-batch-41/repository"
	"final-project-sanbercode-go-batch-41/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCategory(c *gin.Context) {
	var (
		result gin.H
	)

	categories, err := repository.GetAllCategory(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": categories,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCategory(c *gin.Context) {
	var category structs.Category

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	err = repository.InsertCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Category",
	})
}

func UpdateCategory(c *gin.Context) {
	var (
		result gin.H
	)
	var category structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	category.ID = id

	err = repository.UpdateCategory(database.DbConnection, category)

	if err == sql.ErrNoRows {
		result = gin.H{
			"result": "Category dengan id tersebut tidak ditemukan.",
		}
	} else {
		result = gin.H{
			"result": "Success Update Category",
		}
	}

	c.JSON(http.StatusOK, result)
}

func DeleteCategory(c *gin.Context) {
	var (
		result gin.H
	)
	var category structs.Category
	id, err := strconv.Atoi(c.Param("id"))

	category.ID = id
	err = repository.DeleteCategory(database.DbConnection, category)

	if err == sql.ErrNoRows {
		result = gin.H{
			"result": "Category dengan id tersebut tidak ditemukan.",
		}
	} else {
		result = gin.H{
			"result": "Success Delete Category",
		}
	}

	c.JSON(http.StatusOK, result)
}
