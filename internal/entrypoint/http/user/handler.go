// internal/entrypoint/http/user/handler.go
package userhandler

import (
	"github.com/gin-gonic/gin"
	"lawise-go/internal/entrypoint/http"
)

// UserService interface defines the contract for user business logic
type UserService interface {
	// Add your service methods here
	// CreateUser(user *User) error
	// GetUser(id string) (*User, error)
	// etc...
}

type UserHandler struct {
	userService UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

// RegisterRoutes implements the RouteHandler interface
func (h *UserHandler) RegisterRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.POST("", h.CreateUser)
		users.GET("", h.GetUsers)
		users.GET("/:id", h.GetUser)
		users.PUT("/:id", h.UpdateUser)
		users.DELETE("/:id", h.DeleteUser)
	}
}

// Handler methods
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req http.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Call service method
	// user, err := h.userService.CreateUser(&req)
	// if err != nil {
	//     c.JSON(500, gin.H{"error": err.Error()})
	//     return
	// }

	c.JSON(201, gin.H{"message": "User created successfully"})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	// Call service method
	// user, err := h.userService.GetUser(id)
	// if err != nil {
	//     c.JSON(500, gin.H{"error": err.Error()})
	//     return
	// }

	c.JSON(200, gin.H{
		"id":      id,
		"message": "Get user endpoint",
	})
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	// Implementation
	c.JSON(200, gin.H{"message": "Get users endpoint"})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	// Implementation
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Update user endpoint", "id": id})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	// Implementation
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Delete user endpoint", "id": id})
}
