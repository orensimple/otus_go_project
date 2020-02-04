package models

import "time"

// Report main model
type Report struct {
	BannerID   int64
	GroupID    int64
	SlotID     int64
	Show       int64
	Conversion int64
}

type ReportQueueFormat struct {
	Type     int64
	BannerID int64
	GroupID  int64
	SlotID   int64
	Time     time.Time
}
