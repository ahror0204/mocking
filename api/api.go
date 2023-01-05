package api

import (
	"net/http"
	"strconv"

	"github.com/ahror0204/mocking/storage"

	"github.com/gin-gonic/gin"
)

func NewServer(strg storage.StorageI) Server {
	router := gin.Default()

	s := Server{
		storage: strg,
	}

	router.POST("/users", s.CreateUser)
	router.GET("/users/:id", s.GetUser)

	s.Router = router
	return s
}

type Server struct {
	storage storage.StorageI
	Router  *gin.Engine
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type CreateUserRequest struct {
	FirstName   string `json:"first_name" binding:"required,min=2,max=30"`
	LastName    string `json:"last_name" binding:"required,min=2,max=30"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email" binding:"required,email"`
}

func errorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Error: err.Error(),
	}
}

func (s *Server) CreateUser(c *gin.Context) {
	var (
		req CreateUserRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := s.storage.CreateUser(&storage.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *Server) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.storage.GetUser(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}
