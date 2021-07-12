package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	yogo "yogo.io/go-tripPlanner-backend/db"
	"yogo.io/go-tripPlanner-backend/token"
	"yogo.io/go-tripPlanner-backend/util"
)

type createUserRequest struct {
	Email      string         `json:"email" binding:"required"`
	Username   string         `json:"username" binding:"required,alphanum"`
	Password   string         `json:"password" binding:"required"`
	ProfilePic sql.NullString `json:"profilePic"`
	Status     int32          `json:"status" binding:"required"`
}

type userResponse struct {
	ID         int64          `json:"id"`
	Email      string         `json:"email"`
	Username   string         `json:"username"`
	Status     int32          `json:"status"`
	ProfilePic sql.NullString `json:"profilePic"`
	Createdat  time.Time      `json:"createdat"`
	Updatedat  sql.NullTime   `json:"updatedat"`
}

func newUserResponse(user yogo.User) userResponse {
	return userResponse{
		ID:         user.ID,
		Email:      user.Email,
		Username:   user.Username,
		Status:     user.Status,
		ProfilePic: user.ProfilePic,
		Createdat:  user.Createdat,
		Updatedat:  user.Updatedat,
	}
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

	rsp := newUserResponse(user)
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
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if user.Email != authPayload.Username {
		err := errors.New("you are unauthorized to view this account")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	rsp := newUserResponse(user)
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

type loginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.CheckUser(ctx, req.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(
		user.Email,
		server.config.AccessTokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, rsp)
}
