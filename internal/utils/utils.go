package utils

import (
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

// Flog adds returns a logrus.Entry with fields for the current function and file.
func Flog() *logrus.Entry {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Could not get context info for logger!")
	}

	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	fn := funcname[strings.LastIndex(funcname, ".")+1:]
	return logrus.WithField("file", filename).WithField("function", fn)
}
