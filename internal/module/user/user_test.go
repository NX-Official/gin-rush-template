package user

import (
	"gin-rush-template/internal/global/database"
	"gin-rush-template/internal/global/errs"
	"gin-rush-template/test"
	"gin-rush-template/tools"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreate(t *testing.T) {
	test.SetupEnvironment(t)
	t.Run("Success", func(t *testing.T) {
		req := CreateRequest{
			User: User{
				Email:    "test@test.com",
				Password: "123456",
			},
		}
		resp := test.DoRequest(t, Create, req)
		test.NoError(t, resp)
		u := database.Query.User
		userInfo, err := u.Where(u.Email.Eq(req.Email)).First()
		require.NoError(t, err)
		require.Equal(t, true, tools.Compare(req.Password, userInfo.Password))
	})
	t.Run("PassWordTooShort", func(t *testing.T) {
		req := CreateRequest{
			User: User{
				Email:    "test@test.com",
				Password: "123",
			},
		}
		resp := test.DoRequest(t, Create, req)
		test.ErrorEqual(t, errs.InvalidRequest.WithTips(
			`Key: 'CreateRequest.User.Password' Error:Field validation for 'Password' failed on the 'min' tag`,
		), resp)
	})
	t.Run("EmailInvalid", func(t *testing.T) {
		req := CreateRequest{
			User: User{
				Email:    "test",
				Password: "123456",
			},
		}
		resp := test.DoRequest(t, Create, req)
		test.ErrorEqual(t, errs.InvalidRequest.WithTips(
			`Key: 'CreateRequest.User.Email' Error:Field validation for 'Email' failed on the 'email' tag`,
		), resp)
	})
}
