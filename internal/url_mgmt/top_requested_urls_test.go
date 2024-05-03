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
		{
			name: "Success: return top 3 requested URLs",
			args: args{
				urlCounts: map[string]int{
					"/page4": 3,
					"/page1": 10,
					"/page2": 8,
					"/page3": 5,
					"/page5": 1,
				},
				requestedNum: 3,
			},
			want: []URLCount{
				{URL: "/page1", Count: 10},
				{URL: "/page2", Count: 8},
				{URL: "/page3", Count: 5},
			},
			wantErr: false,
		},
		{
			name: "Error: requestCount less than one",
			args: args{
				urlCounts: map[string]int{
					"/page4": 3,
					"/page1": 10,
					"/page2": 8,
					"/page3": 5,
					"/page5": 1,
				},
				requestedNum: 0,
			},
			want:    nil,
			wantErr: true,
		},
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
