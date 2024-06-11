package mirror

import (
	"io/fs"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/ideamans/webpkit/app"
	"github.com/ideamans/webpkit/l10n"
	"github.com/ideamans/webpkit/logger"
)

type MirrorApp struct {
	srcDirPath   string
	destDirPath  string
	LockFilePath string
	LockExpires  time.Duration
	WebPToPng    bool
}

func NewMirrorApp(srcDirPath, destDirPath string) *MirrorApp {
	return &MirrorApp{
		srcDirPath:  srcDirPath,
		destDirPath: destDirPath,
		LockExpires: time.Duration(1) * time.Hour,
	}
}

func (ma *MirrorApp) Run() error {
	// Exclusive control
	if ma.LockFilePath != "" {
		available, err := app.AcquireExclusiveLock(ma.LockFilePath, ma.LockExpires)
		if err != nil {
			return l10n.E("Failed to control exclusive: %v", err)
		}
		if !available {
			logger.Info(func() string { return l10n.T("The Application seems to be already running.") })
			return nil
		}
	}

	srcDirAbsPath, err := filepath.Abs(ma.srcDirPath)
	if err != nil {
		return l10n.E("Failed to get absolute path for %s: %v", ma.srcDirPath, err)
	}
	destDirAbsPath, err := filepath.Abs(ma.destDirPath)
	if err != nil {
		return l10n.E("Failed to get absolute path for %s: %v", ma.destDirPath, err)
	}

	id := app.NewImageDouble(srcDirAbsPath, destDirAbsPath, ma.WebPToPng)

	targetExts := []string{".jpg", ".jpeg", ".png", ".gif", ".webp"}
	err = filepath.WalkDir(srcDirAbsPath, func(fullPath string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			// Skip if destDirPath is included under srcDirPath
			if fullPath == destDirAbsPath {
				return filepath.SkipDir
			}
			return nil
		}

		relPath, err := filepath.Rel(srcDirAbsPath, fullPath)
		if err != nil {
			return l10n.E("Failed to get relative path %s in %s: %v", fullPath, srcDirAbsPath, err)
		}

		// Skip if not an image file
		ext := filepath.Ext(relPath)
		if !slices.Contains(targetExts, strings.ToLower(ext)) {
			return nil
		}

		// Ensure image double
		err = id.Ensure(relPath)
		if err != nil {
			logger.Error(func() string { return l10n.F("Failed to ensure image double for %s: %v", relPath, err) })
		}

		return nil
	})

	return err
}
