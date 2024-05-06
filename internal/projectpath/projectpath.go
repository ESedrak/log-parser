package projectpath

import (
	"path/filepath"
	"runtime"
)

// get current root of package structure
var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	Root = filepath.Join(filepath.Dir(b), "../..")
)
