package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	yogo "yogo.io/go-tripPlanner-backend/db"
)

type createUserRequest struct {
	Email      string `json:"email" binding:"required"`
	Username   string `json:"username" binding:"required"`
	ProfilePic string `json:"profilePic"`
	Status     int32  `json:"status" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := yogo.CreateUserParams{
		Email:    req.Email,
		Username: req.Username,
		Status:   req.Status,
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
