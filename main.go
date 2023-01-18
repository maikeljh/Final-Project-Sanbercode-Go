package main

import (
	"database/sql"
	"final-project-sanbercode-go-batch-41/controllers"
	"final-project-sanbercode-go-batch-41/database"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	// ENV CONFIGURATION
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Print("Failed load file environment")
	} else {
		fmt.Println("Success load file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	// ROUTER GIN
	router := gin.Default()

	// AUTH
	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))

	// ROUTER CATEGORY
	router.GET("/categories", controllers.GetAllCategory)
	authorized.POST("/categories", controllers.InsertCategory)
	authorized.PUT("/categories/:id", controllers.UpdateCategory)
	authorized.DELETE("/categories/:id", controllers.DeleteCategory)

	// ROUTER USER
	router.GET("/users", controllers.GetAllUsers)
	authorized.POST("/users", controllers.InsertUser)
	authorized.PUT("/users/:id", controllers.UpdateUser)
	authorized.DELETE("/users/:id", controllers.DeleteUser)

	// ROUTER PRODUCT
	router.GET("/products", controllers.GetAllProducts)
	authorized.POST("/products", controllers.InsertProduct)
	authorized.PUT("/products/:id", controllers.UpdateProduct)
	authorized.DELETE("/products/:id", controllers.DeleteProduct)
	router.Run("localhost:8080")
}
