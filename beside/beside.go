package beside

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

type BesideApp struct {
	dirPath      string
	LockFilePath string
	LockExpires  time.Duration
	WebPToPng    bool
}

func NewBesideApp(dirPath string) *BesideApp {
	return &BesideApp{
		dirPath:     dirPath,
		LockExpires: time.Duration(1) * time.Hour,
	}
}

func (ba *BesideApp) Run() error {
	// Exclusive control
	if ba.LockFilePath != "" {
		available, err := app.AcquireExclusiveLock(ba.LockFilePath, ba.LockExpires)
		if err != nil {
			return l10n.E("Failed to control exclusive: %v", err)
		}
		if !available {
			logger.Info(func() string { return l10n.T("The Application seems to be already running.") })
			return nil
		}
	}

	id := app.NewImageDouble(ba.dirPath, ba.dirPath, ba.WebPToPng)

	targetExts := []string{".jpg", ".jpeg", ".png", ".gif", ".webp"}
	err := filepath.WalkDir(ba.dirPath, func(fullPath string, info fs.DirEntry, err error) error {
		if err != nil {
			return l10n.E("Got an error while walking %s: %v", fullPath, err)
		}

		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(ba.dirPath, fullPath)
		if err != nil {
			return l10n.E("Failed to get relative path %s in %s: %v", fullPath, ba.dirPath, err)
		}

		// Skip if not an image file or not an image double file (e.g. .jpg.webp)
		lastExt := filepath.Ext(relPath)
		noExt1 := strings.TrimSuffix(relPath, lastExt)
		origExt := filepath.Ext(noExt1)
		if !slices.Contains(targetExts, strings.ToLower(lastExt)) {
			return nil
		}
		if slices.Contains(targetExts, strings.ToLower(origExt)) {
			return nil
		}

		// Ensure image double
		err = id.Ensure(relPath)
		if err != nil {
			logger.Error(func() string { return l10n.F("Failed to ensure the image double for %s: %v", relPath, err) })
			return err
		}

		return nil
	})

	return err
}
