package string

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMonthList(t *testing.T) {
	type args struct {
		start string
		end   string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "aaa",
			args: args{start: "201904", end: "202203"},
			want: []string{"201904", "201905", "201906", "201907", "201908", "201909", "201910", "201911", "201912", "202001", "202002", "202003", "202004", "202005", "202006", "202007", "202008", "202009", "202010", "202011", "202012", "202101", "202102", "202103", "202104", "202105", "202106", "202107", "202108", "202109", "202110", "202111", "202112", "202201", "202202", "202203"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MonthList(tt.args.start, tt.args.end)
			fmt.Println(result)
			if got := result; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MonthList() = %v, want %v", got, tt.want)
			}
		})
	}
}
