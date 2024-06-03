package user

import (
	"gin-rush-template/internal/global/database"
	"gin-rush-template/internal/global/errs"
	"gin-rush-template/internal/global/jwt"
	"gin-rush-template/tools"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	var req User
	if err := c.BindJSON(&req); err != nil {
		errs.Fail(c, errs.InvalidRequest.WithTips(err.Error()))
		return
	}
	u := database.Query.User
	userInfo, err := u.WithContext(c.Request.Context()).Where(u.Email.Eq(req.Email)).First()
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		errs.Fail(c, errs.NotFound)
		return
	case err != nil:
		errs.Fail(c, errs.DatabaseError.WithOrigin(err))
		return
	}

	if !tools.Compare(req.Password, userInfo.Password) {
		errs.Fail(c, errs.InvalidPassword)
		return
	}
	errs.Success(c, map[string]string{
		"token": jwt.CreateToken(jwt.Payload{UserId: userInfo.ID}),
	})
}
