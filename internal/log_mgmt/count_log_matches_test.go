package log_mgmt

import (
	"log-parser/internal/ip_mgmt"
	"log-parser/internal/url_mgmt"
	"reflect"
	"testing"
)

func TestCountLogMatchesIgnoresQuery(t *testing.T) {
	type args struct {
		logMatches [][]string
	}
	tests := []struct {
		name  string
		args  args
		want  *url_mgmt.URLCounter
		want1 *ip_mgmt.IPCounter
	}{
		{
			name: "Success: test with multiple log matches",
			args: args{
				logMatches: [][]string{
					{"192.168.1.1", "", "GET", "/example/"},
					{"193.168.1.1", "", "GET", "/example/"},
					{"192.168.1.1", "", "PUT", "/example/?query=true"},
					{"192.168.1.1", "", "POST", "/example2/"},
				},
			},
			want: &url_mgmt.URLCounter{
				URLCounts: map[string]int{
					"/example/":  3,
					"/example2/": 1,
				},
			},
			want1: &ip_mgmt.IPCounter{
				IPCounts: map[string]int{
					"192.168.1.1": 3,
					"193.168.1.1": 1,
				},
			},
		},
		{
			name: "Success: unwanted HTTP method is not counted",
			args: args{
				logMatches: [][]string{
					{"192.168.1.1", "", "PATCH", "/example/"},
				},
			},
			want: &url_mgmt.URLCounter{
				URLCounts: map[string]int{},
			},
			want1: &ip_mgmt.IPCounter{
				IPCounts: map[string]int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CountLogMatchesIgnoresQuery(tt.args.logMatches)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountLogMatchesNoQuery() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CountLogMatchesNoQuery() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
