package log

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/charmbracelet/lipgloss"
)

type Colors struct {
	Info, Warn, Error, Fatal                                                         lipgloss.Style
	Time, File, Line, Func                                                           lipgloss.Style
	KeyStyle, TypeStyle, StrStyle, NumStyle, BoolStyle, NullStyle, PtrStyle, Bracket lipgloss.Style
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

type Formatter = func(Level string, fields map[string]any, __str string, back int, args ...any) string

type Logger struct {
	Writer    io.Writer
	Formatter Formatter
	back      int

	mu   sync.Mutex // protect Writer.Write
	pool *sync.Pool // reuse buffers
}

type LoggerWithFields struct {
	*Logger
	fields map[string]any
}

// initialize a new logger
func NewLogger(w io.Writer, formatter Formatter, back int) *Logger {
	return &Logger{
		Writer:    w,
		Formatter: formatter,
		back:      back,
		pool: &sync.Pool{
			New: func() any {
				b := make([]byte, 0, 1024)
				return &b
			},
		},
	}
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

	__result.fields = __map
	return __result
}

func (l *Logger) WithMap(obj map[string]any) *LoggerWithFields {
	return &LoggerWithFields{
		Logger: l,
		fields: obj,
	}
}

func (l *Logger) WithErr(err error) *LoggerWithFields {

	result := &LoggerWithFields{
		Logger: l,
		fields: make(map[string]any),
	}
	if err == nil {
		result.fields["error"] = nil
	} else {
		result.fields["error"] = err
	}
	return result
}

// --- the key function you asked to modify ---
func (l *Logger) write(__data []byte) {
	l.mu.Lock()
	defer l.mu.Unlock()

	bufPtr := l.pool.Get().(*[]byte)
	buf := *bufPtr
	buf = append(buf[:0], __data...)

	l.Writer.Write(buf)

	// reset and return to pool
	*bufPtr = buf[:0]
	l.pool.Put(bufPtr)
}

// --- log levels ---

func (l *Logger) Info(__str string, args ...any) {
	__format := l.Formatter("INFO", nil, __str, l.back, args...)
	l.write([]byte(__format))
}

func (l *Logger) Warn(__str string, args ...any) {
	__format := l.Formatter("WARN", nil, __str, l.back, args...)
	l.write([]byte(__format))
}

func (l *Logger) Error(__str string, args ...any) {
	__format := l.Formatter("ERROR", nil, __str, l.back, args...)
	l.write([]byte(__format))
}

func (l *Logger) Fatal(__str string, args ...any) {
	__format := l.Formatter("FATAL", nil, __str, l.back, args...)
	l.write([]byte(__format))
}

// --- WithFields versions ---
func (l *LoggerWithFields) Info(__str string, args ...any) {
	__format := l.Logger.Formatter("INFO", l.fields, __str, l.Logger.back-1, args...)
	l.write([]byte(__format))
}
func (l *LoggerWithFields) Warn(__str string, args ...any) {
	__format := l.Logger.Formatter("WARN", l.fields, __str, l.Logger.back-1, args...)
	l.write([]byte(__format))
}
func (l *LoggerWithFields) Error(__str string, args ...any) {
	__format := l.Logger.Formatter("ERROR", l.fields, __str, l.Logger.back-1, args...)
	l.write([]byte(__format))
}
func (l *LoggerWithFields) Fatal(__str string, args ...any) {
	__format := l.Logger.Formatter("FATAL", l.fields, __str, l.Logger.back-1, args...)
	l.write([]byte(__format))
}
