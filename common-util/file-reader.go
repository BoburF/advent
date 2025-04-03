package common_util

import (
	"fmt"
	"os"
	"path"
)

func ReadFile(year, day int) ([]byte, error) {
	root, err := FindProjectRoot()
	if err != nil {
		return nil, err
	}

	pathToFile := path.Join(root, "input-files", fmt.Sprint(year), fmt.Sprint(day, ".txt"))
	return os.ReadFile(pathToFile)
}
