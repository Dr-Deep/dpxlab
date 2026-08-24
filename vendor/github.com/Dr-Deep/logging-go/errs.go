package logging

import (
	"fmt"
)

// FlushError represents an error that occurred during a flush.
type FlushError struct {
	Err error
}

func (e *FlushError) Error() string {
	return fmt.Sprintf("Logger.Writer flush failed: %v", e.Err)
}

// CloseError represents an error that occurred during file close.
type CloseError struct {
	Err error
}

func (e *CloseError) Error() string {
	return fmt.Sprintf("Logger.File close failed: %v", e.Err)
}
