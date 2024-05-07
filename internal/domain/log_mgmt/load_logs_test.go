package log_mgmt

import (
	"errors"
	"testing"
)

func TestLoadLogs(t *testing.T) {
	tests := []struct {
		name        string
		args        string
		want        string
		wantErr     error
		wantErrBool bool
	}{
		{
			name:        "Success: load one log",
			args:        "./testdata/good.log",
			want:        "177.71.128.21 - - [10/Jul/2018:22:22:08 +0200] \"GET /blog/2018/08/survey-your-opinion-matters/ HTTP/1.1\" 200 3574 \"-\" \"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6\"",
			wantErr:     nil,
			wantErrBool: false,
		},
		{
			name:        "Error: bad filepath",
			args:        "./incorrect/filepath.log",
			want:        "",
			wantErr:     errors.New("open ./incorrect/filepath.log: no such file or directory"),
			wantErrBool: true,
		},
		{
			name:        "empty log",
			args:        "./testdata/empty.log",
			want:        "",
			wantErr:     nil,
			wantErrBool: false,
		},
	}
	for _, tt := range tests {
		// reference: https://www.timothyomargheim.com/posts/testing-channels-in-go/

		t.Run(tt.name, func(t *testing.T) {
			//setup for one log
			logChan := make(chan string, 1)
			errChan := make(chan error, 1)

			// execute
			go LoadLogs(tt.args, logChan, errChan)

			// receive log/error
			got := <-logChan
			gotErr1 := <-errChan

			// check to see if an error is received when no error is expected
			if (gotErr1 != nil) != tt.wantErrBool {
				t.Errorf("LoadLogs() error = %v, wantErr %v", gotErr1, tt.wantErr)
				return
			}
			// check the error message
			if tt.wantErrBool && gotErr1 != nil {
				if gotErr1.Error() != tt.wantErr.Error() {
					t.Errorf("LoadLogs() error message = %v, want %v", gotErr1, tt.wantErr)
				}
			}
			// check no errors and the received outcome matches the expected outcome
			if !tt.wantErrBool && got != tt.want {
				t.Errorf("LoadLogs() got = %v, want %v", got, tt.want)
			}
		})
	}
}
