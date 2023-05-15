package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func (server *Server) hello(c *gin.Context) {
	msg := HelloResponse{
		Message: "Hello World",
	}

	c.JSON(http.StatusOK, msg)
}
