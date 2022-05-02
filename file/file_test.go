//go:build linux || darwin
// +build linux darwin

package file

import (
	"os"
	"testing"
)

func TestLockFile(t *testing.T) {
	fileName := "/tmp/TestLockFile.test"
	f, err := os.Create(fileName)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(fileName)
	defer f.Close()
	type args struct {
		file *os.File
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "case1",
			args:    args{file: f},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LockFile(tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("LockFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			err = UnlockFile(f)
			if err != nil {
				t.Errorf("UnlockFile = %v", err)
			}
		})
	}
}
