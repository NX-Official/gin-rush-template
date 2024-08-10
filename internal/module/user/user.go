package user

import (
	"gin-rush-template/internal/global/database"
	"gin-rush-template/internal/global/errs"
	"gin-rush-template/internal/model"
	"gin-rush-template/tools"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type User struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=12"`
}

type CreateRequest struct {
	User
}

func Create(c *gin.Context) {
	var req CreateRequest
	if err := c.BindJSON(&req); err != nil {
		errs.Fail(c, errs.InvalidRequest.WithTips(err.Error()))
		return
	}
	user := &model.User{}
	_ = copier.Copy(user, &req)
	log.Info("Creating User", "user email", user.Email)
	err := database.Query.User.WithContext(c.Request.Context()).Create(user)
	switch {
	case tools.IsDuplicateKeyError(err):
		log.Info("email exist", "email", user.Email)
		errs.Fail(c, errs.HasExist.WithOrigin(err))
		return
	case err != nil:
		log.Error("database error", "error", err)
		errs.Fail(c, errs.DatabaseError.WithOrigin(err))
		return
	}

	errs.Success(c, nil)
}
