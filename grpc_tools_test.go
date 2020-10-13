package go_tools

import (
	"golang.org/x/net/context"
	"testing"
)

func TestGetClientIP(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetClientIP(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClientIP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetClientIP() got = %v, want %v", got, tt.want)
			}
		})
	}
}
