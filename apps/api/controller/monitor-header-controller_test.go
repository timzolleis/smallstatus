package controller

import (
	"github.com/timzolleis/smallstatus/dto"
	"github.com/timzolleis/smallstatus/model"
	"reflect"
	"testing"
)

func Test_mapMonitorHeaderDTO(t *testing.T) {

	header := &model.MonitorHeader{
		MonitorID: 1,
		Key:       "my-header",
		Value:     "my-value",
	}

	want := dto.MonitorHeaderDTO{
		ID:    header.ID,
		Key:   header.Key,
		Value: header.Value,
	}

	type args struct {
		header *model.MonitorHeader
	}
	tests := []struct {
		name string
		args args
		want dto.MonitorHeaderDTO
	}{
		{
			name: "Valid Header",
			args: args{
				header: header,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapMonitorHeaderDTO(tt.args.header); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapMonitorHeaderDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapMonitorHeaderDTOToModel(t *testing.T) {

	headerDto := dto.MonitorHeaderDTO{
		ID:    1,
		Key:   "my-header",
		Value: "my-value",
	}

	want := model.MonitorHeader{
		Key:   headerDto.Key,
		Value: headerDto.Value,
	}

	type args struct {
		dto       *dto.MonitorHeaderDTO
		monitorId uint
	}
	tests := []struct {
		name string
		args args
		want model.MonitorHeader
	}{
		{
			name: "Valid DTO",
			want: want,
			args: args{
				dto: &headerDto,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapMonitorHeaderDTOToModel(tt.args.dto, tt.args.monitorId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapMonitorHeaderDTOToModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
