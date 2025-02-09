// Package easy allows to easily format output of Logrus logger
package easy

import (
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	// Default log format will output [INFO]: 2006-01-02T15:04:05Z07:00 - Log message
	defaultLogFormat       = "[%lvl%]: %time% - %msg% %fields%"
	defaultTimestampFormat = time.RFC3339
	defaultFieldFormat     = "%k%: %v%"
)

// Formatter implements logrus.Formatter interface.
type Formatter struct {
	// Timestamp format
	TimestampFormat string
	// Available standard keys: time, msg, lvl
	// Also can include custom fields but limited to strings.
	// All of fields need to be wrapped inside %% i.e %time% %msg%
	LogFormat   string
	FieldFormat string
}

// Format building log message.
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	output := f.LogFormat
	if output == "" {
		output = defaultLogFormat
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}
	if f.FieldFormat == "" {
		f.FieldFormat = defaultFieldFormat
	}
	output = strings.Replace(output, "%time%", entry.Time.Format(timestampFormat), 1)

	output = strings.Replace(output, "%msg%", entry.Message, 1)

	level := strings.ToUpper(entry.Level.String())
	output = strings.Replace(output, "%lvl%", level, 1)
	field := ""
	for k, val := range entry.Data {
		if field != "" {
			field += " "
		}
		field += strings.Replace(
			strings.Replace(f.FieldFormat,
				"%k%", fmt.Sprintf("%v", k), 1),
			"%v%", fmt.Sprintf("%v", val), 1)
	}
	output = strings.Replace(output, "%fields%", field, 1)
	return []byte(output), nil
}
