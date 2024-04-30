package controller

import (
	"github.com/timzolleis/smallstatus/dto"
	"github.com/timzolleis/smallstatus/model"
	"reflect"
	"testing"
)

func Test_mapMonitorDTOToModel(t *testing.T) {
	name := "MyMonitor"
	url := "https://example.com"
	interval := 30
	method := "POST"
	retries := 5
	monitorTimeout := 10
	workspace := uint(1)

	type args struct {
		dto       *dto.MonitorDTO
		workspace uint
	}
	tests := []struct {
		name string
		args args
		want model.Monitor
	}{
		{
			name: "Valid DTO",
			want: model.Monitor{
				Name:        name,
				Url:         url,
				Method:      method,
				Interval:    interval,
				Retries:     retries,
				Timeout:     monitorTimeout,
				WorkspaceID: workspace,
			},
			args: args{
				dto: &dto.MonitorDTO{
					Name:     name,
					Url:      url,
					Interval: interval,
					Method:   method,
					Retries:  retries,
					Timeout:  monitorTimeout},
				workspace: workspace,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapMonitorDTOToModel(tt.args.dto, tt.args.workspace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapMonitorDTOToModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapMonitorToDTO(t *testing.T) {
	name := "MyMonitor"
	url := "https://example.com"
	interval := 30
	method := "POST"
	retries := 5
	monitorTimeout := 10
	workspace := uint(1)

	type args struct {
		monitor *model.Monitor
	}
	tests := []struct {
		name string
		args args
		want dto.MonitorDTO
	}{
		{
			name: "Valid Model",
			args: args{
				monitor: &model.Monitor{
					Name:        name,
					Url:         url,
					Method:      method,
					Retries:     retries,
					Timeout:     monitorTimeout,
					WorkspaceID: workspace,
					Interval:    interval,
				},
			},
			want: dto.MonitorDTO{
				Name:     name,
				Url:      url,
				Method:   method,
				Retries:  retries,
				Timeout:  monitorTimeout,
				Interval: interval,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapMonitorToDTO(tt.args.monitor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapMonitorToDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}
