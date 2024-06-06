//go:build !windows

package main

import (
	"os"
	"strconv"
	"syscall"

	"github.com/ideamans/webpkit/l10n"
)

func SetUmaskOrExit(umaskValue string) {
	umask, err := strconv.ParseUint(umaskValue, 8, 32)
	if err != nil {
		os.Stderr.WriteString(l10n.F("Failed to parse %s as a duration", umaskValue))
		os.Exit(1)
	}
	syscall.Umask(int(umask))
}
