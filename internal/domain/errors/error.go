package errors

// ReportError type
type ReportError string

func (ee ReportError) Error() string {
	return string(ee)
}

var (
	ErrOverlaping        = ReportError("another event exists for this date")
	ErrReportExist       = ReportError("event alredy exist")
	ErrReportNotFound    = ReportError("not found event")
	ErrWrangParams       = ReportError("invalid input params")
	ErrConfigWrangParams = ReportError("can not validate config file")
)
