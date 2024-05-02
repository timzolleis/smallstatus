package dto

type MonitorDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	Method   string `json:"method"`
	Timeout  int    `json:"timeout"`
	Retries  int    `json:"retries"`
	Interval int    `json:"interval"`
}

type MonitorHeaderDTO struct {
	ID    uint   `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CreateMonitorDTO struct {
	Name     string                   `json:"name" validate:"required"`
	Url      string                   `json:"url" validate:"required,http_url"`
	Interval int                      `json:"interval" validate:"required,numeric,min=1,max=180"`
	Method   string                   `json:"method" validate:"required,http_method"`
	Timeout  int                      `json:"timeout" validate:"required,numeric,min=1,max=60"`
	Retries  int                      `json:"retries" validate:"required,numeric,min=1,max=10"`
	Headers  []CreateMonitorHeaderDTO `json:"headers"`
}

type CreateMonitorHeaderDTO struct {
	Key   string `json:"key" ,validate:"required"`
	Value string `json:"value" ,validate:"required"`
}
