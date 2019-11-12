package utils

type ErrCode int

const (
	ERR_INVALID_PARAM ErrCode = 1 // 非法参数
	ERR_INVALID_COOKIE ErrCode = 2 // 无效cookie
	// ...
)