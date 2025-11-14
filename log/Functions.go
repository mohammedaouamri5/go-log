package log

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func newColors() *Colors {
	return &Colors{
		Info: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00C7FF")).Bold(true),

		Warn: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD700")).Bold(true).Italic(true),

		Error: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF4500")).Bold(true),

		Fatal: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Background(lipgloss.Color("#2E0000")).
			Bold(true).
			Underline(true),

		Time: lipgloss.NewStyle().Foreground(lipgloss.Color("#999999")).Italic(true),
		File: lipgloss.NewStyle().Foreground(lipgloss.Color("#87CEEB")).Bold(true), // skyblue
		Line: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF69B4")).Bold(true), // pink

		Func:      lipgloss.NewStyle().Foreground(lipgloss.Color("#ADFF2F")).Italic(true).Bold(true), // green-yellow
		KeyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#7D56F4")).Bold(true),
		TypeStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#999999")).Italic(true),
		StrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575")),
		NumStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#FF8800")),
		BoolStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF4444")).Bold(true),
		NullStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#888888")).Italic(true),
		PtrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#00BFFF")).Italic(true),
		Bracket:   lipgloss.NewStyle().Foreground(lipgloss.Color("#AAAAAA")),
	}
}

var indentSize = 2

func prettyPrint(colors *Colors, v any, indent int) string {
	if v == nil {
		return colors.NullStyle.Render("null")
	}

	// --- Handle error type first ---
	if err, ok := v.(error); ok {
		ptr := fmt.Sprintf("%p", err) // memory address like 0x97869
		str := fmt.Sprintf("%q", err.Error())
		return colors.Error.Render(fmt.Sprintf("%s | %s", ptr, str))
	}

	val := reflect.ValueOf(v)

	switch val.Kind() {
	case reflect.Map:
		var sb strings.Builder
		sb.WriteString(colors.Bracket.Render("{") + "\n")
		for _, key := range val.MapKeys() {
			k := key.Interface()
			sb.WriteString(strings.Repeat(" ", indent+indentSize))
			sb.WriteString(colors.KeyStyle.Render(fmt.Sprintf("%q", k)))

			fieldVal := val.MapIndex(key).Interface()

			// --- Detect if it's an error to color the type as (error) ---
			if _, ok := fieldVal.(error); ok {
				sb.WriteString(" " + colors.TypeStyle.Render("(error)"))
			} else {
				sb.WriteString(" " + colors.TypeStyle.Render(fmt.Sprintf("(%s)", reflect.TypeOf(fieldVal).Kind())))
			}

			sb.WriteString(": ")
			sb.WriteString(prettyPrint(colors, fieldVal, indent+indentSize))
			sb.WriteString("\n")
		}
		sb.WriteString(strings.Repeat(" ", indent) + colors.Bracket.Render("}"))
		return sb.String()

	case reflect.Slice, reflect.Array:
		// Special case for []byte
		if val.Type().Elem().Kind() == reflect.Uint8 {
			return colors.StrStyle.Render(fmt.Sprintf("<bytes: len=%d>", val.Len()))
		}

		var sb strings.Builder
		sb.WriteString(colors.Bracket.Render("[") + "\n")
		for i := 0; i < val.Len(); i++ {
			sb.WriteString(strings.Repeat(" ", indent+indentSize))
			elem := val.Index(i).Interface()
			sb.WriteString(prettyPrint(colors, elem, indent+indentSize))
			sb.WriteString("\n")
		}
		sb.WriteString(strings.Repeat(" ", indent) + colors.Bracket.Render("]"))
		return sb.String()

	case reflect.String:
		return colors.StrStyle.Render(fmt.Sprintf("%q", val.String()))
	case reflect.Int, reflect.Int64, reflect.Float64:
		return colors.NumStyle.Render(fmt.Sprint(val.Interface()))
	case reflect.Bool:
		return colors.BoolStyle.Render(fmt.Sprint(val.Bool()))
	default:
		// fallback: try JSON encoding
		j, _ := json.Marshal(val.Interface())
		return colors.StrStyle.Render(string(j))
	}
}

func GetSimpelFormatterText(__color *Colors) Formatter {
	if __color == nil {
		__color = newColors()
	}

	basePath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err.Error())
		return nil
	}

	for {
		if _, err := os.Stat(filepath.Join(basePath, "go.mod")); err == nil {
			break
		}

		parent := filepath.Dir(basePath)
		if parent == basePath {
			fmt.Println("Error:", err.Error())
			return nil
		}
		basePath = parent
	}

	__func :=
		func(level string, fields map[string]any, msg string, back int, args ...any) string {
			// Caller info
			pc := make([]uintptr, 1)
			runtime.Callers(back, pc)
			frame, _ := runtime.CallersFrames(pc).Next()

			// Relative file path
			file := frame.File
			if basePath != "" {
				if rel, err := filepath.Rel(basePath, file); err == nil {
					file = rel
				}
			}

			// Extract parts
			timestamp := time.Now().Format("03:04:05 PM")
			funcName := filepath.Base(frame.Function)
			line := fmt.Sprintf("%d", frame.Line)

			// Styled parts

			timeStr := __color.Time.Render("[" + timestamp + "]")
			levelStr := __color.TheColor(level).Render(strings.ToUpper(level))
			funcStr := __color.Func.Render(funcName)
			fileStr := __color.File.Render(file)
			lineStr := __color.Line.Render(line)

			// Final log
			if fields == nil {
				return fmt.Sprintf("%s  %s (%s | %s:%s)\n\t\t%s\n",
					levelStr,
					timeStr,
					funcStr,
					fileStr,
					lineStr,
					fmt.Sprintf(msg, args...),
				)
			} else {
				return fmt.Sprintf("%s  %s (%s | %s:%s)\n\t\t%s\n%s\n",
					levelStr,
					timeStr,
					funcStr,
					fileStr,
					lineStr,
					fmt.Sprintf(msg, args...),
					prettyPrint(__color, fields, 0),
				)
			}
		}

	return __func
}
