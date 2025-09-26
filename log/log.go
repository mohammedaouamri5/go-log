package log

import (
	"encoding/json"
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
	fields    map[string]any
	Writer    io.Writer
	Formatter func(Level string, fields map[string]any, __str string, args ...any) string
}

func (l *Logger) WithObj(obj any) *Logger {
	data, err := json.Marshal(obj)
	if err != nil {
		return l
	}
	err = json.Unmarshal(data, &l.fields)
	print(err.Error())
	return l
}
func (l *Logger) WithMap(obj map[string]any) *Logger {
	l.fields = obj
	return l
}

func (l *Logger) Info(__str string, args ...any) {
	__format := l.Formatter("INFO", l.fields, __str, args...)
	l.Writer.Write([]byte(__format))
	l.fields = nil
}
func (l *Logger) InfoDebug(__str string, args ...any) {
	__format := l.Formatter("INFO", l.fields, __str, args...)
	l.Writer.Write([]byte(__format))
	l.fields = nil
}
func (l *Logger) Warn(__str string, args ...any) {
	__format := l.Formatter("WARN", l.fields, __str, args...)
	l.Writer.Write([]byte(__format))
	l.fields = nil
}
func (l *Logger) WarnD(__str string, args ...any) {
	__format := l.Formatter("WARN", l.fields, __str, args...)
	l.Writer.Write([]byte(__format))
	l.fields = nil
}
func (l *Logger) Error(__str string, args ...any) {
	__format := l.Formatter("ERROR", l.fields, __str, args...)
	l.Writer.Write([]byte(__format))
	l.fields = nil
}
func (l *Logger) ErrorD(__str string, args ...any) {
	__format := l.Formatter("ERROR", l.fields, __str, args...)
	l.Writer.Write([]byte(__format))
	l.fields = nil
}
func (l *Logger) Fatal(__str string, args ...any) {
	__format := l.Formatter("FATAL", l.fields, __str, args...)
	l.Writer.Write([]byte(__format))
	l.fields = nil
}
func (l *Logger) FatalD(__str string, args ...any) {
	__format := l.Formatter("FATAL", l.fields, __str, args...)
	l.Writer.Write([]byte(__format))
	l.fields = nil
}
