//go:build linux

package envexec

import (
	"os"

	"golang.org/x/sys/unix"
)

// disableEcho 禁用 TTY 回显（Linux）
func disableEcho(f *os.File) {
	t, err := unix.IoctlGetTermios(int(f.Fd()), unix.TCGETS)
	if err != nil {
		return
	}
	// 清除回显与回车回显
	t.Lflag &^= unix.ECHO | unix.ECHONL
	_ = unix.IoctlSetTermios(int(f.Fd()), unix.TCSETS, t)
}
