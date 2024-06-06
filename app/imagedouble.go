package app

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ideamans/webpkit/converter"
	"github.com/ideamans/webpkit/l10n"
	"github.com/ideamans/webpkit/logger"
)

type ImageDoubleInterface interface {
	Ensure(relPath string) error
}

type ImageDouble struct {
	srcDirPath  string
	destDirPath string
}

func DefaultNewImageDouble(srcDirPath, destDirPath string) ImageDoubleInterface {
	id := ImageDouble{
		srcDirPath:  srcDirPath,
		destDirPath: destDirPath,
	}
	return &id
}

var (
	NewImageDouble = DefaultNewImageDouble
)

func (id *ImageDouble) Ensure(relPath string) error {
	// Extension pair
	srcExt := strings.ToLower(filepath.Ext(relPath))
	doubleExt := ".webp"
	if srcExt == ".jpg" || srcExt == ".jpeg" || srcExt == ".png" || srcExt == ".gif" {
		doubleExt = ".webp"
	} else if srcExt == ".webp" {
		doubleExt = ".png"
	}
	errorExt := ".err"

	// Source file
	srcPath := filepath.Join(id.srcDirPath, relPath)
	srcStat, err := os.Stat(srcPath)
	if err != nil {
		return l10n.E("ImageDouble failed to get stat src file %s: %v", srcPath, err)
	}
	srcModTime := srcStat.ModTime()

	// Double file and error file
	var doubleFileValid, errorFileValid bool

	// Double file is valid if exists and has same modTime
	doubleFilePath := filepath.Join(id.destDirPath, relPath+doubleExt)
	doubleFileStat, err := os.Stat(doubleFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			doubleFileValid = false
		} else {
			return l10n.E("ImageDouble failed to get stat double file %s: %v", doubleFilePath, err)
		}
	} else {
		doubleFileValid = doubleFileStat.ModTime().Unix() == srcModTime.Unix()
	}

	// Error file is valid if exists and has same modTime
	errorFilePath := filepath.Join(id.destDirPath, relPath+errorExt)
	errorFileStat, err := os.Stat(errorFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			errorFileValid = false
		} else {
			return l10n.E("ImageDouble failed to get stat error file %s: %v", errorFilePath, err)
		}
	} else {
		errorFileValid = errorFileStat.ModTime().Unix() == srcModTime.Unix()
	}

	// Skip double or error is valid
	if doubleFileValid || errorFileValid {
		logger.Debug(func() string { return l10n.F("Skip %s because of no update", relPath) })
		return nil
	}

	// Convert source file to double file
	os.MkdirAll(filepath.Dir(doubleFilePath), 0777)
	err = converter.Singleton.Convert(srcPath, doubleFilePath)
	if err != nil {
		logger.Error(func() string { return l10n.F("Failed to convert %s: %v", relPath, err) })
		os.RemoveAll(doubleFilePath)
		_, err = os.Create(errorFilePath)
		if err != nil {
			return err
		}
		os.Chtimes(errorFilePath, srcModTime, srcModTime)
	} else {
		logger.Info(func() string { return l10n.F("Converted %s", relPath) })
		os.RemoveAll(errorFilePath)
		os.Chtimes(doubleFilePath, srcModTime, srcModTime)
	}

	return nil
}
