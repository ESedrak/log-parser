package ip_mgmt

import (
	"reflect"
	"testing"
)

func TestUniqueIPs(t *testing.T) {
	type args struct {
		ipCounts map[string]int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Success: return the number of unique IPs",
			args: args{
				ipCounts: map[string]int{
					"172.41.191.9": 5,
					"169.41.191.9": 2,
					"170.41.191.9": 3,
					"171.41.191.9": 4,
					"168.41.191.9": 1,
				},
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "Error: no IP's found",
			args: args{
				ipCounts: map[string]int{},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UniqueIPs(tt.args.ipCounts)
			if (err != nil) != tt.wantErr {
				t.Errorf("UniqueIPs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueIPs() = %v, want %v", got, tt.want)
			}
		})
	}
}
