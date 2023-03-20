package zlog

import (
	"unsafe"

	"github.com/rs/zerolog"
)

// Event is a log event which as same as zerolog.Event
type Event struct {
	buf       []byte
	w         zerolog.LevelWriter
	level     zerolog.Level
	done      func(msg string)
	stack     bool           // enable error stack trace
	ch        []zerolog.Hook // hooks from context
	skipFrame int            // The number of additional frames to skip when printing the caller.
}

// SetLevelWriter set the level writer
func (e *Event) SetLevelWriter(levelWriter zerolog.LevelWriter) {
	e.w = levelWriter
}

// NewEvent covert zerolog.Event to log.Event with unsafe
func NewEvent(e *zerolog.Event) *Event {
	// covert zerolog.Event to log.Event with unsafe
	return (*Event)(unsafe.Pointer(e))
}
