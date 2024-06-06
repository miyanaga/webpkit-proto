package app

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "webpkit")
	os.RemoveAll(path)
	expires := time.Duration(10) * time.Second

	available1, err := AcquireExclusiveLock(path, expires)
	if err != nil {
		t.Errorf("AcquireExclusiveLock() failed with error %v", err)
	}
	if !available1 {
		t.Errorf("AcquireExclusiveLock() expected to be available but locked")
	}

	available2, err := AcquireExclusiveLock(path, expires)
	if err != nil {
		t.Errorf("AcquireExclusiveLock() failed with error %v", err)
	}
	if available2 {
		t.Errorf("AcquireExclusiveLock() expected to be locked but available")
	}

	twentySecAgo := time.Now().Add(-20 * time.Second)
	os.Chtimes(path, twentySecAgo, twentySecAgo)

	available3, err := AcquireExclusiveLock(path, expires)
	if err != nil {
		t.Errorf("AcquireExclusiveLock() failed with error %v", err)
	}
	if !available3 {
		t.Errorf("AcquireExclusiveLock() expected to be available again but locked")
	}
}
