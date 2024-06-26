package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"myappswagger/models"
)

// @title           User Management API
// @version         1.0
// @description     This is a sample server for managing users.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

var users = []models.User{
	{ID: 1, Name: "John Doe", Email: "john.doe@example.com"},
	{ID: 2, Name: "Jane Smith", Email: "jane.smith@example.com"},
}

// getUsers godoc
// @Summary      List users
// @Description  get users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.User
// @Router       /users [get]
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// getUserByID godoc
// @Summary      Get a user by ID
// @Description  get user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int     true  "User ID"
// @Success      200  {object}  models.User
// @Failure      404  
// @Router       /users/{id} [get]
func getUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

// createUser godoc
// @Summary      Create a user
// @Description  create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "New User"
// @Success      201  {object}  models.User
// @Failure      400 
// @Router       /users [post]
func createUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser.ID = len(users) + 1
	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

// updateUser godoc
// @Summary      Update a user
// @Description  update user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "User ID"
// @Param        user  body      models.User  true  "Updated User"
// @Success      200  {object}  models.User
// @Failure      400 
// @Failure      404 
// @Router       /users/{id} [put]
func updateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, user := range users {
		if user.ID == id {
			users[i] = updatedUser
			users[i].ID = id
			c.JSON(http.StatusOK, users[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

// deleteUser godoc
// @Summary      Delete a user
// @Description  delete user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int     true  "User ID"
// @Success      204 
// @Failure      404 
// @Router       /users/{id} [delete]
func deleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusNoContent, gin.H{})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func main() {
	r := gin.Default()

	// Route vers la documentation Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	{
		api.GET("/users", getUsers)
		api.GET("/users/:id", getUserByID)
		api.POST("/users", createUser)
		api.PUT("/users/:id", updateUser)
		api.DELETE("/users/:id", deleteUser)
	}

	r.Run()
}
