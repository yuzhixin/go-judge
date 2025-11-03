//go:build !linux

package envexec

import "os"

// disableEcho 在非 Linux 平台为 no-op，用于保持编译通过
func disableEcho(f *os.File) {}
