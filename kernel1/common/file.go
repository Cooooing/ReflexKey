package common

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// GetFileSize get the length in bytes of file of the specified path.
func (g *GuFile) GetFileSize(path string) int64 {
	fi, err := os.Stat(path)
	if nil != err {
		return -1
	}
	return fi.Size()
}

// IsExist determines whether the file specified by the given path is exists.
func (g *GuFile) IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsBinary determines whether the specified content is a binary file content.
func (g *GuFile) IsBinary(content string) bool {
	for _, b := range content {
		if 0 == b {
			return true
		}
	}
	return false
}

// IsDir determines whether the specified path is a directory.
func (g *GuFile) IsDir(path string) bool {
	fio, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return false
	}

	if nil != err {
		Log.Warn("determines whether [%s] is a directory failed: [%v]", path, err)
		return false
	}
	return fio.IsDir()
}

// WriteFileSaferByReader writes the data to a temp file and atomically move if everything else succeeds.
func (g *GuFile) WriteFileSaferByReader(writePath string, reader io.Reader, perm os.FileMode) (err error) {
	dir, name := filepath.Split(writePath)
	tmp := filepath.Join(dir, name+Rand.String(7)+".tmp")
	f, err := os.OpenFile(tmp, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if nil != err {
		return
	}

	if _, err = io.Copy(f, reader); nil != err {
		return
	}

	if err = f.Sync(); nil != err {
		return
	}

	if err = f.Close(); nil != err {
		return
	}

	if err = os.Chmod(f.Name(), perm); nil != err {
		return
	}

	for i := 0; i < 3; i++ {
		err = os.Rename(f.Name(), writePath) // Windows 上重命名是非原子的
		if nil == err {
			_ = os.Remove(f.Name())
			return
		}

		if errMsg := strings.ToLower(err.Error()); strings.Contains(errMsg, "access is denied") || strings.Contains(errMsg, "used by another process") { // 文件可能是被锁定
			time.Sleep(200 * time.Millisecond)
			continue
		}
		break
	}
	return
}

// WriteFileSafer writes the data to a temp file and atomically move if everything else succeeds.
func (g *GuFile) WriteFileSafer(writePath string, data []byte, perm os.FileMode) (err error) {
	dir, name := filepath.Split(writePath)
	tmp := filepath.Join(dir, name+Rand.String(7)+".tmp")
	f, err := os.OpenFile(tmp, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if nil != err {
		return
	}

	if _, err = f.Write(data); nil != err {
		return
	}

	if err = f.Sync(); nil != err {
		return
	}

	if err = f.Close(); nil != err {
		return
	}

	if err = os.Chmod(f.Name(), perm); nil != err {
		return
	}

	for i := 0; i < 3; i++ {
		err = os.Rename(f.Name(), writePath) // Windows 上重命名是非原子的
		if nil == err {
			_ = os.Remove(f.Name())
			return
		}

		if errMsg := strings.ToLower(err.Error()); strings.Contains(errMsg, "access is denied") || strings.Contains(errMsg, "used by another process") { // 文件可能是被锁定
			time.Sleep(200 * time.Millisecond)
			continue
		}
		break
	}
	return
}
