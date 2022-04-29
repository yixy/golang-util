package encrypt

import (
	"testing"
)

func TestLoadKey(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{name: "1", args: args{file: "pemfile/pkcs1_pri.pem"}, wantErr: false},
		{name: "2", args: args{file: "pemfile/pkcs1_pub.pem"}, wantErr: false},
		{name: "3", args: args{file: "pemfile/pkcs8_pri.pem"}, wantErr: false},
		{name: "4", args: args{file: "pemfile/x509_pub.pem"}, wantErr: false},
		{name: "5", args: args{file: "pemfile/wrong.pem"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LoadKey(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("LoadKey() = %v, want %v", got, tt.want)
			// }
		})
	}
}
