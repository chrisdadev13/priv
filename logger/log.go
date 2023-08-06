package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/motemen/go-colorine"
)

var logger = colorine.NewLogger(
	colorine.Prefixes{
		"git":      colorine.Verbose,
		"hg":       colorine.Verbose,
		"svn":      colorine.Verbose,
		"darcs":    colorine.Verbose,
		"bzr":      colorine.Verbose,
		"fossil":   colorine.Verbose,
		"skip":     colorine.Verbose,
		"cd":       colorine.Verbose,
		"resolved": colorine.Verbose,

		"open":    colorine.Warn,
		"exists":  colorine.Warn,
		"warning": colorine.Warn,

		"authorized": colorine.Notice,

		"error": colorine.Error,
	}, colorine.Info)

func init() {
	SetOutput(os.Stderr)
}

func SetOutput(w io.Writer) {
	logger.SetOutput(w)
}

func Log(prefix, message string) {
	logger.Log(prefix, message)
}

func Logf(prefix, msg string, args ...interface{}) {
	Log(prefix, fmt.Sprintf(msg, args...))
}
