package dwebp

import (
	"testing"

	"github.com/ideamans/webpkit/imagetype"
	"github.com/phuslu/fastimage"
)

func TestDWebPNormalCase(t *testing.T) {
	tmp := t.TempDir()

	srcPath := "../testdata/simple/simple.webp"
	destPath := tmp + "/sample.png"
	code := DWebP("-quiet", "-o", destPath, srcPath)
	if code != 0 {
		t.Errorf("dwebp command exited with code %d", code)
	}

	webp, err := imagetype.FastImage(srcPath)
	if err != nil {
		t.Errorf("Failed to read webp file: %v", err)
	}
	png, err := imagetype.FastImage(destPath)
	if err != nil {
		t.Errorf("Failed to read png file: %v", err)
	}

	if png.Type != fastimage.PNG {
		t.Errorf("Output file is not a png image")
	}
	if webp.Width != png.Width || webp.Height != png.Height {
		t.Errorf("Image dimensions are not %dx%d but %dx%d", webp.Width, webp.Height, png.Width, png.Height)
	}
}
