package str

import (
	"testing"
)

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

func TestMatch(t *testing.T) {
	type args struct {
		src     string
		mathers []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1",
			args: args{src: "test", mathers: []string{"foo", "bar"}},
			want: false,
		},
		{name: "case2",
			args: args{src: "bar", mathers: []string{"foo", "bar"}},
			want: true,
		},
		{name: "case3",
			args: args{src: "foo", mathers: []string{"foo", "bar"}},
			want: true,
		},
		{name: "case4",
			args: args{src: "test", mathers: []string{"test"}},
			want: true,
		},
		{name: "case5",
			args: args{src: "test", mathers: []string{"foo"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Match(tt.args.src, tt.args.mathers...); got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
