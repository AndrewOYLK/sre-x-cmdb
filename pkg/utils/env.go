package utils

import "os"

func Getenv(name string, value string) string {
	v := os.Getenv(name)
	if len(v) == 0 {
		return value
	}
	return v
}

func Setenv(name string, value string) {
	v := os.Getenv(name)
	if len(v) == 0 {
		os.Setenv(name, value)
	}
}
