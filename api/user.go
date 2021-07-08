package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	yogo "yogo.io/go-tripPlanner-backend/db"
	"yogo.io/go-tripPlanner-backend/util"
)

type createUserRequest struct {
	Email      string         `json:"email" binding:"required"`
	Username   string         `json:"username" binding:"required,alphanum"`
	Password   string         `json:"password" binding:"required"`
	ProfilePic sql.NullString `json:"profilePic"`
	Status     int32          `json:"status" binding:"required"`
}

type createUserResponse struct {
	ID         int64          `json:"id"`
	Email      string         `json:"email"`
	Username   string         `json:"username"`
	Status     int32          `json:"status"`
	ProfilePic sql.NullString `json:"profilePic"`
	Createdat  time.Time      `json:"createdat"`
	Updatedat  sql.NullTime   `json:"updatedat"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := yogo.CreateUserParams{
		Email:      req.Email,
		Username:   req.Username,
		Password:   hashedPassword,
		ProfilePic: req.ProfilePic,
		Status:     req.Status,
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {

		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := createUserResponse{
		ID:         user.ID,
		Email:      user.Email,
		Username:   user.Username,
		Status:     user.Status,
		ProfilePic: user.ProfilePic,
		Createdat:  user.Createdat,
		Updatedat:  user.Updatedat,
	}
	ctx.JSON(http.StatusOK, rsp)
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
	rsp := createUserResponse{
		ID:         user.ID,
		Email:      user.Email,
		Username:   user.Username,
		Status:     user.Status,
		ProfilePic: user.ProfilePic,
		Createdat:  user.Createdat,
		Updatedat:  user.Updatedat,
	}
	ctx.JSON(http.StatusOK, rsp)
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
