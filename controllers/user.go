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

func GetAllUsers(c *gin.Context) {
	var (
		result gin.H
	)

	users, err := repository.GetAllUsers(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": users,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertUser(c *gin.Context) {
	var user structs.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}

	err = repository.InsertUser(database.DbConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Create User",
	})
}

func UpdateUser(c *gin.Context) {
	var (
		result gin.H
	)
	var user structs.User
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}

	user.ID = int64(id)

	err = repository.UpdateUser(database.DbConnection, user)

	if err == sql.ErrNoRows {
		result = gin.H{
			"result": "User dengan id tersebut tidak ditemukan.",
		}
	} else {
		result = gin.H{
			"result": "Success Update User",
		}
	}

	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	var (
		result gin.H
	)
	var user structs.User
	id, err := strconv.Atoi(c.Param("id"))

	user.ID = int64(id)
	err = repository.DeleteUser(database.DbConnection, user)

	if err == sql.ErrNoRows {
		result = gin.H{
			"result": "User dengan id tersebut tidak ditemukan.",
		}
	} else {
		result = gin.H{
			"result": "Success Delete User",
		}
	}

	c.JSON(http.StatusOK, result)
}
