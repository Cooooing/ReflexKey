package common

import (
	"errors"
	"os"
	"strings"
	"sync"
)

var (
	lockMutex      = sync.Mutex{}
	operatingFiles = map[string]*sync.Mutex{}
)

func (g *GuFileLock) Lock(filePath string) {
	g.lock(filePath)
}

func (g *GuFileLock) Unlock(filePath string) {
	g.unlock(filePath)
}

func (g *GuFileLock) OpenFile(filePath string, flag int, perm os.FileMode) (file *os.File, err error) {
	g.lock(filePath)

	file, err = os.OpenFile(filePath, flag, perm)
	if g.isDenied(err) {
		Log.Fatal(ExitCodeFileSysErr, "open file [%s] failed: %s", filePath, err)
		return
	}
	return
}

func (g *GuFileLock) CloseFile(file *os.File) (err error) {
	if nil == file {
		return
	}

	defer g.unlock(file.Name())

	err = file.Close()
	if g.isDenied(err) {
		Log.Fatal(ExitCodeFileSysErr, "close file [%s] failed: %s", file.Name(), err)
		return
	}
	return
}

func (g *GuFileLock) IsExist(filePath string) (ret bool) {
	g.lock(filePath)
	defer g.unlock(filePath)

	return File.IsExist(filePath)
}

func (g *GuFileLock) ReadFile(filePath string) (data []byte, err error) {
	g.lock(filePath)
	defer g.unlock(filePath)

	data, err = os.ReadFile(filePath)
	if g.isDenied(err) {
		Log.Fatal(ExitCodeFileSysErr, "read file [%s] failed: %s", filePath, err)
		return
	}
	return
}

func (g *GuFileLock) WriteFile(filePath string, data []byte) (err error) {
	g.lock(filePath)
	defer g.unlock(filePath)

	err = File.WriteFileSafer(filePath, data, 0644)
	if g.isDenied(err) {
		Log.Fatal(ExitCodeFileSysErr, "write file [%s] failed: %s", filePath, err)
		return
	}
	return
}

func (g *GuFileLock) isDenied(err error) bool {
	if nil == err {
		return false
	}

	if errors.Is(err, os.ErrPermission) {
		return true
	}

	errMsg := strings.ToLower(err.Error())
	return strings.Contains(errMsg, "access is denied") || strings.Contains(errMsg, "used by another process")
}

func (g *GuFileLock) lock(filePath string) {
	lockMutex.Lock()
	mutex := operatingFiles[filePath]
	if nil == mutex {
		mutex = &sync.Mutex{}
		operatingFiles[filePath] = mutex
	}
	lockMutex.Unlock()
	mutex.Lock()
}

func (g *GuFileLock) unlock(filePath string) {
	lockMutex.Lock()
	mutex := operatingFiles[filePath]
	//delete(operatingFiles, filePath) 删了的话并发情况下会导致死锁，得考虑新的回收机制
	lockMutex.Unlock()
	if nil != mutex {
		mutex.Unlock()
	}
}
