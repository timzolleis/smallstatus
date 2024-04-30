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
	Name     string                   `json:"name"`
	Url      string                   `json:"url"`
	Interval int                      `json:"interval"`
	Method   string                   `json:"method"`
	Timeout  int                      `json:"timeout"`
	Retries  int                      `json:"retries"`
	Headers  []CreateMonitorHeaderDTO `json:"headers"`
}

type CreateMonitorHeaderDTO struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
