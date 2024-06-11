package app

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	cp "github.com/otiai10/copy"
)

func TestImageDoubleSuccessThenError(t *testing.T) {
	tmp := t.TempDir()

	// First time, testdata/tree/dir/simple.jpg
	srcRootPath := tmp + "/tree/"
	cp.Copy("../testdata/tree/", srcRootPath)
	destRootPath := tmp + "/.mirror/"

	id := NewImageDouble(srcRootPath, destRootPath, true)
	relPath := "dir/simple.jpg"
	err := id.Ensure(relPath)

	if err != nil {
		t.Errorf("Failed to ensure image double for %s: %v", relPath, err)
	}
	srcPath := filepath.Join(srcRootPath, relPath)
	srcStat, err := os.Stat(srcPath)
	if err != nil {
		t.Errorf("Failed to get stat src file: %v", err)
	}
	doublePath := filepath.Join(destRootPath, relPath+".webp")
	doubleStat, err := os.Stat(doublePath)
	if err != nil {
		t.Errorf("Failed to get stat dest file: %v", err)
	}
	if doubleStat.ModTime().Unix() != srcStat.ModTime().Unix() {
		t.Errorf("ModTime is different between src and double expected %v but got %v", srcStat.ModTime(), doubleStat.ModTime())
	}

	// Update src file to get error
	cp.Copy("../testdata/tree/dir/cmyk.jpg", srcPath)
	future := time.Now().Add(10 * time.Second)
	os.Chtimes(srcPath, future, future)

	err = id.Ensure(relPath)
	if err != nil {
		t.Errorf("Failed to ensure image double for %s: %v", relPath, err)
	}

	srcStat2, err := os.Stat(srcPath)
	if err != nil {
		t.Errorf("Failed to get stat src file: %v", err)
	}
	errorPath := filepath.Join(destRootPath, relPath+".err")
	errorStat, err := os.Stat(errorPath)
	if err != nil {
		t.Errorf("Failed to get stat error file: %v", err)
	}
	if errorStat.ModTime().Unix() != srcStat2.ModTime().Unix() {
		t.Errorf("ModTime is different between src and error expected %v but got %v", srcStat2.ModTime(), errorStat.ModTime())
	}

	_, err = os.Stat(doublePath)
	if err == nil {
		t.Errorf("Double file should be removed but exists")
	}
}

func TestImageDoubleErrorThenSuccess(t *testing.T) {
	tmp := t.TempDir()

	// First time, testdata/tree/dir/cmyk.jpg
	srcRootPath := tmp + "/tree/"
	cp.Copy("../testdata/tree/", srcRootPath)
	destRootPath := tmp + "/.mirror/"

	id := NewImageDouble(srcRootPath, destRootPath, true)
	relPath := "dir/cmyk.jpg"
	err := id.Ensure(relPath)

	if err != nil {
		t.Errorf("Failed to ensure image double for %s: %v", relPath, err)
	}
	srcPath := filepath.Join(srcRootPath, relPath)
	srcStat, err := os.Stat(srcPath)
	if err != nil {
		t.Errorf("Failed to get stat src file: %v", err)
	}
	errorPath := filepath.Join(destRootPath, relPath+".err")
	errorStat, err := os.Stat(errorPath)
	if err != nil {
		t.Errorf("Failed to get stat dest file: %v", err)
	}
	if errorStat.ModTime().Unix() != srcStat.ModTime().Unix() {
		t.Errorf("ModTime is different between src and error expected %v but got %v", srcStat.ModTime(), errorStat.ModTime())
	}

	// Update src file to get success
	cp.Copy("../testdata/tree/dir/simple.jpg", srcPath)
	future := time.Now().Add(10 * time.Second)
	os.Chtimes(srcPath, future, future)

	err = id.Ensure(relPath)
	if err != nil {
		t.Errorf("Failed to ensure image double for %s: %v", relPath, err)
	}

	srcStat2, err := os.Stat(srcPath)
	if err != nil {
		t.Errorf("Failed to get stat src file: %v", err)
	}
	doublePath := filepath.Join(destRootPath, relPath+".webp")
	doubleStat, err := os.Stat(doublePath)
	if err != nil {
		t.Errorf("Failed to get stat error file: %v", err)
	}
	if doubleStat.ModTime().Unix() != srcStat2.ModTime().Unix() {
		t.Errorf("ModTime is different between src and double expected %v but got %v", srcStat2.ModTime(), errorStat.ModTime())
	}

	_, err = os.Stat(errorPath)
	if err == nil {
		t.Errorf("Error file should be removed but exists")
	}
}
