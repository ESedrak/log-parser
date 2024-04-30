package ip_mgmt

import (
	"reflect"
	"testing"
)

func TestIPCounter_MostActiveIP(t *testing.T) {
	type fields struct {
		IPCounts map[string]int
	}
	type args struct {
		requestedCount int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []IPCount
		wantErr bool
	}{
		{
			name: "Success: return top 3 most active IPs",
			fields: fields{
				IPCounts: map[string]int{
					"172.41.191.9": 5,
					"169.41.191.9": 2,
					"170.41.191.9": 3,
					"171.41.191.9": 4,
					"168.41.191.9": 1,
				},
			},
			args: args{
				requestedCount: 3,
			},
			want:    []IPCount{{IP: "172.41.191.9", Count: 5}, {IP: "171.41.191.9", Count: 4}, {IP: "170.41.191.9", Count: 3}},
			wantErr: false,
		},
		{
			name: "Error: requestCount less than one",
			fields: fields{
				IPCounts: map[string]int{},
			},
			args: args{
				requestedCount: 0,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IPCounter{
				IPCounts: tt.fields.IPCounts,
			}
			got, err := i.MostActiveIP(tt.args.requestedCount)
			if (err != nil) != tt.wantErr {
				t.Errorf("IPCounter.MostActiveIPs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IPCounter.MostActiveIPs() = %v, want %v", got, tt.want)
			}
		})
	}
}
