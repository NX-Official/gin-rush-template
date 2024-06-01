package errs

import (
	"encoding/json"
	"fmt"
	"gin-rush-template/config"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

type ResponseBody struct {
	Code   int32  `json:"code"`
	Msg    string `json:"msg"`
	Origin string `json:"origin,omitempty"`
	Data   any    `json:"data,omitempty"`
}

func Success(c *gin.Context, data ...any) {
	response := ResponseBody{
		Code: success.Code,
		Msg:  success.Message,
		Data: nil,
	}
	if len(data) > 0 {
		response.Data = data[0]
	}
	writeResponse(c, http.StatusOK, response)
}

func Fail(c *gin.Context, err error) {
	var response ResponseBody

	var e *Error
	ok := errors.As(err, &e)
	if !ok {
		e = serverInternal.WithOrigin(err)
	}

	response.Code = e.Code
	response.Msg = e.Message

	if config.Get().Mode == config.DebugMode {
		response.Origin = e.Origin
	}
	c.Set(ErrorContextKey, *e)
	writeResponse(c, int(e.Code/100), response)
	c.Abort()
}

func writeResponse(c *gin.Context, code int, response ResponseBody) {
	body, _ := json.Marshal(response)
	span := trace.SpanFromContext(c.Request.Context())
	span.SetAttributes(attribute.KeyValue{Key: "http.response.body", Value: attribute.StringValue(string(body))})
	c.JSON(code, response)
}

func Recovery(c *gin.Context) {
	info := recover()
	if info != nil {
		err, ok := info.(error)
		if ok {
			Fail(c, errors.WithStack(err))
		} else {
			Fail(c, errors.New(fmt.Sprintf("%+v", info)))
		}
	}
}
