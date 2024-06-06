package app

import (
	"os"
	"path/filepath"
	"time"

	"github.com/ideamans/webpkit/l10n"
)

func AcquireExclusiveLock(path string, expires time.Duration) (bool, error) {
	stat, err := os.Stat(path)
	if err == nil {
		if time.Since(stat.ModTime()) > expires {
			os.RemoveAll(path)
		} else {
			return false, nil
		}
	}

	err = os.MkdirAll(filepath.Dir(path), 0777)
	if err != nil {
		return false, l10n.E("Failed to make directory for lock file %s: %v", path, err)
	}
	_, err = os.Create(path)
	if err != nil {
		return false, l10n.E("Failed to create lock file %s: %v", path, err)
	}

	return true, nil
}
