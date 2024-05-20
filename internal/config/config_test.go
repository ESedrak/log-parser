package config

import (
	"reflect"
	"testing"
)

func Test_loadConfig(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    Config
		wantErr bool
	}{
		{
			name: "Success scenario",
			args: args{
				filePath: "internal/config/testdata/good.json",
			},
			want:    getExpectedConfig(),
			wantErr: false,
		},
		{
			name: "read config error",
			args: args{
				filePath: "wrong",
			},
			want:    Config{},
			wantErr: true,
		},
		{
			name: "unmarshal bad json",
			args: args{
				filePath: "internal/config/testdata/bad.json",
			},
			want:    Config{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadConfig(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getExpectedConfig() Config {
	return Config{
		RequestedNum: RequestedNum{
			IP:  3,
			URL: 3,
		},
		Regex: Regex{
			MatchIPsURLsIgnoreQuery: "(\\d+\\.\\d+\\.\\d+\\.\\d+).+(?:GET|POST|PUT|DELETE|HEAD)\\s([^ ?]+)",
		},
		Path: Path{
			LogPath: "logs/log_file.log",
		},
	}
}
