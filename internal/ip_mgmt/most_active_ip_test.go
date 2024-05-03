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
		// TODO: Add test cases.
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
