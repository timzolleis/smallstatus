package dto

type MonitorDto struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	Type     string `json:"type"`
	Interval int    `json:"interval"`
}
