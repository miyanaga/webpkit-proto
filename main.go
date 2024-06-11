package main

import (
	"fmt"
	"os"

	"github.com/ideamans/webpkit/l10n"
)

var (
	Version = "1.0.0-b2"
)

func main() {
	l10n.DetectLanguage()
	rootCmd := BuildCommand()

	if err := rootCmd.Execute(); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("%v\n", err))
		os.Exit(1)
	}
}
