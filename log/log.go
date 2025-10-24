package log

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/charmbracelet/lipgloss"
)

type Colors struct {
	Info  lipgloss.Style
	Warn  lipgloss.Style
	Error lipgloss.Style
	Fatal lipgloss.Style

	Time lipgloss.Style
	File lipgloss.Style
	Line lipgloss.Style
	Func lipgloss.Style

	KeyStyle  lipgloss.Style
	TypeStyle lipgloss.Style
	StrStyle  lipgloss.Style
	NumStyle  lipgloss.Style
	BoolStyle lipgloss.Style
	NullStyle lipgloss.Style
	PtrStyle  lipgloss.Style
	Bracket   lipgloss.Style
}

func (c *Colors) TheColor(Level string) *lipgloss.Style {
	switch Level {
	case "INFO":
		return &c.Info
	case "WARN":
		return &c.Warn
	case "ERROR":
		return &c.Error
	case "FATAL":
		return &c.Fatal
	default:
		return nil
	}
}

type Logger struct {
	Writer    io.Writer
	Formatter func(Level string, fields map[string]any, __str string, args ...any) string
}

type LoggerWithFields struct {
	*Logger
	fields map[string]any
}

func (l *Logger) WithObj(__obj any) *LoggerWithFields {
	__result := &LoggerWithFields{
		Logger: l,
		fields: nil,
	}

	if __obj == nil {
		return __result
	}

	__data, __err := json.Marshal(__obj)
	if __err != nil {
		fmt.Printf("WithObj: failed to marshal object: %v\n", __err)
		return __result
	}

	var __map map[string]any
	__err = json.Unmarshal(__data, &__map)
	if __err != nil {
		fmt.Printf("WithObj: failed to unmarshal into map: %v\n", __err)
		return __result
	}
	fmt.Printf("WithObj: map: %v\n", __map)
	fmt.Printf("WithObj: obj: %v\n", __obj)

	__result.fields = __map
	return __result
}



func (l *Logger) WithMap(obj map[string]any) *LoggerWithFields {
	__result := LoggerWithFields{
		Logger: l,
		fields: obj,
	}
	return &__result
}

func (l *Logger) Info(__str string, args ...any) {
	__format := l.Formatter("INFO", nil, __str, args...)
	l.Writer.Write([]byte(__format))
}
func (l *Logger) InfoDebug(__str string, args ...any) {
	__format := l.Formatter("INFO", nil, __str, args...)
	l.Writer.Write([]byte(__format))
}
func (l *Logger) Warn(__str string, args ...any) {
	__format := l.Formatter("WARN", nil, __str, args...)
	l.Writer.Write([]byte(__format))
}
func (l *Logger) WarnD(__str string, args ...any) {
	__format := l.Formatter("WARN", nil, __str, args...)
	l.Writer.Write([]byte(__format))
}
func (l *Logger) Error(__str string, args ...any) {
	__format := l.Formatter("ERROR", nil, __str, args...)
	l.Writer.Write([]byte(__format))
}
func (l *Logger) ErrorD(__str string, args ...any) {
	__format := l.Formatter("ERROR", nil, __str, args...)
	l.Writer.Write([]byte(__format))
}
func (l *Logger) Fatal(__str string, args ...any) {
	__format := l.Formatter("FATAL", nil, __str, args...)
	l.Writer.Write([]byte(__format))
}
func (l *Logger) FatalD(__str string, args ...any) {
	__format := l.Formatter("FATAL", nil, __str, args...)
	l.Writer.Write([]byte(__format))
}

// -------------
// -------------
// -------------

func (l *LoggerWithFields) Info(__str string, args ...any) {
	__format := l.Logger.Formatter("INFO", l.fields, __str, args...)
	l.Writer.Write([]byte(__format))
}
func (l *LoggerWithFields) InfoDebug(__str string, args ...any) {
	__format := l.Logger.Formatter("INFO", l.fields, __str, args...)
	l.Writer.Write([]byte(__format))
}
func (l *LoggerWithFields) Warn(__str string, args ...any) {
	__format := l.Formatter("WARN", l.fields, __str, args...)
	l.Writer.Write([]byte(__format))
}
func (l *LoggerWithFields) WarnD(__str string, args ...any) {
	__format := l.Logger.Formatter("WARN", l.fields, __str, args...)
	l.Writer.Write([]byte(__format))
}
func (l *LoggerWithFields) Error(__str string, args ...any) {
	__format := l.Logger.Formatter("ERROR", l.fields, __str, args...)
	l.Writer.Write([]byte(__format))
}
func (l *LoggerWithFields) ErrorD(__str string, args ...any) {
	__format := l.Logger.Formatter("ERROR", l.fields, __str, args...)
	l.Writer.Write([]byte(__format))
}
func (l *LoggerWithFields) Fatal(__str string, args ...any) {
	__format := l.Logger.Formatter("FATAL", l.fields, __str, args...)
	l.Writer.Write([]byte(__format))
}
func (l *LoggerWithFields) FatalD(__str string, args ...any) {
	__format := l.Logger.Formatter("FATAL", l.fields, __str, args...)
	l.Writer.Write([]byte(__format))
}
