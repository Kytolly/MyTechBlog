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
	ERROR_TOKEN_NOT_EXIST   = 1007 // token不存在
	ERROR_USER_NO_RIGHT		= 1008 // 用户无管理权限
	// CODE=2000 + XX 分类模块的错误
	ERROR_CATEGORY_USED 	= 2001 // 分类重复
	ERROR_CATEGORY_NOT_EXIST= 2002 // 分类不存在
	// CODE=3000 + XX 文章模块的错误 
	ERROR_ARTICLE_NOT_EXIST = 3001 //文章不在指定分类中
)

var CodeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR: "FAIL",
	ERROR_USERNAME_USED: "the user is already exsited!", 	
	ERROR_PASSWORD_WRONG: "the password or account is wrong!",
	ERROR_USER_NOT_EXIST: "the user is not existed!",
	ERROR_TOKEN_INVALID: "the token becomes invalid!",
	ERROR_TOKEN_WRONG: "the token is wrong!",
	ERROR_TOKEN_NOT_EXIST: "the token is not existed!",
	ERROR_TOKEN_TYPE_WRONG: "the token's type is wrong!",
	ERROR_CATEGORY_USED: "the category is already exsited!",
	ERROR_CATEGORY_NOT_EXIST: "the category is not exists!",
	ERROR_ARTICLE_NOT_EXIST: "the article is not exsited!",
	ERROR_USER_NO_RIGHT: "the user dont have enough right!",
}

func GetErrorMessage(code int) string{
	return CodeMsg[code]
}