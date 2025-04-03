package common_util

import (
	"os"
)

func FindProjectRoot() (string, error) {
	return os.Getwd()
}
