package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	colorRed    = 31
	colorYellow = 33
	colorBlue   = 36
	colorGray   = 37

	y1  = `0123456789`
	y2  = `0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789`
	y3  = `0000000000111111111122222222223333333333444444444455555555556666666666777777777788888888889999999999`
	y4  = `0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789`
	mo1 = `000000000111`
	mo2 = `123456789012`
	d1  = `0000000001111111111222222222233`
	d2  = `1234567890123456789012345678901`
	h1  = `000000000011111111112222`
	h2  = `012345678901234567890123`
	mi1 = `000000000011111111112222222222333333333344444444445555555555`
	mi2 = `012345678901234567890123456789012345678901234567890123456789`
	s1  = `000000000011111111112222222222333333333344444444445555555555`
	s2  = `012345678901234567890123456789012345678901234567890123456789`
	ns1 = `0123456789`
)

// Fields type, used to pass to `WithFields`.
type Fields map[string]interface{}

// Formatter - logrus formatter, implements logrus.Formatter
type Formatter struct {
	JSONFormat        bool
	DisableColors     bool
	DisableHTMLEscape bool
	PrettyPrint       bool
}

// Format an log entry
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	if f.JSONFormat {
		return f.jsonFormat(entry)
	}
	levelColor := getColorByLevel(entry.Level)

	// output buffer
	b := &bytes.Buffer{}

	// write time
	timeFormat := formatTime(time.Now())
	b.WriteString(string(timeFormat) + " ")

	// caller
	if entry.HasCaller() {
		fmt.Fprintf(b, "(%s:%d)", entry.Caller.File, entry.Caller.Line)
	}

	// level
	if !f.DisableColors {
		fmt.Fprintf(b, "\x1b[%dm", levelColor)
	}
	level := strings.ToUpper(entry.Level.String())
	b.WriteString(" [" + level[:4] + "]")
	if !f.DisableColors {
		b.WriteString("\x1b[0m")
	}
	b.WriteString(" ")

	// fields
	f.writeFields(b, entry)

	// message
	b.WriteString(entry.Message)

	b.WriteByte('\n')

	return b.Bytes(), nil
}

func (f *Formatter) jsonFormat(entry *logrus.Entry) ([]byte, error) {
	data := make(Fields, len(entry.Data)+4)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	data["time"] = string(formatTime(time.Now()))
	data["level"] = entry.Level.String()
	data["msg"] = entry.Message
	if entry.HasCaller() {
		data["file"] = fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	encoder := json.NewEncoder(b)
	encoder.SetEscapeHTML(!f.DisableHTMLEscape)
	if f.PrettyPrint {
		encoder.SetIndent("", "  ")
	}
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}

	return b.Bytes(), nil
}

func getColorByLevel(level logrus.Level) int {
	switch level {
	case logrus.DebugLevel:
		return colorGray
	case logrus.WarnLevel:
		return colorYellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return colorRed
	default:
		return colorBlue
	}
}

func (f *Formatter) writeFields(b *bytes.Buffer, entry *logrus.Entry) {
	if len(entry.Data) != 0 {
		fields := make([]string, 0, len(entry.Data))
		for field := range entry.Data {
			fields = append(fields, field)
		}

		sort.Strings(fields)

		for _, field := range fields {
			fmt.Fprintf(b, "[%s:%v]", field, entry.Data[field])
			b.WriteString(" ")
		}
	}
}

func formatTime(when time.Time) []byte {
	y, mo, d := when.Date()
	h, mi, s := when.Clock()
	ns := when.Nanosecond() / 1000000
	//len("2006/01/02 15:04:05.123")==23
	var buf [23]byte

	buf[0] = y1[y/1000%10]
	buf[1] = y2[y/100]
	buf[2] = y3[y-y/100*100]
	buf[3] = y4[y-y/100*100]
	buf[4] = '/'
	buf[5] = mo1[mo-1]
	buf[6] = mo2[mo-1]
	buf[7] = '/'
	buf[8] = d1[d-1]
	buf[9] = d2[d-1]
	buf[10] = ' '
	buf[11] = h1[h]
	buf[12] = h2[h]
	buf[13] = ':'
	buf[14] = mi1[mi]
	buf[15] = mi2[mi]
	buf[16] = ':'
	buf[17] = s1[s]
	buf[18] = s2[s]
	buf[19] = '.'
	buf[20] = ns1[ns/100]
	buf[21] = ns1[ns%100/10]
	buf[22] = ns1[ns%10]

	//buf[23] = ' '

	return buf[0:]
}
