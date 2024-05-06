package model

type Monitor struct {
	Base
	Name        string
	Url         string
	Interval    int
	Timeout     int
	Retries     int
	Method      string
	WorkspaceID uint
	Headers     []RequestHeader `gorm:"foreignKey:MonitorID"`
}
