package test

import (
	"gin-rush-template/internal/global/errs"
	"github.com/stretchr/testify/require"
	"testing"
)

func ErrorEqual(t *testing.T, expected *errs.Error, resp errs.ResponseBody) {
	require.Equal(t, expected.Code, resp.Code)
	require.Equal(t, expected.Message, resp.Msg)
}

func NoError(t *testing.T, resp errs.ResponseBody) {
	require.Equal(t, int32(200), resp.Code)
}
