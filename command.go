package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/ideamans/webpkit/beside"
	"github.com/ideamans/webpkit/converter"
	"github.com/ideamans/webpkit/cwebp"
	"github.com/ideamans/webpkit/dwebp"
	"github.com/ideamans/webpkit/gif2webp"
	"github.com/ideamans/webpkit/l10n"
	"github.com/ideamans/webpkit/logger"
	"github.com/ideamans/webpkit/mirror"
	"github.com/ideamans/webpkit/webpinfo"
	"github.com/spf13/cobra"
)

type AppOptions struct {
	LogLevel     string
	WebPToPng    bool
	Umask        string
	LockFilePath string
	LockExpires  string
}

var (
	GlobalAppOptions AppOptions
)

func SetAppOptions(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&GlobalAppOptions.WebPToPng, "webp-to-png", "", false, l10n.T("Convert WebP to PNG (default is WebP to JPEG)"))
	cmd.PersistentFlags().StringVarP(&GlobalAppOptions.LogLevel, "log-level", "l", "info", l10n.T("Log level to display (trace, debug, info, warn, error, fatal, silent)"))
	cmd.PersistentFlags().StringVarP(&GlobalAppOptions.LockFilePath, "lock-file", "", "", "Exclusive lock file path to control exclusive")
	cmd.PersistentFlags().StringVarP(&GlobalAppOptions.LockExpires, "lock-expires", "", "1h", "Exclusive lock expires (e.g. 1h, 1d)")
	cmd.PersistentFlags().StringVarP(&GlobalAppOptions.Umask, "umask", "", "022", l10n.T("Umask for file and directory creation"))
}

func ParseAppOptionsOrExit() (string, time.Duration, bool) {
	logger.LogLevel = logger.ParseLogLevel(GlobalAppOptions.LogLevel)
	lockExpires := ParseDurationOrExit(GlobalAppOptions.LockExpires)
	return GlobalAppOptions.LockFilePath, lockExpires, GlobalAppOptions.WebPToPng
}

func ParseDurationOrExit(durationStr string) time.Duration {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		os.Stderr.WriteString(l10n.F("Failed to parse %s as a umask value", durationStr))
		os.Exit(1)
	}
	return duration
}

func BuildCommand() *cobra.Command {
	rootCmd := cobra.Command{
		Use:   "webpkit",
		Short: l10n.T("Toolkit for converting conventionally formatted Web images to WebP"),
	}
	rootCmd.Flags().SetInterspersed(false)

	// version sub command
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: l10n.T("Print the version number"),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s %s %s", Version, runtime.GOOS, runtime.GOARCH)
		},
	}
	rootCmd.AddCommand(versionCmd)

	// raw command
	rawCmd := &cobra.Command{
		Use:   "raw [cwebp|dwebp|gif2webp|webpinfo] <...args>",
		Short: l10n.T("Alias for a raw command of libwebp"),
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			command := args[0]
			args = args[1:]
			if command == "cwebp" {
				os.Exit(cwebp.CWebP(args...))
			} else if command == "dwebp" {
				os.Exit(dwebp.DWebP(args...))
			} else if command == "gif2webp" {
				os.Exit(gif2webp.Gif2WebP(args...))
			} else if command == "webpinfo" {
				os.Exit(webpinfo.WebPInfo(args...))
			} else {
				os.Stderr.WriteString(l10n.F("%s is unavailable command", command))
			}
		},
	}
	rawCmd.Flags().SetInterspersed(false)
	rootCmd.AddCommand(rawCmd)

	// convert sub command
	convertCmd := cobra.Command{
		Use:   "convert [srcFilePath] [destFilePath]",
		Short: l10n.T("Convert a single image file to WebP format"),
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			srcPath := args[0]
			destPath := args[1]
			err := converter.Singleton.Convert(srcPath, destPath)
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("%v\n", err))
				os.Exit(1)
			}
		},
	}
	convertCmd.PersistentFlags().StringVarP(&GlobalAppOptions.Umask, "umask", "", "022", l10n.T("Umask for file and directory creation"))
	rootCmd.AddCommand(&convertCmd)

	// mirror sub command
	mirrorCmd := cobra.Command{
		Use:   "mirror [srcDirPath] [destDirPath]",
		Short: l10n.T("Convert image files under the directory to WebP format (or Reverse) as another tree"),
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			srcDirPath := args[0]
			destDirPath := args[1]
			app := mirror.NewMirrorApp(srcDirPath, destDirPath)

			SetUmaskOrExit(GlobalAppOptions.Umask)
			app.LockFilePath, app.LockExpires, app.WebPToPng = ParseAppOptionsOrExit()

			err := app.Run()
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("%v\n", err))
				os.Exit(1)
			}
		},
	}
	SetAppOptions(&mirrorCmd)
	rootCmd.AddCommand(&mirrorCmd)

	// beside command
	besideCmd := cobra.Command{
		Use:   "beside [dirPath]",
		Short: l10n.T("Convert image files under the directory to WebP format (or Reverse) beside each file"),
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dirPath := args[0]
			app := beside.NewBesideApp(dirPath)

			SetUmaskOrExit(GlobalAppOptions.Umask)
			app.LockFilePath, app.LockExpires, app.WebPToPng = ParseAppOptionsOrExit()

			err := app.Run()
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("%v\n", err))
				os.Exit(1)
			}
		},
	}
	SetAppOptions(&besideCmd)
	rootCmd.AddCommand(&besideCmd)

	return &rootCmd
}
