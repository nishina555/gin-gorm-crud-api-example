package main

import (
	"gin-gorn-crud-api-example/database"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id    int    `json:"id" param:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// func hello(c *gin.Context) {
// 	c.String(http.StatusOK, "Hello, World!")
// }

func getUsers(c *gin.Context) {
	users := []User{}
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	user := User{}
	if err := c.Bind(&user); err != nil {
		panic(err.Error())
	}
	database.DB.Take(&user)

	c.JSON(http.StatusOK, user)
}

func updateUser(c *gin.Context) {
	user := User{}
	if err := c.Bind(&user); err != nil {
		panic(err.Error())
	}
	database.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

func createUser(c *gin.Context) {
	user := User{}
	if err := c.Bind(&user); err != nil {
		panic(err.Error())
	}
	database.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&User{}, id)
	c.Status(http.StatusNoContent)
}

func main() {
	r := gin.Default()
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	// 疎通確認用
	// r.GET("/", hello)

	r.GET("/users", getUsers)
	r.GET("/users/:id", getUser)
	r.PUT("/users/:id", updateUser)
	r.POST("/users", createUser)
	r.DELETE("/users/:id", deleteUser)
	r.Run(":3000")
}
