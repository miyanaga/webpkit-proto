package beside

import (
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"
	"time"

	cp "github.com/otiai10/copy"
)

func recursiveEntries(rootPath string) []string {
	entries := make([]string, 0)
	filepath.WalkDir(rootPath, func(fullPath string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(rootPath, fullPath)
		if err != nil {
			return err
		}
		entries = append(entries, relPath)

		return nil
	})

	sort.Strings(entries)
	return entries
}

func TestBeside(t *testing.T) {
	tmp := t.TempDir()
	cp.Copy("../testdata/tree/", tmp+"/")

	ba := NewBesideApp(tmp)
	ba.WebPToPng = true
	err := ba.Run()
	if err != nil {
		t.Errorf("Failed to run BesideApp: %v", err)
	}

	entriesFirst := recursiveEntries(tmp)
	expectedFirst := []string{
		"dir/cmyk.jpg",
		"dir/cmyk.jpg.err",
		"dir/not-image.txt",
		"dir/simple.jpg",
		"dir/simple.jpg.webp",
		"simple.png",
		"simple.png.webp",
		"simple.webp",
		"simple.webp.png",
	}
	if !reflect.DeepEqual(entriesFirst, expectedFirst) {
		t.Errorf("Unexpected file entries. Expected %v but got %v", expectedFirst, entriesFirst)
	}

	// Fix the error
	cp.Copy(filepath.Join(tmp, "dir/simple.jpg"), filepath.Join(tmp, "dir/cmyk.jpg"))
	future := time.Now().Add(10 * time.Second)
	os.Chtimes(filepath.Join(tmp, "dir/cmyk.jpg"), future, future)

	err = ba.Run()
	if err != nil {
		t.Errorf("Failed to run BesideApp: %v", err)
	}

	entriesSecond := recursiveEntries(tmp)
	expectedSecond := []string{
		"dir/cmyk.jpg",
		"dir/cmyk.jpg.webp",
		"dir/not-image.txt",
		"dir/simple.jpg",
		"dir/simple.jpg.webp",
		"simple.png",
		"simple.png.webp",
		"simple.webp",
		"simple.webp.png",
	}
	if !reflect.DeepEqual(entriesSecond, expectedSecond) {
		t.Errorf("Unexpected file entries. Expected %v but got %v", expectedSecond, entriesSecond)
	}
}
