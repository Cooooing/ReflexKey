package common

import (
	"bytes"
	"errors"
	"github.com/shirou/gopsutil/v3/host"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

// IsWindows determines whether current OS is Windows.
func (g *GuOs) IsWindows() bool {
	return "windows" == runtime.GOOS
}

// IsLinux determines whether current OS is Linux.
func (g *GuOs) IsLinux() bool {
	return "linux" == runtime.GOOS
}

// IsDarwin determines whether current OS is Darwin.
func (g *GuOs) IsDarwin() bool {
	return "darwin" == runtime.GOOS
}

// Home returns the home directory for the executing user.
//
// This uses an OS-specific method for discovering the home directory.
// An error is returned if a home directory cannot be detected.
func (os *GuOs) Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	// cross compile support

	if os.IsWindows() {
		return os.homeWindows()
	}

	// Unix-like system, so just assume Unix
	return os.homeUnix()
}

func (g *GuOs) homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func (g *GuOs) homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}

func (g *GuOs) GetOSPlatform() (plat string) {
	plat, _, _, err := host.PlatformInformation()
	if nil != err {
		Log.Warn("get os platform failed: %s", err)
		return "Unknown"
	}
	return
}
