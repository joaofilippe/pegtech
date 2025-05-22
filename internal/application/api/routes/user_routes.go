package routes

import (
	"net/http"

	"github.com/joaofilippe/pegtech/internal/domain/iservices"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// UserRoutes handles all user-related routes
type UserRoutes struct {
	userService iservices.UserService
}

// NewUserRoutes creates a new instance of UserRoutes
func NewUserRoutes(userService iservices.UserService) *UserRoutes {
	return &UserRoutes{
		userService: userService,
	}
}

// Register registers all user routes
func (r *UserRoutes) Register(e *echo.Echo) {
	e.POST("/users", r.createUser)
	e.GET("/users/:id", r.getUserByID)
	e.GET("/users/email/:email", r.getUserByEmail)
	e.PUT("/users/:id", r.updateUser)
	e.DELETE("/users/:id", r.deleteUser)
	e.POST("/login", r.login)
}

// createUser handles user creation
func (r *UserRoutes) createUser(c echo.Context) error {
	var CreateUserInput struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&CreateUserInput); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	user, err := r.userService.CreateUser(CreateUserInput.Username, CreateUserInput.Email, CreateUserInput.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

// getUserByID handles user retrieval by ID
func (r *UserRoutes) getUserByID(c echo.Context) error {
	id := c.Param("id")

	user, err := r.userService.GetUserByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// getUserByEmail handles user retrieval by email
func (r *UserRoutes) getUserByEmail(c echo.Context) error {
	email := c.Param("email")

	user, err := r.userService.GetUserByEmail(email)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// updateUser handles user updates
func (r *UserRoutes) updateUser(c echo.Context) error {
	id := c.Param("id")

	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	user, err := r.userService.UpdateUser(id, input.Username, input.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// deleteUser handles user deletion
func (r *UserRoutes) deleteUser(c echo.Context) error {
	id := c.Param("id")

	if err := r.userService.DeleteUser(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

// login handles user authentication
func (r *UserRoutes) login(c echo.Context) error {
	var LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&LoginInput); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	user, err := r.userService.GetUserByEmail(LoginInput.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	// Compare the provided password with the stored hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(LoginInput.Password))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	// Remove sensitive information before sending response
	user.Password = ""

	return c.JSON(http.StatusOK, user)
}
