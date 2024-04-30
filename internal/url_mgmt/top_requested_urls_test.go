package url_mgmt

import (
	"reflect"
	"testing"
)

func TestURLCounter_TopRequestedURLs(t *testing.T) {
	type fields struct {
		URLCounts map[string]int
	}
	type args struct {
		requestedCount int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []URLCount
		wantErr bool
	}{
		{
			name: "Success: return top 3 requested URLs",
			fields: fields{
				URLCounts: map[string]int{
					"/page4": 3,
					"/page1": 10,
					"/page2": 8,
					"/page3": 5,
					"/page5": 1,
				},
			},
			args: args{
				requestedCount: 3,
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
			fields: fields{
				URLCounts: map[string]int{
					"/page1": 10,
					"/page2": 8,
					"/page3": 5,
				},
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
			u := &URLCounter{
				URLCounts: tt.fields.URLCounts,
			}
			got, err := u.TopRequestedURLs(tt.args.requestedCount)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLCounter.TopRequestedURLs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URLCounter.TopRequestedURLs() = %v, want %v", got, tt.want)
			}
		})
	}
}
