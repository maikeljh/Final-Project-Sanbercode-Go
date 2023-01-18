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

func GetAllOrder(c *gin.Context) {
	var (
		result gin.H
	)

	order, err := repository.GetAllOrder(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": order,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertOrder(c *gin.Context) {
	var order structs.Order

	err := c.ShouldBindJSON(&order)
	if err != nil {
		panic(err)
	}

	err = repository.InsertOrder(database.DbConnection, order)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Create Order",
	})
}

func UpdateOrder(c *gin.Context) {
	var (
		result gin.H
	)
	var order structs.Order
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&order)
	if err != nil {
		panic(err)
	}

	order.ID = int64(id)

	err = repository.UpdateOrder(database.DbConnection, order)

	if err == sql.ErrNoRows {
		result = gin.H{
			"result": "Order dengan id tersebut tidak ditemukan.",
		}
	} else {
		result = gin.H{
			"result": "Success Update Order",
		}
	}

	c.JSON(http.StatusOK, result)
}

func DeleteOrder(c *gin.Context) {
	var (
		result gin.H
	)
	var order structs.Order
	id, err := strconv.Atoi(c.Param("id"))

	order.ID = int64(id)
	err = repository.DeleteOrder(database.DbConnection, order)

	if err == sql.ErrNoRows {
		result = gin.H{
			"result": "Order dengan id tersebut tidak ditemukan.",
		}
	} else {
		result = gin.H{
			"result": "Success Delete Order",
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetOrderByUserId(c *gin.Context) {
	var (
		result gin.H
	)

	id, _ := strconv.Atoi(c.Param("id"))

	err1, order := repository.GetOrderByUserId(database.DbConnection, int64(id))

	if err1 != nil {
		result = gin.H{
			"result": err1,
		}
	} else {
		result = gin.H{
			"result": order,
		}
	}

	c.JSON(http.StatusOK, result)
}
