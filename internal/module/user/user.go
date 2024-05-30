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

type createRequest struct {
	User
}

func Create(c *gin.Context) {
	var req createRequest
	if err := c.BindJSON(&req); err != nil {
		errs.Fail(c, errs.InvalidRequest.WithTips(err.Error()))
		return
	}
	user := &model.User{}
	_ = copier.Copy(user, &req)
	err := database.Query.User.Create(user)
	switch {
	case tools.IsDuplicateKeyError(err):
		errs.Fail(c, errs.HasExist)
		return
	case err != nil:
		errs.Fail(c, errs.DatabaseError.WithOrigin(err))
		return
	}

	errs.Success(c, nil)
}
