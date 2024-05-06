package model

type Response struct {
	Base
	MonitorID  uint
	StatusCode int
	Body       string
	Headers    []ResponseHeader `gorm:"foreignKey:ResponseID"`
	Duration   int
	Retries    int
}
