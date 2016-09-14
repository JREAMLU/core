package io

import (
	"io/ioutil"
	"os"
)

func ReadAll(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

func ReadAllBytes(path string) ([]byte, error) {
	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	return ioutil.ReadAll(fi)
}

func WriteFile(path, content string, isOverride bool) error {
	var flag int
	if isOverride {
		flag = os.O_CREATE | os.O_RDWR | os.O_TRUNC
	} else {
		flag = os.O_CREATE | os.O_EXCL | os.O_RDWR
	}
	f, err := os.OpenFile(path, flag, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(content)
	return nil
}

func MkdireAll(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// create controller directory
		if err := os.MkdirAll(path, 0777); err != nil {
			return err
		}
	}
	return nil
}

func CheckFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsExist(err) {
		return true
	}
	return false
}

func CreateFileOfTrunc(filePath string) (*os.File, error) {
	if CheckFileExists(filePath) {
		return os.OpenFile(filePath, os.O_TRUNC, 0666)
	}
	return os.Create(filePath)
}

func CreateFileOfAppend(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
}
