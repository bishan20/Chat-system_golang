package api

import (
	db "Chat-system_golang/db/sqlc"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type storeMessageRequest struct {
	Message    string `json:"message" binding:"required"`
	SenderId   int32  `json:"sender_id" binding:"required"`
	ReceiverId int32  `json:"receiver_id" binding:"required"`
}

func (server *Server) storeMessage(ctx *gin.Context) {

	var req storeMessageRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.StoreMessageParams{
		Message:    req.Message,
		SenderID:   req.SenderId,
		ReceiverID: req.ReceiverId,
	}

	message, err := server.store.StoreMessage(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, message)
}

type updateMessageRequest struct {
	Id       int32  `json:"id" binding:"required"`
	SenderId int32  `json:"sender_id" binding:"required"`
	Message  string `json:"message" binding:"required"`
}

func (server *Server) updateMessage(ctx *gin.Context) {

	var req updateMessageRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateMessageParams{
		ID:       req.Id,
		SenderID: req.SenderId,
		Message:  req.Message,
	}

	message, err := server.store.UpdateMessage(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, message)
}

type deleteMessageRequest struct {
	Id       int32 `json:"id" binding:"required"`
	SenderId int32 `json:"sender_id" binding:"required"`
}

func (server *Server) deleteMessage(ctx *gin.Context) {

	var req deleteMessageRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.DeleteMessageParams{
		ID:       req.Id,
		SenderID: req.SenderId,
	}

	if err := server.store.DeleteMessage(ctx, arg); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
}
