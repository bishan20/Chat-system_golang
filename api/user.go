package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Name string `json:"name" binding:"required"`
}

// type userResponse struct {
// 	Id   int32  `json:"id"`
// 	Name string `json:"name"`
// }

// func newUserResponse(user db.User) userResponse {
// 	return userResponse{
// 		Id:   user.ID,
// 		Name: user.Name,
// 	}
// }

func (server *Server) createUser(ctx *gin.Context) {

	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	name := req.Name

	user, err := server.store.CreateUser(ctx, name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// response := newUserResponse(user)
	ctx.JSON(http.StatusOK, user)
}
