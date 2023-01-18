package main

import (
	"database/sql"
	"final-project-sanbercode-go-batch-41/controllers"
	"final-project-sanbercode-go-batch-41/database"
	"final-project-sanbercode-go-batch-41/repository"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

var secret = []byte("secret")

const userkey = "user"

type UserAttempt struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	// ENV CONFIGURATION
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Print("Failed load file environment")
	} else {
		fmt.Println("Success load file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))

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
	router.Use(sessions.Sessions("mysession", sessions.NewCookieStore(secret)))

	// AUTH
	authorized := router.Group("/")

	admin := router.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))

	// ROUTER REGISTER AND LOGIN
	router.POST("/login", login)
	router.POST("/logout", logout)
	router.POST("/users", controllers.InsertUser)

	// ROUTER CATEGORY
	authorized.Use(AuthRequired).GET("/categories", controllers.GetAllCategory)
	authorized.Use(AuthRequired).GET("/categories/:id/products", controllers.GetProductByCategoryID)
	admin.POST("/categories", controllers.InsertCategory)
	admin.PUT("/categories/:id", controllers.UpdateCategory)
	admin.DELETE("/categories/:id", controllers.DeleteCategory)

	// ROUTER USER
	admin.Use(AuthRequired).GET("/users", controllers.GetAllUsers)
	authorized.Use(AuthRequired).GET("/users/:id/carts", controllers.GetCartByUserId)
	authorized.Use(AuthRequired).GET("/users/:id/orders", controllers.GetOrderByUserId)
	authorized.Use(AuthRequired).PUT("/users/:id", controllers.UpdateUser)
	authorized.Use(AuthRequired).DELETE("/users/:id", controllers.DeleteUser)

	// ROUTER PRODUCT
	authorized.Use(AuthRequired).GET("/products", controllers.GetAllProducts)
	admin.POST("/products", controllers.InsertProduct)
	admin.PUT("/products/:id", controllers.UpdateProduct)
	admin.DELETE("/products/:id", controllers.DeleteProduct)

	// ROUTER CART
	authorized.Use(AuthRequired).GET("/carts", controllers.GetAllCart)
	authorized.Use(AuthRequired).POST("/carts", controllers.InsertCart)
	authorized.Use(AuthRequired).PUT("/carts/:id", controllers.UpdateCart)
	authorized.Use(AuthRequired).DELETE("/carts/:id", controllers.DeleteCart)

	// ROUTER ORDER
	authorized.Use(AuthRequired).GET("/orders", controllers.GetAllOrder)
	authorized.Use(AuthRequired).POST("/orders", controllers.InsertOrder)
	admin.PUT("/orders/:id", controllers.UpdateOrder)
	admin.DELETE("/orders/:id", controllers.DeleteOrder)

	router.Run(":" + os.Getenv("PORT"))
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, please log in first."})
		return
	}
	c.Next()
}

func login(c *gin.Context) {
	session := sessions.Default(c)
	attempt := UserAttempt{}
	err = c.BindJSON(&attempt)
	if err != nil {
		panic(err)
	}
	username := attempt.Username
	password := attempt.Password

	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	_, users := repository.GetAllUsers(database.DbConnection)
	var valid bool = false

	for _, user := range users {
		if user.Username == username && user.Password == password {
			valid = true
			break
		}
	}

	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	session.Set(userkey, username)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(userkey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
