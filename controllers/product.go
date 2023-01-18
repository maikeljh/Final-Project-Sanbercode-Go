package controllers

import (
	"database/sql"
	"final-project-sanbercode-go-batch-41/database"
	"final-project-sanbercode-go-batch-41/repository"
	"final-project-sanbercode-go-batch-41/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllProducts(c *gin.Context) {
	var (
		result gin.H
	)

	products, err := repository.GetAllProducts(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": products,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertProduct(c *gin.Context) {
	var product structs.Product

	err := c.ShouldBindJSON(&product)
	if err != nil {
		panic(err)
	}

	err = repository.InsertProduct(database.DbConnection, product)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Create Product",
	})
}

func UpdateProduct(c *gin.Context) {
	var (
		result gin.H
	)
	var product structs.Product
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&product)
	if err != nil {
		panic(err)
	}

	product.ID = int64(id)

	err = repository.UpdateProduct(database.DbConnection, product)

	if err == sql.ErrNoRows {
		result = gin.H{
			"result": "Product dengan id tersebut tidak ditemukan.",
		}
	} else {
		result = gin.H{
			"result": "Success Update Product",
		}
	}

	c.JSON(http.StatusOK, result)
}

func DeleteProduct(c *gin.Context) {
	var (
		result gin.H
	)
	var product structs.Product
	id, err := strconv.Atoi(c.Param("id"))

	product.ID = int64(id)
	err = repository.DeleteProduct(database.DbConnection, product)

	if err == sql.ErrNoRows {
		result = gin.H{
			"result": "Product dengan id tersebut tidak ditemukan.",
		}
	} else {
		result = gin.H{
			"result": "Success Delete Product",
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetProductByCategoryID(c *gin.Context) {
	var (
		result gin.H
	)
	id, _ := strconv.Atoi(c.Param("id"))

	err1, products := repository.GetProductByCategoryID(database.DbConnection, id)

	if err1 != nil {
		result = gin.H{
			"result": err1,
		}
	} else {
		result = gin.H{
			"result": products,
		}
	}

	c.JSON(http.StatusOK, result)
}
