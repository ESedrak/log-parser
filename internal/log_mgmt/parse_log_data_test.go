package log_mgmt

import (
	"reflect"
	"testing"
)

func TestParseLogData(t *testing.T) {
	type args struct {
		logsData []byte
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "Success: valid log data",
			args: args{
				[]byte(`177.71.128.21 - - [10/Jul/2018:22:22:08 +0200] "GET /blog/2018/08/survey-your-opinion-matters/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6"`),
			},
			want: [][]string{
				{"177.71.128.21", "10/Jul/2018:22:22:08 +0200", "GET", "/blog/2018/08/survey-your-opinion-matters/", "HTTP/1.1", "200", "3574", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6"},
			},
			wantErr: false,
		},
		{
			name: "Error: no log data matches",
			args: args{
				nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLogData(tt.args.logsData)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseLogData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLogData() = %v, want %v", got, tt.want)
			}
		})
	}
}
