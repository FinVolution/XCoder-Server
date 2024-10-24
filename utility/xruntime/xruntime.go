package xruntime

import (
	"bytes"
	"fmt"
	"runtime"
)

func MStack(deep, skip int) string {
	pc := make([]uintptr, deep+1) // at least 1 entry needed
	n := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc[:n])
	mstack := bytes.NewBufferString("")
	for {
		frame, more := frames.Next()
		mstack.WriteString(fmt.Sprintf("%s:%d;%s;", frame.File, frame.Line, frame.Function))
		if !more {
			break
		}
	}
	return mstack.String()
}
