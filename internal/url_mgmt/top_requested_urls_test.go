package url_mgmt

import (
	"reflect"
	"testing"
)

func TestTopRequestedURLs(t *testing.T) {
	type args struct {
		urlCounts    map[string]int
		requestedNum int
	}
	tests := []struct {
		name    string
		args    args
		want    []URLCount
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TopRequestedURLs(tt.args.urlCounts, tt.args.requestedNum)
			if (err != nil) != tt.wantErr {
				t.Errorf("TopRequestedURLs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopRequestedURLs() = %v, want %v", got, tt.want)
			}
		})
	}
}
