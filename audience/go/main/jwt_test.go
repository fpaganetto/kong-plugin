package main

import (
	"reflect"
	"testing"
)

func TestGetToken(t *testing.T) {
	type args struct {
		jwt string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestGetToken",
			args: args{
				jwt: "Bearer <token>",
			},
			want: "<token>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetToken(tt.args.jwt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestGetToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAudience(t *testing.T) {
	type args struct {
		jwt string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestGetAudience",
			args: args{
				jwt: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJhdWQiOlsidGVzdGluZyJdfQ.c-XYfRuLIRdH0WwDbcY76T9tvGw1onwr0V1HyZUX_qI",
			},
			want: []string{"testing"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAudience(tt.args.jwt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAudience() = %v, want %v", got, tt.want)
			}
		})
	}
}
