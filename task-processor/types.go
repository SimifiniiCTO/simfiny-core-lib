package taskprocessor

// ProcessingInterval represents the various processing intervals for a recurring task
type ProcessingInterval string

const (
	EveryHour      ProcessingInterval = "@every 1h"
	Every3Hours    ProcessingInterval = "@every 3h"
	Every6Hours    ProcessingInterval = "@every 6h"
	Every12Hours   ProcessingInterval = "@every 12h"
	Every24Hours   ProcessingInterval = "@every 24h"
	Every30Minutes ProcessingInterval = "@every 30m"
	Every15Minutes ProcessingInterval = "@every 15m"
)

// The `String()` method is a method defined on the `ProcessingInterval` type. It is used to convert a
// `ProcessingInterval` value to its corresponding string representation.
func (p ProcessingInterval) String() string {
	return string(p)
}
