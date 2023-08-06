package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title     string
	Completed bool
}

var db *gorm.DB

func getAllTodos(c *gin.Context) {
	var todos []Todo
	db.Find(&todos)
	c.JSON(200, todos)
}

func createTodo(c *gin.Context) {
	var todos []Todo
	if err := c.ShouldBindJSON(&todos); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	for _, todo := range todos {
		db.Create(&todo)
	}
	c.JSON(201, todos)
}

func main() {
	dsn := "user=admin password=Myoff2*** dbname=a sslmode=disable"
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = dbConn
	db.AutoMigrate(&Todo{})

	r := gin.Default()
	r.GET("/todos", getAllTodos)
	r.POST("/todos", createTodo)

	r.Run(":8080")
}
