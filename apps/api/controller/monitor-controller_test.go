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
	monitorType := "http"
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
				Type:        monitorType,
				Interval:    interval,
				WorkspaceID: workspace,
			},
			args: args{
				dto: &dto.MonitorDTO{
					Name:     name,
					Url:      url,
					Interval: interval,
					Type:     monitorType},
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
	monitorType := "http"

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
					Name:     name,
					Url:      url,
					Type:     monitorType,
					Interval: interval,
				},
			},
			want: dto.MonitorDTO{
				Name:     name,
				Url:      url,
				Type:     monitorType,
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