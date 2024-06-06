package cwebp

import (
	"testing"

	"github.com/ideamans/webpkit/imagetype"
	"github.com/phuslu/fastimage"
)

func TestCWebPNormalCase(t *testing.T) {
	tmp := t.TempDir()

	srcPath := "../testdata/simple/simple.jpg"
	destPath := tmp + "/simple.webp"
	code := CWebP("-quiet", "-o", destPath, srcPath)
	if code != 0 {
		t.Errorf("cwebp command exited with code %d", code)
	}

	jpg, err := imagetype.FastImage(srcPath)
	if err != nil {
		t.Errorf("Failed to read jpg file %s: %v", srcPath, err)
	}
	webp, err := imagetype.FastImage(destPath)
	if err != nil {
		t.Errorf("Failed to read webp file %s: %v", destPath, err)
	}

	if webp.Type != fastimage.WEBP {
		t.Errorf("Output file is not a webp image: %v", webp.Type)
	}
	if jpg.Width != webp.Width || jpg.Height != webp.Height {
		t.Errorf("Dimensions are not %dx%d but %dx%d", jpg.Width, jpg.Height, webp.Width, webp.Height)
	}
}

func TestCWebPErrorCase(t *testing.T) {
	tmp := t.TempDir()

	srcPath := "../testdata/error/cmyk.jpg"
	destPath := tmp + "/cmyk.webp"
	code := CWebP("-quiet", "-o", destPath, srcPath)
	if code == 0 {
		t.Errorf("cwebp should be failed")
	}
}
