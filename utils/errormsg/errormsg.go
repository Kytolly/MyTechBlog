package errormsg

const (
	SUCCESS = 200
	ERROR   = 500

	// CODE=1000 + XX 用户模块错误
	ERROR_USERNAME_USED 	= 1001 // 用户重复
	ERROR_PASSWORD_WRONG 	= 1002 // 密码错误
	ERROR_USER_NOT_EXIST 	= 1003 // 用户不存在
	ERROR_TOKEN_INVALID 	= 1004 // token超时或过期
	ERROR_TOKEN_WRONG       = 1005 // token错误
	ERROR_TOKEN_TYPE_WRONG  = 1006 // token格式错误

	// CODE=2000 + XX 文章模块的错误

	// CODE=3000 + XX 分类模块的错误 
)

var CodeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR: "FAIL",
	ERROR_USERNAME_USED: "the user is already exsited!", 	
	ERROR_PASSWORD_WRONG: "the password or account is wrong!",
	ERROR_USER_NOT_EXIST: "the user is not existed!",
	ERROR_TOKEN_INVALID: "the token becomes invalid!",
	ERROR_TOKEN_WRONG: "the token is wrong!",
	ERROR_TOKEN_TYPE_WRONG: "the token's type is wrong!",
}

func GetErrorMessage(code int) string{
	return CodeMsg[code]
}