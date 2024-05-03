package ip_mgmt

import (
	"reflect"
	"testing"
)

func TestMostActiveIP(t *testing.T) {
	type args struct {
		ipCounts     map[string]int
		requestedNum int
	}
	tests := []struct {
		name    string
		args    args
		want    []IPCount
		wantErr bool
	}{
		{
			name: "Success: return top 3 most active IPs",
			args: args{
				ipCounts: map[string]int{
					"172.41.191.9": 5,
					"169.41.191.9": 2,
					"170.41.191.9": 3,
					"171.41.191.9": 4,
					"168.41.191.9": 1,
				},
				requestedNum: 3,
			},
			want:    []IPCount{{IP: "172.41.191.9", Count: 5}, {IP: "171.41.191.9", Count: 4}, {IP: "170.41.191.9", Count: 3}},
			wantErr: false,
		},
		{
			name: "Error: requestCount less than one",
			args: args{
				ipCounts: map[string]int{
					"172.41.191.9": 5,
					"169.41.191.9": 2,
					"170.41.191.9": 3,
					"171.41.191.9": 4,
					"168.41.191.9": 1,
				},
				requestedNum: 0,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MostActiveIP(tt.args.ipCounts, tt.args.requestedNum)
			if (err != nil) != tt.wantErr {
				t.Errorf("MostActiveIP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MostActiveIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
