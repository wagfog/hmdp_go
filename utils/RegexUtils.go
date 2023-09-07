package utils

import "regexp"

func IsPhoneInvalid(phone string) bool {
	reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return !rgx.MatchString(phone)
}
