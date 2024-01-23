package errs

// This error code refers to the semantics of HTTP status codes
// to facilitate the identification of error types

// Type |Type description
// 1xx: Informational - Request received, continuing process
// 2xx: Success - The action was successfully received, understood, and accepted
// 3xx: Redirection - Further action must be taken in order to complete the request
// 4xx: Client Error - The request contains bad syntax or cannot be fulfilled
// 5xx: Server Error - The server failed to fulfill an apparently valid request

// 本错误码参照了HTTP状态码的语义，方便识别错误类型

// 分类 |分类描述
// 1**	信息，服务器收到请求，需要请求者继续执行操作
// 2**	成功，操作被成功接收并处理
// 3**	重定向，需要进一步的操作以完成请求
// 4**	客户端错误，请求包含语法错误或无法完成请求
// 5**	服务器错误，服务器在处理请求的过程中发生了错误

// 200 OK
var (
	success = newError(200, "Success")
)

// 400 BAD REQUEST
var (
	InvalidRequest  = newError(40001, "无效的请求")
	NotFound        = newError(40002, "目标不存在")
	HasExist        = newError(40003, "目标已存在")
	InvalidPassword = newError(40004, "密码错误")
	ErrTokenInvalid = newError(40005, "无效的token")
)

// 500 INTERNAL ERROR
var (
	serverInternal = newError(50001, "服务器内部错误")
	DatabaseError  = newError(50002, "数据库错误")
)
