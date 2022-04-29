package string

import (
	"testing"
)

func TestRandomString(t *testing.T) {
	result := RandomString(32)
	if len(result) != 32 {
		t.Error(len(result))
	}
	t.Log(result)

}

func TestHasPrefixes(t *testing.T) {
	type args struct {
		src      string
		prefixes []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1",
			args: args{src: "foo_test", prefixes: []string{"foo"}},
			want: true,
		},
		{name: "case2",
			args: args{src: "foo", prefixes: []string{"fooo"}},
			want: false,
		},
		{name: "case3",
			args: args{src: "foo", prefixes: []string{"foo"}},
			want: true,
		},
		{name: "case4",
			args: args{src: "foo_test", prefixes: []string{"foo", "bar"}},
			want: true,
		},
		{name: "case5",
			args: args{src: "bar_test", prefixes: []string{"foo", "bar"}},
			want: true,
		},
		{name: "case6",
			args: args{src: "test", prefixes: []string{"foo", "bar"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasPrefixes(tt.args.src, tt.args.prefixes...); got != tt.want {
				t.Errorf("HasPrefixes() = %v, want %v", got, tt.want)
			}
		})
	}
}
