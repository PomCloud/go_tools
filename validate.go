package go_tools

import (
	"regexp"
	"strings"
)

func EmptyString(text string) bool {
	str := strings.Replace(strings.Replace(text, " ", "", -1), "\n", "", -1)
	if len(str) <= 0 {
		return false
	}
	return true
}

func ValidInt(t int) bool {
	if t <= 0 {
		return false
	}
	return true
}

func ValidIP4(ipAddress string) bool {
	ipAddress = strings.Trim(ipAddress, " ")
	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if re.MatchString(ipAddress) {
		return true
	}
	return false
}
