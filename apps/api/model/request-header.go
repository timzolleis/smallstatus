package model

type RequestHeader struct {
	Base
	Key       string
	Value     string
	MonitorID uint
}
