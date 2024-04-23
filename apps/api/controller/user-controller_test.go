package controller

import (
	"github.com/timzolleis/smallstatus/dto"
	"github.com/timzolleis/smallstatus/model"
	"reflect"
	"testing"
)

func Test_mapToDTO(t *testing.T) {
	user := &model.User{
		Name:     "MyUser",
		Email:    "test@me.com",
		Password: "MyPassword",
	}
	want := dto.UserDTO{
		Name:  user.Name,
		Email: user.Email,
	}

	type args struct {
		user *model.User
	}
	tests := []struct {
		name string
		args args
		want dto.UserDTO
	}{
		{
			args: args{
				user: user,
			},
			want: want,
			name: "Valid User",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapToDTO(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapToDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}
