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

func GetAllCart(c *gin.Context) {
	var (
		result gin.H
	)

	cart, err := repository.GetAllCart(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": cart,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCart(c *gin.Context) {
	var cart structs.Cart

	err := c.ShouldBindJSON(&cart)
	if err != nil {
		panic(err)
	}

	err = repository.InsertCart(database.DbConnection, cart)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Create Cart",
	})
}

func UpdateCart(c *gin.Context) {
	var (
		result gin.H
	)
	var cart structs.Cart
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&cart)
	if err != nil {
		panic(err)
	}

	cart.ID = int64(id)

	err = repository.UpdateCart(database.DbConnection, cart)

	if err == sql.ErrNoRows {
		result = gin.H{
			"result": "Cart dengan id tersebut tidak ditemukan.",
		}
	} else {
		result = gin.H{
			"result": "Success Update Cart",
		}
	}

	c.JSON(http.StatusOK, result)
}

func DeleteCart(c *gin.Context) {
	var (
		result gin.H
	)
	var cart structs.Cart
	id, err := strconv.Atoi(c.Param("id"))

	cart.ID = int64(id)
	err = repository.DeleteCart(database.DbConnection, cart)

	if err == sql.ErrNoRows {
		result = gin.H{
			"result": "Cart dengan id tersebut tidak ditemukan.",
		}
	} else {
		result = gin.H{
			"result": "Success Delete Cart",
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetCartByUserId(c *gin.Context) {
	var (
		result gin.H
	)

	id, _ := strconv.Atoi(c.Param("id"))

	err1, cart := repository.GetCartByUserId(database.DbConnection, int64(id))

	if err1 != nil {
		result = gin.H{
			"result": err1,
		}
	} else {
		result = gin.H{
			"result": cart,
		}
	}

	c.JSON(http.StatusOK, result)
}
