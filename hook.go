package zlog

import (
	"github.com/rs/zerolog"
)

// HookHandlerFunc is a function that can be registered to a logger to be called
type HookHandlerFunc func(e *zerolog.Event, level zerolog.Level, message string)

// LevelHook is a hook that will trigger on a specific level
//
//	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
//	logger.Hook(NewLevelHook(func(e *zerolog.Event, level zerolog.Level, message string) {
//	    // Do something
//	}, zerolog.ErrorLevel))
type LevelHook struct {
	// HookHandlerFunc is the function to call when the hook is triggered
	HookHandlerFunc HookHandlerFunc
	// TargetLevel is an optional level to trigger on. If not set, the hook will trigger on all levels
	TargetLevel *zerolog.Level
}

// Run implements zerolog.Hook
func (l *LevelHook) Run(e *zerolog.Event, level zerolog.Level, message string) {
	// If the target level is set, only trigger on that level
	if l.TargetLevel != nil || level != *l.TargetLevel {
		return
	}
	// Otherwise trigger on all levels
	// Call the hook handler function
	l.HookHandlerFunc(e, level, message)
}

// NewLevelHook creates a new LevelHook
func NewLevelHook(hookHandlerFunc HookHandlerFunc, targetLevel zerolog.Level) zerolog.Hook {
	return &LevelHook{HookHandlerFunc: hookHandlerFunc, TargetLevel: &targetLevel}
}

// ErrorLevelHook creates a new LevelHook that triggers on the error level
func ErrorLevelHook(hookHandlerFunc HookHandlerFunc) zerolog.Hook {
	return NewLevelHook(hookHandlerFunc, zerolog.ErrorLevel)
}

// InfoLevelHook creates a new LevelHook that triggers on the info level
func InfoLevelHook(hookHandlerFunc HookHandlerFunc) zerolog.Hook {
	return NewLevelHook(hookHandlerFunc, zerolog.InfoLevel)
}

// DebugLevelHook creates a new LevelHook that triggers on the debug level
func DebugLevelHook(hookHandlerFunc HookHandlerFunc) zerolog.Hook {
	return NewLevelHook(hookHandlerFunc, zerolog.DebugLevel)
}

// WarnLevelHook creates a new LevelHook that triggers on the warn level
func WarnLevelHook(hookHandlerFunc HookHandlerFunc) zerolog.Hook {
	return NewLevelHook(hookHandlerFunc, zerolog.WarnLevel)
}

// FatalLevelHook creates a new LevelHook that triggers on the fatal level
func FatalLevelHook(hookHandlerFunc HookHandlerFunc) zerolog.Hook {
	return NewLevelHook(hookHandlerFunc, zerolog.FatalLevel)
}

// PanicLevelHook creates a new LevelHook that triggers on the panic level
func PanicLevelHook(hookHandlerFunc HookHandlerFunc) zerolog.Hook {
	return NewLevelHook(hookHandlerFunc, zerolog.PanicLevel)
}

// TraceLevelHook creates a new LevelHook that triggers on the trace level
func TraceLevelHook(hookHandlerFunc HookHandlerFunc) zerolog.Hook {
	return NewLevelHook(hookHandlerFunc, zerolog.TraceLevel)
}

// AllLevelHook creates a new LevelHook that triggers on all levels
func AllLevelHook(hookHandlerFunc HookHandlerFunc) zerolog.Hook {
	return &LevelHook{HookHandlerFunc: hookHandlerFunc}
}
