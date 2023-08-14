package goalipay

func IsError(code string) bool {
	if code != "10000" {
		return true
	}
	return false
}
