package log_mgmt

import "testing"

func TestCountLogMatchesIgnoresQuery(t *testing.T) {
	type args struct {
		logChan      <-chan string
		urlCountChan chan<- map[string]int
		ipCountChan  chan<- map[string]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CountLogMatchesIgnoresQuery(tt.args.logChan, tt.args.urlCountChan, tt.args.ipCountChan)
		})
	}
}
