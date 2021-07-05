package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	yogo "yogo.io/go-tripPlanner-backend/db"
)

type createUserRequest struct {
	Email      string         `json:"email" binding:"required"`
	Username   string         `json:"username" binding:"required"`
	ProfilePic sql.NullString `json:"profilePic"`
	Status     int32          `json:"status" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := yogo.CreateUserParams{
		Email:      req.Email,
		Username:   req.Username,
		ProfilePic: req.ProfilePic,
		Status:     req.Status,
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type getUserRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// type listUsersRequest struct {
// 	pageID   int32 `form:"page_id" binding:"required"`
// 	pageSize int32 `form:"page_size" binding:"required"`
// }

func (server *Server) listUsers(ctx *gin.Context) {
	// var req listUsersRequest
	// if err := ctx.ShouldBindQuery(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	// 	return
	// }

	// arg := yogo.Lis
	users, err := server.store.ListUsers(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}
