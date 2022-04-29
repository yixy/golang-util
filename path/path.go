package path

import (
	"os"
	"os/exec"
	"path/filepath"
)

var proPath string

// GetProgramPath
func GetProgramPath() (string, error) {
	if proPath != "" {
		return proPath, nil
	}
	dir, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	if p, err := os.Readlink(dir); err == nil {
		dir = p
	}
	proPath, err = filepath.Abs(filepath.Dir(dir))
	if err != nil {
		return "", err
	}
	return proPath, nil
}
