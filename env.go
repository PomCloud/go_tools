package go_tools

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Get env var `name`  `defvalue` or return an error if it's missing.
func Get(name, defvalue string) string {
	if s := os.Getenv(strings.ToUpper(name)); s == "" {
		return defvalue
	} else {
		return s
	}
}

// GetInt env var `name`  `defvalue` or return int and error if it's missing.
func GetInt(name string, defvalue int) int {
	env := os.Getenv(strings.ToUpper(name))
	if env == "" {
		return defvalue
	}
	v, err := strconv.Atoi(env)
	if err != nil {
		fmt.Println(fmt.Sprintf("%s Error: os.Getenv %s <-> %d", "GetInt", strings.ToUpper(name), defvalue))
		os.Exit(0)
	}
	return int(v)
}

// GetBool env var `name`  `defvalue` or return and bool if it's missing.
func GetBool(name string, defvalue bool) bool {
	env := os.Getenv(strings.ToUpper(name))
	if env == "" {
		return defvalue
	}
	if strings.ToLower(env) == "true" {
		return true
	}
	return false
}
