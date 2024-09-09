package common

import (
	"os"
)

// GetFileSize get the length in bytes of file of the specified path.
func GetFileSize(path string) int64 {
	fi, err := os.Stat(path)
	if nil != err {
		return -1
	}
	return fi.Size()
}

// IsExist determines whether the file specified by the given path is exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsBinary determines whether the specified content is a binary file content.
func IsBinary(content string) bool {
	for _, b := range content {
		if 0 == b {
			return true
		}
	}
	return false
}

// IsDir determines whether the specified path is a directory.
func IsDir(path string) bool {
	fio, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return false
	}

	if nil != err {
		Warn("determines whether [%s] is a directory failed: [%v]", path, err)
		return false
	}
	return fio.IsDir()
}
