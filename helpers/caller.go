package helpers

import (
	"runtime"
	"strings"
)

// OuterCaller returns the first caller outside current package
func OuterCaller() (pc uintptr, file string, line int, ok bool) {
	// Skip more frames to include this function
	const skip = 2 // skip GetOuterCaller and its immediate caller
	var currentPkg string

	// Get current package name from the first valid frame
	pcArray := make([]uintptr, 1)
	if runtime.Callers(skip, pcArray) > 0 {
		if frame, _ := runtime.CallersFrames(pcArray).Next(); frame.Func != nil {
			currentPkg = getPackageName(frame.Function)
		}
	}

	// Search for external caller
	for i := skip + 1; ; i++ {
		pc, file, line, ok = runtime.Caller(i)
		if !ok {
			return 0, "", 0, false
		}

		// Get function info
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}

		// Get package name from function
		callerPkg := getPackageName(fn.Name())

		// Check if caller is from different package
		if callerPkg != currentPkg && callerPkg != "" {
			return pc, file, line, true
		}
	}
}

func getPackageName(funcName string) string {
	// Remove method receiver if present
	if idx := strings.LastIndex(funcName, "."); idx != -1 {
		// Check if there's another dot before (for methods with receiver)
		if prevIdx := strings.LastIndex(funcName[:idx], "."); prevIdx != -1 {
			return funcName[:prevIdx]
		}
		return funcName[:idx]
	}
	return ""
}
