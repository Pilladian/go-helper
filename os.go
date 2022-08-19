package helper

import (
	"errors"
	"os"
	"strings"
)

// Takes input string "path" and returns wether the file/directory exists on the hosts system.
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Takes input string "path" and recursively creates the directories on the host system. If an error occurs, it will be returned.
func CreatePath(path string) error {
	existent, err := Exists(path)
	if err != nil {
		return errors.New(err.Error())
	}

	if !existent {
		split_path := strings.Split(path, "/")
		CreatePath(strings.Join(split_path[0:len(split_path)-1], "/"))

		err := os.Mkdir(path, 0700)
		if err != nil {
			return errors.New(err.Error())
		}
	}
	return nil
}
