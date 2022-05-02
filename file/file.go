//go:build linux || darwin
// +build linux darwin

package file

import (
	"os"
	"syscall"
)

func LockFile(file *os.File) error {
	//exclude and no-block
	return syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
}

func UnlockFile(file *os.File) error {
	//unlock
	return syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
}
