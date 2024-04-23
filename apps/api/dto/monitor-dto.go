package dto

type MonitorDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	Type     string `json:"type"`
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
	Type     string                   `json:"type"`
	Headers  []CreateMonitorHeaderDTO `json:"headers"`
}

type CreateMonitorHeaderDTO struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
