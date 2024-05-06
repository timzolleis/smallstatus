package helper

import (
	"reflect"
	"testing"
)

func TestNewErrorResponse(t *testing.T) {
	type args struct {
		message string
		code    int
	}
	tests := []struct {
		name string
		args args
		want *ErrorResponse
	}{
		{
			name: "Test NewErrorResponse",
			args: args{
				message: "error",
				code:    400,
			},
			want: &ErrorResponse{
				Code:    400,
				Message: "error",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewErrorResponse(tt.args.message, tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewErrorResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSuccessResponse(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *SuccessResponse
	}{
		{
			name: "Test NewSuccessResponse",
			args: args{message: "success"},
			want: &SuccessResponse{
				Message: "success",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSuccessResponse(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSuccessResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
