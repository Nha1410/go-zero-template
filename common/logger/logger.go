package logger

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

// Logger wraps go-zero logger with additional context
type Logger struct {
	logx.Logger
}

// NewLogger creates a new logger instance
func NewLogger() *Logger {
	return &Logger{
		Logger: logx.WithContext(context.Background()),
	}
}

// WithContext returns a logger with context
func (l *Logger) WithContext(ctx context.Context) *Logger {
	return &Logger{
		Logger: logx.WithContext(ctx),
	}
}

// WithFields returns a logger with fields
func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	ctx := context.WithValue(context.Background(), "fields", fields)
	return &Logger{
		Logger: logx.WithContext(ctx),
	}
}

// LogRequest logs HTTP request
func (l *Logger) LogRequest(method, path string, statusCode int, duration int64) {
	l.Infov(map[string]interface{}{
		"method":   method,
		"path":     path,
		"status":   statusCode,
		"duration": duration,
	})
}

// LogError logs error with context
func (l *Logger) LogError(err error, fields ...logx.LogField) {
	if err != nil {
		// Combine error field with additional fields
		allFields := make([]logx.LogField, 0, len(fields)+1)
		allFields = append(allFields, logx.Field("error", err.Error()))
		allFields = append(allFields, fields...)
		// Error method accepts variadic any arguments
		args := make([]interface{}, 0, len(allFields)+1)
		args = append(args, "Error occurred")
		for _, f := range allFields {
			args = append(args, f)
		}
		l.Error(args...)
	}
}

