package goalipay

import "errors"

var (
	MissParamErr  = errors.New("缺少必需的参数")
	SignErr       = errors.New("签名错误")
	VerifySignErr = errors.New("签名验证错误")
)

func IsError(code string) bool {
	if code != "10000" {
		return true
	}
	return false
}
