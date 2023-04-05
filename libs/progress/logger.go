package progress

import (
	"encoding/json"
	"io"
	"os"

	"github.com/databricks/bricks/libs/flags"
)

type Logger struct {
	Mode   flags.ProgressLogFormat
	Writer io.Writer

	isFirstEvent bool
}

func NewLogger(mode flags.ProgressLogFormat) *Logger {
	return &Logger{
		Mode:         mode,
		Writer:       os.Stderr,
		isFirstEvent: true,
	}
}

func (l *Logger) Log(event Event) {
	switch l.Mode {
	case flags.ModeInplace:
		if l.isFirstEvent {
			l.Writer.Write([]byte("\n"))
		}
		l.Writer.Write([]byte("\033[1F"))
		l.Writer.Write([]byte(event.String()))
		l.Writer.Write([]byte("\n"))

	case flags.ModeJson:
		b, err := json.MarshalIndent(event, "", "  ")
		if err != nil {
			// we panic because there we cannot catch this in jobs.RunNowAndWait
			panic(err)
		}
		l.Writer.Write([]byte(b))
		l.Writer.Write([]byte("\n"))

	case flags.ModeAppend:
		l.Writer.Write([]byte(event.String()))
		l.Writer.Write([]byte("\n"))

	default:
		// we panic because errors are not captured in some log sides like
		// jobs.RunNowAndWait
		panic("unknown progress logger mode: " + l.Mode.String())
	}
	l.isFirstEvent = false
}
