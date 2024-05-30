package user

import (
	"gin-rush-template/internal/global/database"
	"gin-rush-template/test"
	"gin-rush-template/tools"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreate(t *testing.T) {
	testEmail := "test@test.com"
	testPassword := "123456"

	test.SetupEnvironment(t)
	req := createRequest{
		User: User{
			Email:    testEmail,
			Password: testPassword,
		},
	}

	resp := test.DoRequest(t, Create, req)
	test.NoError(t, resp)

	u := database.Query.User
	userInfo, err := u.Where(u.Email.Eq(testEmail)).First()
	require.NoError(t, err)
	require.Equal(t, false, tools.Compare(testPassword, userInfo.Password))
}
