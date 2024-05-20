package log_mgmt

import (
	"log-parser/internal/config"
	"reflect"
	"testing"
)

func init() {
	config.Init("/config/config.json")
}

func TestCountLogMatch(t *testing.T) {
	// reference: https://www.timothyomargheim.com/posts/testing-channels-in-go/
	tests := []struct {
		name  string
		args  []string
		want  map[string]int
		want1 map[string]int
	}{
		{
			name: "Success: log match",
			args: []string{
				"177.71.128.21 - - [10/Jul/2018:22:22:08 +0200] \"GET /blog/2018/08/survey-your-opinion-matters/ HTTP/1.1\" 200 3574 \"-\" \"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6\"",
			},
			want: map[string]int{
				"/blog/2018/08/survey-your-opinion-matters/": 1,
			},
			want1: map[string]int{
				"177.71.128.21": 1,
			},
		},
		{
			name: "Warn: unwanted HTTP method is not counted/log not parsed",
			args: []string{
				"177.71.128.21 - - [10/Jul/2018:22:22:08 +0200] \"PATCH /blog/2018/08/survey-your-opinion-matters/ HTTP/1.1\" 200 3574 \"-\" \"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6\"",
			},
			want:  map[string]int{},
			want1: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup
			cfg := config.Values
			urlCountChan := make(chan map[string]int)
			ipCountChan := make(chan map[string]int)
			logChan := make(chan string, len(tt.args))

			// execute
			go CountLogMatch(cfg.Regex.MatchIPsURLsIgnoreQuery, logChan, urlCountChan, ipCountChan)

			// create a loop
			for _, log := range tt.args {
				logChan <- log
			}
			// close channel
			close(logChan)

			// receive URL/IP counts
			got := <-urlCountChan
			got1 := <-ipCountChan

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountLogMatchesNoQuery() got = %v, expected %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CountLogMatchesNoQuery() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
